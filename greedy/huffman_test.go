package greedy

import (
	"testing"
)

func TestBuildFrequencyMap(t *testing.T) {
	data := []byte("hello world")
	freq := BuildFrequencyMap(data)

	if freq['h'] != 1 {
		t.Errorf("Expected freq['h'] = 1, got %d", freq['h'])
	}
	if freq['l'] != 3 {
		t.Errorf("Expected freq['l'] = 3, got %d", freq['l'])
	}
	if freq['o'] != 2 {
		t.Errorf("Expected freq['o'] = 2, got %d", freq['o'])
	}
}

func TestHuffmanTreeBuild(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{
		'a': 5,
		'b': 9,
		'c': 12,
		'd': 13,
		'e': 16,
		'f': 45,
	}

	ht.Build(freq)

	codes := ht.GetCodes()
	if len(codes) != len(freq) {
		t.Errorf("Expected %d codes, got %d", len(freq), len(codes))
	}
}

func TestHuffmanTreeEncodeDecode(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{
		'a': 5,
		'b': 9,
		'c': 12,
		'd': 13,
		'e': 16,
		'f': 45,
	}

	ht.Build(freq)

	data := []byte("abcdef")
	encoded := ht.Encode(data)
	decoded := ht.Decode(encoded)

	if string(decoded) != string(data) {
		t.Errorf("Decoded data %s != original data %s", decoded, data)
	}
}

func TestHuffmanTreeEncodeString(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{
		'h': 1,
		'e': 1,
		'l': 3,
		'o': 2,
		' ': 1,
		'w': 1,
		'r': 1,
		'd': 1,
	}

	ht.Build(freq)

	data := "hello world"
	encoded := ht.EncodeString(data)
	decoded := ht.DecodeString(string(encoded))

	if string(decoded) != data {
		t.Errorf("Decoded data %s != original data %s", decoded, data)
	}
}

func TestHuffmanTreeGetCode(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{
		'a': 5,
		'b': 9,
		'c': 12,
	}

	ht.Build(freq)

	code, ok := ht.GetCode('a')
	if !ok {
		t.Errorf("Expected code for 'a'")
	}
	if code == "" {
		t.Errorf("Expected non-empty code for 'a'")
	}

	_, ok = ht.GetCode('z')
	if ok {
		t.Errorf("Expected no code for 'z'")
	}
}

func TestCompressDecompress(t *testing.T) {
	data := []byte("hello world hello world hello world")

	encoded, ht := Compress(data)
	decoded := Decompress(encoded, ht)

	if string(decoded) != string(data) {
		t.Errorf("Decoded data %s != original data %s", decoded, data)
	}
}

func TestCompressDecompressString(t *testing.T) {
	data := "hello world hello world hello world"

	encoded, ht := CompressString(data)
	decoded := DecompressString(string(encoded), ht)

	if string(decoded) != data {
		t.Errorf("Decoded data %s != original data %s", decoded, data)
	}
}

func TestCalculateCompressionRatio(t *testing.T) {
	data := []byte("aaaaaaaa")

	encoded, _ := Compress(data)
	ratio := CalculateCompressionRatio(data, encoded)

	if ratio > 1.0 {
		t.Errorf("Compression ratio %f should be <= 1.0 for highly repetitive data", ratio)
	}
}

func TestHuffmanTreeEmpty(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{}

	ht.Build(freq)

	data := []byte("")
	encoded := ht.Encode(data)
	decoded := ht.Decode(encoded)

	if len(decoded) != 0 {
		t.Errorf("Expected empty decoded data, got %v", decoded)
	}
}

func TestHuffmanTreeSingleChar(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{
		'a': 1,
	}

	ht.Build(freq)

	data := []byte("a")
	encoded := ht.Encode(data)
	decoded := ht.Decode(encoded)

	if string(decoded) != string(data) {
		t.Errorf("Decoded data %s != original data %s", decoded, data)
	}
}

func TestHuffmanTreeConsistency(t *testing.T) {
	data := []byte("this is a test for huffman coding")

	ht := NewHuffmanTree()
	freq := BuildFrequencyMap(data)
	ht.Build(freq)

	encoded := ht.Encode(data)
	decoded := ht.Decode(encoded)

	if string(decoded) != string(data) {
		t.Errorf("Decoded data %s != original data %s", decoded, data)
	}

	codes := ht.GetCodes()
	for char, code := range codes {
		if code == "" {
			t.Errorf("Code for char %c should not be empty", char)
		}
	}
}

func TestHuffmanTreeMultipleChars(t *testing.T) {
	ht := NewHuffmanTree()
	freq := make(map[byte]int)

	for i := 0; i < 26; i++ {
		freq[byte('a'+i)] = i + 1
	}

	ht.Build(freq)

	codes := ht.GetCodes()
	if len(codes) != 26 {
		t.Errorf("Expected 26 codes, got %d", len(codes))
	}
}

func TestHuffmanCodePrefixProperty(t *testing.T) {
	ht := NewHuffmanTree()
	freq := map[byte]int{
		'a': 5,
		'b': 9,
		'c': 12,
		'd': 13,
		'e': 16,
		'f': 45,
	}

	ht.Build(freq)

	codes := ht.GetCodes()

	for c1, code1 := range codes {
		for c2, code2 := range codes {
			if c1 != c2 {
				if len(code1) < len(code2) && code2[:len(code1)] == code1 {
					t.Errorf("Code for %c (%s) is prefix of code for %c (%s)", c1, code1, c2, code2)
				}
				if len(code2) < len(code1) && code1[:len(code2)] == code2 {
					t.Errorf("Code for %c (%s) is prefix of code for %c (%s)", c2, code2, c1, code1)
				}
			}
		}
	}
}
