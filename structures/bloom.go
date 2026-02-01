package structures

import (
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bits   []bool
	size   int
	k      int
	hashes []hashFunc
}

type hashFunc func(data []byte) uint32

func NewBloomFilter(n int, p float64) *BloomFilter {
	m := optimalSize(n, p)
	k := optimalHashFunctions(m, n)

	bf := &BloomFilter{
		bits:   make([]bool, m),
		size:   m,
		k:      k,
		hashes: make([]hashFunc, k),
	}

	for i := 0; i < k; i++ {
		seed := uint32(i)
		bf.hashes[i] = makeHashFunc(seed)
	}

	return bf
}

func makeHashFunc(seed uint32) hashFunc {
	return func(data []byte) uint32 {
		h := fnv.New32a()
		h.Write(data)
		return h.Sum32() + seed
	}
}

func NewBloomFilterWithParams(m, k int) *BloomFilter {
	bf := &BloomFilter{
		bits:   make([]bool, m),
		size:   m,
		k:      k,
		hashes: make([]hashFunc, k),
	}

	for i := 0; i < k; i++ {
		seed := uint32(i)
		bf.hashes[i] = makeHashFunc(seed)
	}

	return bf
}

func optimalSize(n int, p float64) int {
	return int(math.Ceil(-float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
}

func optimalHashFunctions(m, n int) int {
	return int(math.Ceil(float64(m) / float64(n) * math.Log(2)))
}

func (bf *BloomFilter) Add(data []byte) {
	for i := 0; i < bf.k; i++ {
		hash := bf.hashes[i](data)
		index := hash % uint32(bf.size)
		bf.bits[index] = true
	}
}

func (bf *BloomFilter) AddString(s string) {
	bf.Add([]byte(s))
}

func (bf *BloomFilter) Contains(data []byte) bool {
	for i := 0; i < bf.k; i++ {
		hash := bf.hashes[i](data)
		index := hash % uint32(bf.size)
		if !bf.bits[index] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) ContainsString(s string) bool {
	return bf.Contains([]byte(s))
}

func (bf *BloomFilter) FalsePositiveRate(n int) float64 {
	return math.Pow(1-math.Pow(1-1.0/float64(bf.size), float64(bf.k)*float64(n)), float64(bf.k))
}

func (bf *BloomFilter) Size() int {
	return bf.size
}

func (bf *BloomFilter) HashFunctions() int {
	return bf.k
}

func (bf *BloomFilter) BitsSet() int {
	count := 0
	for _, bit := range bf.bits {
		if bit {
			count++
		}
	}
	return count
}

func (bf *BloomFilter) EstimatedBitsSet(n int) float64 {
	return float64(bf.size) * (1 - math.Pow(1-1.0/float64(bf.size), float64(bf.k)*float64(n)))
}

func (bf *BloomFilter) Reset() {
	for i := range bf.bits {
		bf.bits[i] = false
	}
}

func (bf *BloomFilter) Union(other *BloomFilter) bool {
	if bf.size != other.size || bf.k != other.k {
		return false
	}

	for i := 0; i < bf.size; i++ {
		bf.bits[i] = bf.bits[i] || other.bits[i]
	}
	return true
}

func (bf *BloomFilter) Intersection(other *BloomFilter) bool {
	if bf.size != other.size || bf.k != other.k {
		return false
	}

	for i := 0; i < bf.size; i++ {
		bf.bits[i] = bf.bits[i] && other.bits[i]
	}
	return true
}
