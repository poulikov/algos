package greedy

import (
	"container/heap"
)

type HuffmanNode struct {
	char   byte
	freq   int
	left   *HuffmanNode
	right  *HuffmanNode
	isLeaf bool
	index  int
}

type HuffmanTree struct {
	root  *HuffmanNode
	codes map[byte]string
	freq  map[byte]int
}

type MinHeap []*HuffmanNode

func (mh MinHeap) Len() int { return len(mh) }

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].freq < mh[j].freq || (mh[i].freq == mh[j].freq && mh[i].char < mh[j].char)
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *MinHeap) Push(x interface{}) {
	n := len(*mh)
	item := x.(*HuffmanNode)
	item.index = n
	*mh = append(*mh, item)
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*mh = old[:n-1]
	return item
}

func NewHuffmanTree() *HuffmanTree {
	return &HuffmanTree{
		codes: make(map[byte]string),
		freq:  make(map[byte]int),
	}
}

func (ht *HuffmanTree) Build(freq map[byte]int) {
	ht.freq = freq

	minHeap := make(MinHeap, 0)
	heap.Init(&minHeap)

	for char, frequency := range freq {
		node := &HuffmanNode{
			char:   char,
			freq:   frequency,
			isLeaf: true,
		}
		heap.Push(&minHeap, node)
	}

	for minHeap.Len() > 1 {
		left := heap.Pop(&minHeap).(*HuffmanNode)
		right := heap.Pop(&minHeap).(*HuffmanNode)

		parent := &HuffmanNode{
			freq:   left.freq + right.freq,
			left:   left,
			right:  right,
			isLeaf: false,
		}

		heap.Push(&minHeap, parent)
	}

	if minHeap.Len() == 1 {
		ht.root = heap.Pop(&minHeap).(*HuffmanNode)
	}

	if ht.root == nil && len(freq) > 0 {
		for char, frequency := range freq {
			ht.root = &HuffmanNode{
				char:   char,
				freq:   frequency,
				isLeaf: true,
			}
			break
		}
	}

	ht.generateCodes(ht.root, "")
}

func (ht *HuffmanTree) generateCodes(node *HuffmanNode, code string) {
	if node == nil {
		return
	}

	if node.isLeaf {
		if code == "" {
			ht.codes[node.char] = "0"
		} else {
			ht.codes[node.char] = code
		}
		return
	}

	ht.generateCodes(node.left, code+"0")
	ht.generateCodes(node.right, code+"1")
}

func (ht *HuffmanTree) Encode(data []byte) []byte {
	encoded := ""
	for _, b := range data {
		code, ok := ht.codes[b]
		if !ok {
			code = ht.codes[b]
		}
		encoded += code
	}
	return []byte(encoded)
}

func (ht *HuffmanTree) EncodeString(s string) []byte {
	return ht.Encode([]byte(s))
}

func (ht *HuffmanTree) Decode(encoded []byte) []byte {
	var result []byte
	current := ht.root

	if ht.root == nil {
		return result
	}

	if ht.root.isLeaf && len(ht.codes) == 1 {
		for range encoded {
			result = append(result, ht.root.char)
		}
		return result
	}

	for _, bit := range encoded {
		if current == nil {
			break
		}

		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if current != nil && current.isLeaf {
			result = append(result, current.char)
			current = ht.root
		}
	}

	return result
}

func (ht *HuffmanTree) DecodeString(encoded string) []byte {
	return ht.Decode([]byte(encoded))
}

func (ht *HuffmanTree) GetCode(char byte) (string, bool) {
	code, ok := ht.codes[char]
	return code, ok
}

func (ht *HuffmanTree) GetCodes() map[byte]string {
	return ht.codes
}

func (ht *HuffmanTree) GetFreq() map[byte]int {
	return ht.freq
}

func BuildFrequencyMap(data []byte) map[byte]int {
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}
	return freq
}

func BuildFrequencyMapString(s string) map[byte]int {
	return BuildFrequencyMap([]byte(s))
}

func Compress(data []byte) ([]byte, *HuffmanTree) {
	freq := BuildFrequencyMap(data)
	ht := NewHuffmanTree()
	ht.Build(freq)
	encoded := ht.Encode(data)
	return encoded, ht
}

func CompressString(s string) ([]byte, *HuffmanTree) {
	return Compress([]byte(s))
}

func Decompress(encoded []byte, ht *HuffmanTree) []byte {
	if ht == nil || len(encoded) == 0 {
		return nil
	}
	return ht.Decode(encoded)
}

func DecompressString(encoded string, ht *HuffmanTree) []byte {
	return Decompress([]byte(encoded), ht)
}

func CalculateCompressionRatio(original, compressed []byte) float64 {
	if len(original) == 0 {
		return 0
	}
	return float64(len(compressed)) / float64(len(original)*8)
}
