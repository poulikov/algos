package trees

import (
	"fmt"
	"strings"
)

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	value    rune
}

// Trie represents a prefix tree (trie) for efficient string operations
type Trie struct {
	root *TrieNode
	size int
}

// NewTrie creates a new empty Trie
// Time complexity: O(1)
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a word to the trie
// Time complexity: O(m) where m is the length of the word
func (t *Trie) Insert(word string) {
	node := t.root

	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
				value:    ch,
			}
		}
		node = node.children[ch]
	}

	if !node.isEnd {
		node.isEnd = true
		t.size++
	}
}

// Search checks if a word exists in the trie
// Time complexity: O(m) where m is the length of the word
func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEnd
}

// StartsWith checks if there is any word in the trie that starts with the given prefix
// Time complexity: O(m) where m is the length of the prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

func (t *Trie) searchNode(prefix string) *TrieNode {
	node := t.root

	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return nil
		}
		node = node.children[ch]
	}

	return node
}

// Delete removes a word from the trie
// Time complexity: O(m) where m is the length of the word
func (t *Trie) Delete(word string) bool {
	if !t.Search(word) {
		return false
	}

	t.deleteHelper(t.root, word, 0)
	t.size--
	return true
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		node.isEnd = false
		return len(node.children) == 0
	}

	ch := rune(word[index])
	child, exists := node.children[ch]
	if !exists {
		return false
	}

	shouldDeleteChild := t.deleteHelper(child, word, index+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		return len(node.children) == 0 && !node.isEnd
	}

	return false
}

// CountWordsWithPrefix returns the number of words in the trie with the given prefix
// Time complexity: O(n) where n is the number of nodes in the subtree
func (t *Trie) CountWordsWithPrefix(prefix string) int {
	node := t.searchNode(prefix)
	if node == nil {
		return 0
	}
	return t.countWords(node)
}

func (t *Trie) countWords(node *TrieNode) int {
	count := 0
	if node.isEnd {
		count++
	}

	for _, child := range node.children {
		count += t.countWords(child)
	}

	return count
}

// GetAllWordsWithPrefix returns all words in the trie with the given prefix
// Time complexity: O(n * m) where n is the number of words and m is the average word length
func (t *Trie) GetAllWordsWithPrefix(prefix string) []string {
	node := t.searchNode(prefix)
	if node == nil {
		return []string{}
	}

	words := make([]string, 0)
	t.collectWords(node, prefix, &words)
	return words
}

func (t *Trie) collectWords(node *TrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	for ch, child := range node.children {
		t.collectWords(child, prefix+string(ch), words)
	}
}

// GetAllWords returns all words stored in the trie
// Time complexity: O(n * m) where n is the number of words and m is the average word length
func (t *Trie) GetAllWords() []string {
	words := make([]string, 0)
	t.collectWords(t.root, "", &words)
	return words
}

// LongestCommonPrefix returns the longest common prefix among all words in the trie
// Time complexity: O(m) where m is the length of the LCP
func (t *Trie) LongestCommonPrefix() string {
	if t.IsEmpty() {
		return ""
	}

	node := t.root
	var sb strings.Builder

	for {
		if node.isEnd || len(node.children) != 1 {
			break
		}

		for ch, child := range node.children {
			sb.WriteRune(ch)
			node = child
			break
		}
	}

	return sb.String()
}

// IsEmpty returns true if the trie contains no words
// Time complexity: O(1)
func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

// Size returns the number of words in the trie
// Time complexity: O(1)
func (t *Trie) Size() int {
	return t.size
}

// Clear removes all words from the trie
// Time complexity: O(1)
func (t *Trie) Clear() {
	t.root = &TrieNode{
		children: make(map[rune]*TrieNode),
	}
	t.size = 0
}

// Copy creates a new trie with the same words
// Time complexity: O(n * m) where n is the number of words and m is the average word length
func (t *Trie) Copy() *Trie {
	newTrie := NewTrie()

	for _, word := range t.GetAllWords() {
		newTrie.Insert(word)
	}

	return newTrie
}

// String returns a string representation of the trie
func (t *Trie) String() string {
	words := t.GetAllWords()
	if len(words) == 0 {
		return "[]"
	}
	return fmt.Sprintf("%v", words)
}

// HasWord checks if a word exists (alias for Search)
// Time complexity: O(m) where m is the length of the word
func (t *Trie) HasWord(word string) bool {
	return t.Search(word)
}

// GetNode returns the node corresponding to the given word
// This is useful for custom traversals
// Time complexity: O(m) where m is the length of the word
func (t *Trie) GetNode(word string) *TrieNode {
	return t.searchNode(word)
}

// GetChildren returns all children of the node corresponding to the given prefix
// Time complexity: O(k) where k is the number of children
func (t *Trie) GetChildren(prefix string) []rune {
	node := t.searchNode(prefix)
	if node == nil {
		return []rune{}
	}

	children := make([]rune, 0, len(node.children))
	for ch := range node.children {
		children = append(children, ch)
	}

	return children
}
