/*
Вот реализация шардированного in-memory кэша с конкурентным доступом на Go.
Кэш разбит на несколько сегментов (шардов), каждый из которых имеет свой мьютекс для обеспечения безопасности при одновременном доступе.
Ключи распределяются по шардам с помощью хэширования, что позволяет равномерно загрузить систему и избежать блокировок между шардами.

Основные компоненты:

1. **Интерфейс Cache** - определяет основные методы работы с кэшем
2. **lruNode** - узел двусвязного списка для реализации LRU (без boxing)
3. **shard** - сегмент кэша, содержащий карту данных, LRU список и мьютекс для конкурентного доступа
4. **shardCache** - основной контейнер кэша с набором шард
5. **getShardIndex** - вычисляет индекс шарда для заданного ключа через type-aware хеширование
6. **Set** - устанавливает значение с блокировкой соответствующего шарда и обновлением LRU
7. **Get** - получает значение с блокировкой и обновлением LRU

Ошибка `ErrNotFound` возвращается, когда ключ не найден в кэше.

Оптимизации производительности:
- Type-aware хеширование избегает fmt.Sprintf для базовых типов
- Специализированный LRU список без interface{}
- Значения типа V хранятся без boxing
- Использование sync.Pool для переиспользования fnv хешеров
*/
package cache

import (
	"fmt"
	"hash"
	"hash/fnv"
	"sync"
)

// Cache представляет интерфейс кэша с методами Set и Get
type Cache[K comparable, V any] interface {
	Set(key K, val V)
	Get(key K) (V, error)
}

// lruNode представляет узел в LRU списке
// value теперь типа V, что избегает boxing
type lruNode[K comparable, V any] struct {
	key   K
	value V
	prev  *lruNode[K, V]
	next  *lruNode[K, V]
}

// lruList представляет оптимизированный двусвязный список для LRU
// Специализированный для наших нужд, без interface{}
type lruList[K comparable, V any] struct {
	head *lruNode[K, V]
	tail *lruNode[K, V]
	len  int
}

// newLruList создает новый LRU список
func newLruList[K comparable, V any]() *lruList[K, V] {
	head := &lruNode[K, V]{}
	tail := &lruNode[K, V]{}
	head.next = tail
	tail.prev = head
	return &lruList[K, V]{
		head: head,
		tail: tail,
	}
}

// Len возвращает длину списка
func (l *lruList[K, V]) Len() int {
	return l.len
}

// PushFront добавляет узел в начало списка
func (l *lruList[K, V]) PushFront(node *lruNode[K, V]) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	l.len++
}

// MoveToFront перемещает существующий узел в начало списка
func (l *lruList[K, V]) MoveToFront(node *lruNode[K, V]) {
	// Удаляем узел из текущей позиции
	node.prev.next = node.next
	node.next.prev = node.prev

	// Вставляем в начало
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

// Remove удаляет узел из списка
func (l *lruList[K, V]) Remove(node *lruNode[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.len--
}

// Back возвращает последний узел списка (наименее используемый)
func (l *lruList[K, V]) Back() *lruNode[K, V] {
	if l.tail.prev == l.head {
		return nil
	}
	return l.tail.prev
}

// shard представляет отдельный сегмент кэша с мьютексом для конкурентного доступа
type shard[K comparable, V any] struct {
	data     map[K]*lruNode[K, V]
	lruList  *lruList[K, V]
	capacity int
	mu       sync.RWMutex
}

// shardCache представляет шардированный кэш
type shardCache[K comparable, V any] struct {
	shards    []*shard[K, V]
	count     int
	shardSize int
}

// NewCache создает новый экземпляр шардированного кэша с LRU
// shardsCount - количество шардов для распределения ключей
// shardSize - максимальное количество элементов в каждом шарде
func NewCache[K comparable, V any](shardsCount int, shardSize int) Cache[K, V] {
	if shardsCount <= 0 {
		shardsCount = 1
	}
	if shardSize <= 0 {
		shardSize = 1000
	}

	cache := &shardCache[K, V]{
		shards:    make([]*shard[K, V], shardsCount),
		count:     shardsCount,
		shardSize: shardSize,
	}

	for i := 0; i < shardsCount; i++ {
		cache.shards[i] = &shard[K, V]{
			data:     make(map[K]*lruNode[K, V]),
			lruList:  newLruList[K, V](),
			capacity: shardSize,
		}
	}

	return cache
}

// fnvPool для переиспользования fnv хешеров, чтобы избежать аллокаций
var fnvPool = sync.Pool{
	New: func() any {
		return fnv.New32a()
	},
}

// getShardIndex возвращает индекс шарда для заданного ключа
// Использует type-aware хеширование для базовых типов без fmt.Sprintf
// Для сложных типов использует fnv через sync.Pool
func (c *shardCache[K, V]) getShardIndex(key K) int {
	// Для базовых типов используем быстрое хеширование без fmt.Sprintf
	switch v := any(key).(type) {
	case int:
		return int(uint32(v) % uint32(c.count))
	case int8:
		return int(uint32(v) % uint32(c.count))
	case int16:
		return int(uint32(v) % uint32(c.count))
	case int32:
		return int(uint32(v) % uint32(c.count))
	case int64:
		return int((uint32(v) ^ uint32(v>>32)) % uint32(c.count))
	case uint:
		return int(uint32(v) % uint32(c.count))
	case uint8:
		return int(uint32(v) % uint32(c.count))
	case uint16:
		return int(uint32(v) % uint32(c.count))
	case uint32:
		return int(v % uint32(c.count))
	case uint64:
		return int(uint32((v ^ v>>32)) % uint32(c.count))
	case uintptr:
		return int((uint32(v) ^ uint32(v>>32)) % uint32(c.count))
	case string:
		return int(hashString(v) % uint32(c.count))
	case bool:
		if v {
			return 1 % c.count
		}
		return 0
	default:
		// Для сложных типов используем fnv через pool
		return int(hashGeneric(key) % uint32(c.count))
	}
}

// hashString вычисляет хеш строки без аллокаций
func hashString(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = h*31 + uint32(s[i])
	}
	return h
}

// hashGeneric хеширует сложные типы через fnv с использованием sync.Pool
func hashGeneric[T comparable](key T) uint32 {
	hasher := fnvPool.Get().(hash.Hash32)
	defer fnvPool.Put(hasher)

	hasher.Reset()

	// Используем fmt.Sprintf только для сложных типов
	str := fmt.Sprintf("%v", key)
	hasher.Write([]byte(str))

	return hasher.Sum32()
}

// Set устанавливает значение в кэш
func (c *shardCache[K, V]) Set(key K, val V) {
	shardIdx := c.getShardIndex(key)
	shard := c.shards[shardIdx]

	shard.mu.Lock()
	defer shard.mu.Unlock()

	// Если ключ уже существует, обновляем значение и перемещаем в начало LRU
	if node, exists := shard.data[key]; exists {
		node.value = val
		shard.lruList.MoveToFront(node)
		return
	}

	// Если кэш достиг максимального размера, удаляем наименее недавно использованный элемент
	if shard.lruList.Len() >= shard.capacity {
		backNode := shard.lruList.Back()
		if backNode != nil {
			delete(shard.data, backNode.key)
			shard.lruList.Remove(backNode)
		}
	}

	// Добавляем новый элемент в начало LRU списка
	newNode := &lruNode[K, V]{key: key, value: val}
	shard.lruList.PushFront(newNode)
	shard.data[key] = newNode
}

// Get получает значение из кэша
// Используем write lock сразу, чтобы избежать race condition между RUnlock и Lock
func (c *shardCache[K, V]) Get(key K) (V, error) {
	shardIdx := c.getShardIndex(key)
	shard := c.shards[shardIdx]

	shard.mu.Lock()
	defer shard.mu.Unlock()

	node, exists := shard.data[key]
	if !exists {
		var zero V
		return zero, ErrNotFound
	}

	// Перемещаем элемент в начало LRU списка
	shard.lruList.MoveToFront(node)

	// Возвращаем значение напрямую без type assertion (благодаря V вместо any)
	return node.value, nil
}

// ErrNotFound представляет ошибку, когда ключ не найден в кэше
var ErrNotFound = &NotFoundError{}

// NotFoundError это структура ошибки для несуществующего ключа
type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "key not found in cache"
}
