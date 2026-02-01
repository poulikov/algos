package trees

import (
	"fmt"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	if trie == nil {
		t.Fatal("NewTrie() returned nil")
	}
	if !trie.IsEmpty() {
		t.Fatal("NewTrie() trie should be empty")
	}
	if trie.Size() != 0 {
		t.Fatalf("NewTrie() trie should have size 0, got %d", trie.Size())
	}
}

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	if trie.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", trie.Size())
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("helium")

	tests := []struct {
		word     string
		expected bool
	}{
		{"hello", true},
		{"world", true},
		{"helium", true},
		{"hell", false},
		{"helloworld", false},
		{"", false},
	}

	for _, test := range tests {
		if result := trie.Search(test.word); result != test.expected {
			t.Fatalf("Search(%q) = %v, expected %v", test.word, result, test.expected)
		}
	}
}

func TestTrieSearchEmpty(t *testing.T) {
	trie := NewTrie()
	if trie.Search("hello") {
		t.Fatal("Empty trie should not contain any words")
	}
}

func TestTrieStartsWith(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")
	trie.Insert("world")

	tests := []struct {
		prefix   string
		expected bool
	}{
		{"hel", true},
		{"he", true},
		{"hello", true},
		{"wor", true},
		{"xyz", false},
		{"", true},
	}

	for _, test := range tests {
		if result := trie.StartsWith(test.prefix); result != test.expected {
			t.Fatalf("StartsWith(%q) = %v, expected %v", test.prefix, result, test.expected)
		}
	}
}

func TestTrieDelete(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")

	if !trie.Delete("hello") {
		t.Fatal("Delete() should return true")
	}
	if trie.Search("hello") {
		t.Fatal("Word 'hello' still exists after deletion")
	}
	if trie.Size() != 1 {
		t.Fatalf("Expected size 1 after deletion, got %d", trie.Size())
	}
}

func TestTrieDeleteNonExistent(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")

	if trie.Delete("world") {
		t.Fatal("Delete() should return false for non-existent word")
	}
	if trie.Size() != 1 {
		t.Fatalf("Size should remain 1, got %d", trie.Size())
	}
}

func TestTrieDeleteWithSharedPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")
	trie.Insert("help")

	trie.Delete("helium")
	if trie.Search("helium") {
		t.Fatal("Word 'helium' still exists")
	}
	if !trie.Search("hello") {
		t.Fatal("Word 'hello' should still exist")
	}
	if !trie.Search("help") {
		t.Fatal("Word 'help' should still exist")
	}
}

func TestTrieDeletePrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("hell")

	trie.Delete("hell")
	if trie.Search("hell") {
		t.Fatal("Word 'hell' still exists")
	}
	if !trie.Search("hello") {
		t.Fatal("Word 'hello' should still exist")
	}
}

func TestTrieCountWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")
	trie.Insert("help")
	trie.Insert("world")

	tests := []struct {
		prefix   string
		expected int
	}{
		{"hel", 3},
		{"he", 3},
		{"hell", 1},
		{"w", 1},
		{"xyz", 0},
		{"", 4},
	}

	for _, test := range tests {
		if result := trie.CountWordsWithPrefix(test.prefix); result != test.expected {
			t.Fatalf("CountWordsWithPrefix(%q) = %d, expected %d", test.prefix, result, test.expected)
		}
	}
}

func TestTrieGetAllWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")
	trie.Insert("help")
	trie.Insert("world")

	words := trie.GetAllWordsWithPrefix("hel")
	if len(words) != 3 {
		t.Fatalf("Expected 3 words, got %d", len(words))
	}

	expectedWords := map[string]bool{
		"hello":  true,
		"helium": true,
		"help":   true,
	}

	for _, word := range words {
		if !expectedWords[word] {
			t.Fatalf("Unexpected word: %s", word)
		}
	}
}

func TestTrieGetAllWordsWithPrefixEmpty(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")

	words := trie.GetAllWordsWithPrefix("xyz")
	if len(words) != 0 {
		t.Fatalf("Expected 0 words, got %d", len(words))
	}
}

func TestTrieGetAllWords(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("go")

	words := trie.GetAllWords()
	if len(words) != 3 {
		t.Fatalf("Expected 3 words, got %d", len(words))
	}

	expectedWords := map[string]bool{
		"hello": true,
		"world": true,
		"go":    true,
	}

	for _, word := range words {
		if !expectedWords[word] {
			t.Fatalf("Unexpected word: %s", word)
		}
	}
}

func TestTrieLongestCommonPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")
	trie.Insert("help")

	lcp := trie.LongestCommonPrefix()
	if lcp != "hel" {
		t.Fatalf("Expected LCP 'hel', got '%s'", lcp)
	}
}

func TestTrieLongestCommonPrefixEmpty(t *testing.T) {
	trie := NewTrie()

	lcp := trie.LongestCommonPrefix()
	if lcp != "" {
		t.Fatalf("Expected empty LCP, got '%s'", lcp)
	}
}

func TestTrieLongestCommonPrefixSingleWord(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")

	lcp := trie.LongestCommonPrefix()
	if lcp != "hello" {
		t.Fatalf("Expected 'hello' as LCP for single word, got '%s'", lcp)
	}
}

func TestTrieLongestCommonPrefixNoCommonPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	lcp := trie.LongestCommonPrefix()
	if lcp != "" {
		t.Fatalf("Expected empty LCP, got '%s'", lcp)
	}
}

func TestTrieIsEmpty(t *testing.T) {
	trie := NewTrie()
	if !trie.IsEmpty() {
		t.Fatal("New trie should be empty")
	}

	trie.Insert("hello")
	if trie.IsEmpty() {
		t.Fatal("Trie with words should not be empty")
	}
}

func TestTrieSize(t *testing.T) {
	trie := NewTrie()
	if trie.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", trie.Size())
	}

	trie.Insert("hello")
	if trie.Size() != 1 {
		t.Fatalf("Expected size 1, got %d", trie.Size())
	}

	trie.Insert("world")
	if trie.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", trie.Size())
	}

	trie.Insert("hello")
	if trie.Size() != 2 {
		t.Fatalf("Expected size 2 after duplicate insert, got %d", trie.Size())
	}
}

func TestTrieClear(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	trie.Clear()

	if !trie.IsEmpty() {
		t.Fatal("Clear() should make trie empty")
	}
	if trie.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", trie.Size())
	}
	if trie.Search("hello") {
		t.Fatal("Clear() should remove all words")
	}
}

func TestTrieCopy(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	trieCopy := trie.Copy()

	if trieCopy.Size() != trie.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", trieCopy.Size(), trie.Size())
	}

	trie.Insert("go")
	trie.Delete("hello")

	if trieCopy.Search("go") {
		t.Fatal("Copy should be independent, but found new word in copy")
	}
	if !trieCopy.Search("hello") {
		t.Fatal("Copy should be independent, but missing original word")
	}
}

func TestTrieString(t *testing.T) {
	trie := NewTrie()
	str := trie.String()
	if str != "[]" {
		t.Fatalf("Expected '[]', got '%s'", str)
	}

	trie.Insert("hello")
	str = trie.String()
	if str == "[]" {
		t.Fatal("String should not be '[]' after insert")
	}
}

func TestTrieHasWord(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")

	if !trie.HasWord("hello") {
		t.Fatal("HasWord should return true")
	}
	if trie.HasWord("world") {
		t.Fatal("HasWord should return false for non-existent word")
	}
}

func TestTrieGetNode(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")

	node := trie.GetNode("hello")
	if node == nil {
		t.Fatal("GetNode should return node for existing word")
	}

	node = trie.GetNode("world")
	if node != nil {
		t.Fatal("GetNode should return nil for non-existent word")
	}
}

func TestTrieGetChildren(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("helium")
	trie.Insert("world")

	children := trie.GetChildren("he")
	if len(children) != 1 {
		t.Fatalf("Expected 1 child for prefix 'he', got %d", len(children))
	}
	if children[0] != 'l' {
		t.Fatalf("Expected child 'l', got '%c'", children[0])
	}

	children = trie.GetChildren("hel")
	if len(children) != 2 {
		t.Fatalf("Expected 2 children for prefix 'hel', got %d", len(children))
	}

	children = trie.GetChildren("xyz")
	if len(children) != 0 {
		t.Fatalf("Expected 0 children for non-existent prefix, got %d", len(children))
	}
}

func TestTrieUnicode(t *testing.T) {
	trie := NewTrie()
	trie.Insert("привет")
	trie.Insert("мир")
	trie.Insert("приветствие")

	if !trie.Search("привет") {
		t.Fatal("Trie should support Unicode")
	}
	if !trie.Search("мир") {
		t.Fatal("Trie should support Unicode")
	}

	prefixWords := trie.GetAllWordsWithPrefix("при")
	if len(prefixWords) != 2 {
		t.Fatalf("Expected 2 words with prefix 'при', got %d", len(prefixWords))
	}
}

func TestTrieSpecialCharacters(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hello-world")
	trie.Insert("hello_world")

	if !trie.Search("hello-world") {
		t.Fatal("Trie should handle special characters")
	}
	if !trie.Search("hello_world") {
		t.Fatal("Trie should handle special characters")
	}
}

func BenchmarkTrieInsert(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Insert(words[i%len(words)])
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	for _, word := range words {
		trie.Insert(word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Search(words[i%len(words)])
	}
}

func BenchmarkTrieGetAllWordsWithPrefix(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	for _, word := range words {
		trie.Insert(word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.GetAllWordsWithPrefix(words[i%len(words)][:3])
	}
}

func generateWords(n int) []string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = fmt.Sprintf("word%d", i)
	}
	return words
}
