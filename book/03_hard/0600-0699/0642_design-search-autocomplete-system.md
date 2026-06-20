# 0642 — Design Search Autocomplete System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func charIdx(c byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie, Prefix Sum

**Kompleksitas Waktu:** O(N * L) where N=#sentences, L=avg length  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #642: Design Search Autocomplete System
// https://leetcode.com/problems/design-search-autocomplete-system/
// Difficulty: Hard [Paid]
//
// Design an autocomplete system: given sentences and their occurrence counts,
// return the top 3 most relevant sentences as the user types character by character.
// Results sorted by frequency (desc), then lexicographically (asc).
// '#' signals end of current sentence, which is saved to the history.

// TrieNode represents a node in the trie.
type TrieNode struct {
	children [27]*TrieNode // a-z + ' ' mapped to 0-26
	times    int           // frequency of sentence ending at this node
}

// AutocompleteSystem implements the search autocomplete system.
type AutocompleteSystem struct {
	root       *TrieNode
	prefix     strings.Builder
	currNode   *TrieNode // current Trie position; nil if prefix not in trie
	currSent   strings.Builder
}

// charIdx maps 'a'-'z' -> 0-25 and ' ' -> 26.
func charIdx(c byte) int {
	if c == ' ' {
		return 26
	}
	return int(c - 'a')
}

// NewAutocompleteSystem creates the system with initial sentences and frequencies.
// Time: O(N * L) where N=#sentences, L=avg length
func NewAutocompleteSystem(sentences []string, times []int) *AutocompleteSystem {
	as := &AutocompleteSystem{
		root:     &TrieNode{},
		currNode: nil,
	}
	for i, s := range sentences {
		as.insert(s, times[i])
	}
	as.currNode = as.root
	return as
}

// insert adds a sentence into the trie with the given frequency.
func (as *AutocompleteSystem) insert(sentence string, times int) {
	node := as.root
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(sentence); i++ {
		idx := charIdx(sentence[i])
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.times += times
}

// traverseAndCollect collects all sentences under a node with their frequencies.
func (as *AutocompleteSystem) traverseAndCollect(node *TrieNode, prefix string, results *[]sentenceFreq) {
	if node == nil {
		return
	}
	if node.times > 0 {
		*results = append(*results, sentenceFreq{sentence: prefix, times: node.times})
	}
	for i := 0; i < 27; i++ {
		if node.children[i] != nil {
			var ch byte
			if i == 26 {
				ch = ' '
			} else {
				ch = byte('a' + i)
			}
			as.traverseAndCollect(node.children[i], prefix+string(ch), results)
		}
	}
}

type sentenceFreq struct {
	sentence string
	times    int
}

// Input processes a character typed by the user and returns top 3 autocomplete results.
// '#' marks the end of the current sentence.
func (as *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		// Save the current sentence.
		sentence := as.currSent.String()
		as.insert(sentence, 1)
		as.currSent.Reset()
		as.prefix.Reset()
		as.currNode = as.root
		return nil
	}

	as.currSent.WriteByte(c)
	as.prefix.WriteByte(c)

	if as.currNode == nil {
		return nil
	}

	idx := charIdx(c)
	as.currNode = as.currNode.children[idx]
	if as.currNode == nil {
		return nil
	}

	// Collect all sentences under current node.
	var candidates []sentenceFreq
	as.traverseAndCollect(as.currNode, as.prefix.String(), &candidates)

	// Sort by frequency desc, then lexicographically asc.
  // Custom sort dengan comparator
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].times != candidates[j].times {
			return candidates[i].times > candidates[j].times
		}
		return candidates[i].sentence < candidates[j].sentence
	})

	// Return top 3.
	top := make([]string, 0, 3)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(candidates) && i < 3; i++ {
		top = append(top, candidates[i].sentence)
	}
	return top
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0642 Design Search Autocomplete System ===")

	// Test 1: Basic autocomplete.
	sentences := []string{"i love you", "island", "ironman", "i love leetcode"}
	times := []int{5, 3, 2, 2}
	ac := NewAutocompleteSystem(sentences, times)

	fmt.Println("Input: 'i'")
	r1 := ac.Input('i')
	fmt.Printf("  Results: %v\n", r1)

	fmt.Println("Input: ' '")
	r2 := ac.Input(' ')
	fmt.Printf("  Results: %v\n", r2)

	fmt.Println("Input: 'a'")
	r3 := ac.Input('a')
	fmt.Printf("  Results: %v\n", r3)

	fmt.Println("Input: '#' (end sentence)")
	r4 := ac.Input('#')
	fmt.Printf("  Results: %v\n", r4)

	// After adding "i a", it should now be a candidate.
	fmt.Println("Input: 'i' (after adding 'i a')")
	r5 := ac.Input('i')
	fmt.Printf("  Results: %v\n", r5)

	fmt.Println("Input: ' '")
	r6 := ac.Input(' ')
	fmt.Printf("  Results: %v\n", r6)

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Empty system.
	ac2 := NewAutocompleteSystem(nil, nil)
	fmt.Println("Input 'a' in empty system:")
	r7 := ac2.Input('a')
	fmt.Printf("  Results: %v (expected [])\n", r7)
	ac2.Input('#')

	// Exact match then more typing (prefix not found).
	ac3 := NewAutocompleteSystem([]string{"hello"}, []int{10})
	ac3.Input('h')
	ac3.Input('e')
	ac3.Input('l')
	ac3.Input('l')
	ac3.Input('o')
	fmt.Println("After 'hello' then 'x' (no match):")
	r8 := ac3.Input('x')
	fmt.Printf("  Results: %v (expected [])\n", r8)

	// Multiple sentences with same frequency.
	ac4 := NewAutocompleteSystem(
		[]string{"abc", "abd", "abe"},
		[]int{1, 1, 1},
	)
	fmt.Println("Same frequency, should sort lexicographically:")
	r9 := ac4.Input('a')
	fmt.Printf("  Results: %v (expected [abc abd abe])\n", r9)
}
```
