package cache

import (
	"fmt"
	"sync"
	"testing"
)

// TestSetAndGet тестирует базовые операции Set и Get
func TestSetAndGet(t *testing.T) {
	cache := NewCache[int, string](10, 100)

	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "three")

	if val, err := cache.Get(1); err != nil {
		t.Errorf("Expected to get value for key 1, got error: %v", err)
	} else if val != "one" {
		t.Errorf("Expected value 'one', got '%s'", val)
	}

	if val, err := cache.Get(3); err != nil {
		t.Errorf("Expected to get value for key 3, got error: %v", err)
	} else if val != "three" {
		t.Errorf("Expected value 'three', got '%s'", val)
	}
}

// TestGetNotFound тестирует получение несуществующего ключа
func TestGetNotFound(t *testing.T) {
	cache := NewCache[int, string](10, 100)

	if _, err := cache.Get(999); err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got: %v", err)
	}

	// Проверяем, что возвращается нулевое значение для типа V
	if _, err := cache.Get(999); err == nil {
		t.Error("Expected error, got nil")
	}
}

// TestConcurrentAccess тестирует конкурентный доступ к кэшу
func TestConcurrentAccess(t *testing.T) {
	cache := NewCache[int, string](10, 100)
	var wg sync.WaitGroup

	// Сначала заполняем кэш данными
	for i := 0; i < 100; i++ {
		cache.Set(i, "value")
	}

	// Затем запускаем горутины для конкурентного чтения
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(key int) {
			defer wg.Done()
			val, _ := cache.Get(key)
			if val != "value" {
				t.Errorf("Expected 'value' for key %d, got '%s'", key, val)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcurrentSet тестирует конкурентные операции Set
func TestConcurrentSet(t *testing.T) {
	cache := NewCache[int, string](10, 100)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int, value string) {
			defer wg.Done()
			cache.Set(key, value)
		}(i, "value")
	}

	wg.Wait()

	// Проверяем, что все значения были установлены
	for i := 0; i < 100; i++ {
		val, err := cache.Get(i)
		if err != nil || val != "value" {
			t.Errorf("Key %d not found or wrong value", i)
		}
	}
}

// TestDifferentKeyTypes тестирует работу с разными типами ключей
func TestDifferentKeyTypes(t *testing.T) {
	cache := NewCache[string, int](5, 100)

	// Строковые ключи
	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("user:alice", 100)

	if val, err := cache.Get("key1"); err != nil {
		t.Errorf("Expected to get value for key 'key1', got error: %v", err)
	} else if val != 1 {
		t.Errorf("Expected value 1, got %d", val)
	}

	// Структуры как ключи
	type User struct {
		ID   int
		Name string
	}
	user := User{ID: 42, Name: "John"}
	// Преобразуем структуру в строку для использования в качестве ключа
	userKey := fmt.Sprintf("%v", user)
	cache.Set(userKey, 999)

	if val, err := cache.Get(userKey); err != nil {
		t.Errorf("Expected to get value for struct key, got error: %v", err)
	} else if val != 999 {
		t.Errorf("Expected value 999, got %d", val)
	}
}

// TestShardDistribution тестирует распределение ключей по шардам
func TestShardDistribution(t *testing.T) {
	cache := NewCache[int, int](4, 100)

	// Добавляем много ключей и проверяем, что они корректно распределены
	for i := 0; i < 100; i++ {
		cache.Set(i, i)
	}

	// Проверяем, что все ключи доступны
	for i := 0; i < 100; i++ {
		val, err := cache.Get(i)
		if err != nil || val != i {
			t.Errorf("Key %d not found or wrong value", i)
		}
	}
}

// TestOverwrite тестирует перезапись значения
func TestOverwrite(t *testing.T) {
	cache := NewCache[string, int](5, 100)

	cache.Set("key", 1)
	if val, err := cache.Get("key"); err != nil || val != 1 {
		t.Errorf("Expected value 1, got %d", val)
	}

	// Перезаписываем значение
	cache.Set("key", 2)
	if val, err := cache.Get("key"); err != nil || val != 2 {
		t.Errorf("Expected value 2 after overwrite, got %d", val)
	}
}

// TestGetReturnsZeroValue тестирует, что при ошибке возвращается нулевое значение типа V
func TestGetReturnsZeroValue(t *testing.T) {
	cache := NewCache[int, string](10, 100)

	// Для строки нулевое значение - это ""
	val, err := cache.Get(999)
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got: %v", err)
	}
	if val != "" {
		t.Errorf("Expected zero value for string, got '%s'", val)
	}
}

// TestErrorType тестирует тип ошибки NotFoundError
func TestErrorType(t *testing.T) {
	cache := NewCache[int, string](10, 100)

	_, err := cache.Get(999)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %T", err)
	}

	// Проверяем, что ошибка не является nil
	if err != nil {
		if _, ok := err.(*NotFoundError); !ok {
			t.Errorf("Error should be of type *NotFoundError, got %T", err)
		}
	}
}

// BenchmarkConcurrentRead тестирует производительность конкурентного чтения
func BenchmarkConcurrentRead(b *testing.B) {
	cache := NewCache[int, int](10, 100)

	// Подготовка данных
	for i := 0; i < 100; i++ {
		cache.Set(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Get(i % 100)
		}(i)
	}
	wg.Wait()
}

// BenchmarkConcurrentReadLarge тестирует производительность конкурентного чтения с 100K элементами
func BenchmarkConcurrentReadLarge(b *testing.B) {
	cache := NewCache[int, int](10, 100000)

	// Подготовка данных
	for i := 0; i < 100000; i++ {
		cache.Set(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Get(i % 100000)
		}(i)
	}
	wg.Wait()
}

// BenchmarkConcurrentReadHuge тестирует производительность конкурентного чтения с 10M элементами
func BenchmarkConcurrentReadHuge(b *testing.B) {
	cache := NewCache[int, int](10, 10000000)

	// Подготовка данных
	for i := 0; i < 10000000; i++ {
		cache.Set(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Get(i % 10000000)
		}(i)
	}
	wg.Wait()
}

// BenchmarkConcurrentWrite тестирует производительность конкурентной записи
func BenchmarkConcurrentWrite(b *testing.B) {
	cache := NewCache[int, int](10, 100)

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkConcurrentWriteLarge тестирует производительность конкурентной записи с 100K элементами
func BenchmarkConcurrentWriteLarge(b *testing.B) {
	cache := NewCache[int, int](10, 100000)

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkConcurrentWriteHuge тестирует производительность конкурентной записи с 10M элементами
func BenchmarkConcurrentWriteHuge(b *testing.B) {
	cache := NewCache[int, int](10, 10000000)

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapRead тестирует производительность чтения из sync.Map
func BenchmarkSyncMapRead(b *testing.B) {
	var syncMap sync.Map

	// Подготовка данных
	for i := 0; i < 100; i++ {
		syncMap.Store(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = syncMap.Load(i % 100)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapReadLarge тестирует производительность чтения из sync.Map с 100K элементами
func BenchmarkSyncMapReadLarge(b *testing.B) {
	var syncMap sync.Map

	// Подготовка данных
	for i := 0; i < 100000; i++ {
		syncMap.Store(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = syncMap.Load(i % 100000)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapReadHuge тестирует производительность чтения из sync.Map с 10M элементами
func BenchmarkSyncMapReadHuge(b *testing.B) {
	var syncMap sync.Map

	// Подготовка данных
	for i := 0; i < 10000000; i++ {
		syncMap.Store(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = syncMap.Load(i % 10000000)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapWrite тестирует производительность записи в sync.Map
func BenchmarkSyncMapWrite(b *testing.B) {
	var syncMap sync.Map

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap.Store(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapWriteLarge тестирует производительность записи в sync.Map с 100K элементами
func BenchmarkSyncMapWriteLarge(b *testing.B) {
	var syncMap sync.Map

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap.Store(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapWriteHuge тестирует производительность записи в sync.Map с 10M элементами
func BenchmarkSyncMapWriteHuge(b *testing.B) {
	var syncMap sync.Map

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap.Store(i, i)
		}(i)
	}
	wg.Wait()
}

// mutexMap представляет карту с мьютексом для конкурентного доступа
type mutexMap struct {
	data map[int]int
	mu   sync.RWMutex
}

func newMutexMap() *mutexMap {
	return &mutexMap{
		data: make(map[int]int),
	}
}

func (m *mutexMap) Set(key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *mutexMap) Get(key int) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

// BenchmarkMutexMapRead тестирует производительность чтения из карты с мьютексом
func BenchmarkMutexMapRead(b *testing.B) {
	mutexMap := newMutexMap()

	// Подготовка данных
	for i := 0; i < 100; i++ {
		mutexMap.Set(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = mutexMap.Get(i % 100)
		}(i)
	}
	wg.Wait()
}

// BenchmarkMutexMapReadLarge тестирует производительность чтения из карты с мьютексом с 100K элементами
func BenchmarkMutexMapReadLarge(b *testing.B) {
	mutexMap := newMutexMap()

	// Подготовка данных
	for i := 0; i < 100000; i++ {
		mutexMap.Set(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = mutexMap.Get(i % 100000)
		}(i)
	}
	wg.Wait()
}

// BenchmarkMutexMapReadHuge тестирует производительность чтения из карты с мьютексом с 10M элементами
func BenchmarkMutexMapReadHuge(b *testing.B) {
	mutexMap := newMutexMap()

	// Подготовка данных
	for i := 0; i < 10000000; i++ {
		mutexMap.Set(i, i)
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = mutexMap.Get(i % 10000000)
		}(i)
	}
	wg.Wait()
}

// BenchmarkMutexMapWrite тестирует производительность записи в карту с мьютексом
func BenchmarkMutexMapWrite(b *testing.B) {
	mutexMap := newMutexMap()

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutexMap.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkMutexMapWriteLarge тестирует производительность записи в карту с мьютексом с 100K элементами
func BenchmarkMutexMapWriteLarge(b *testing.B) {
	mutexMap := newMutexMap()

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutexMap.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkMutexMapWriteHuge тестирует производительность записи в карту с мьютексом с 10M элементами
func BenchmarkMutexMapWriteHuge(b *testing.B) {
	mutexMap := newMutexMap()

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutexMap.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// TestLRU тестирует работу LRU механизма
func TestLRU(t *testing.T) {
	// Создаем кэш с размером шарда 3, чтобы легко проверить LRU
	cache := NewCache[int, string](1, 3)

	// Добавляем 3 элемента
	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "three")

	// Проверяем, что все элементы есть в кэше
	if val, err := cache.Get(1); err != nil || val != "one" {
		t.Errorf("Expected 'one', got '%s', error: %v", val, err)
	}
	if val, err := cache.Get(2); err != nil || val != "two" {
		t.Errorf("Expected 'two', got '%s', error: %v", val, err)
	}
	if val, err := cache.Get(3); err != nil || val != "three" {
		t.Errorf("Expected 'three', got '%s', error: %v", val, err)
	}

	// Добавляем 4-й элемент, что должно вытеснить наименее недавно использованный (1)
	cache.Set(4, "four")

	// Проверяем, что первый элемент был вытеснен
	if _, err := cache.Get(1); err != ErrNotFound {
		t.Errorf("Expected ErrNotFound for key 1, got: %v", err)
	}

	// Проверяем, что остальные элементы на месте
	if val, err := cache.Get(2); err != nil || val != "two" {
		t.Errorf("Expected 'two', got '%s', error: %v", val, err)
	}
	if val, err := cache.Get(3); err != nil || val != "three" {
		t.Errorf("Expected 'three', got '%s', error: %v", val, err)
	}
	if val, err := cache.Get(4); err != nil || val != "four" {
		t.Errorf("Expected 'four', got '%s', error: %v", val, err)
	}

	// Получаем элемент 2, чтобы он стал самым недавно использованным
	if val, err := cache.Get(2); err != nil || val != "two" {
		t.Errorf("Expected 'two', got '%s', error: %v", val, err)
	}

	// Добавляем 5-й элемент, что должно вытеснить наименее недавно использованный (3)
	cache.Set(5, "five")

	// Проверяем, что элемент 3 был вытеснен
	if _, err := cache.Get(3); err != ErrNotFound {
		t.Errorf("Expected ErrNotFound for key 3, got: %v", err)
	}

	// Проверяем, что остальные элементы на месте
	if val, err := cache.Get(2); err != nil || val != "two" {
		t.Errorf("Expected 'two', got '%s', error: %v", val, err)
	}
	if val, err := cache.Get(4); err != nil || val != "four" {
		t.Errorf("Expected 'four', got '%s', error: %v", val, err)
	}
	if val, err := cache.Get(5); err != nil || val != "five" {
		t.Errorf("Expected 'five', got '%s', error: %v", val, err)
	}
}
