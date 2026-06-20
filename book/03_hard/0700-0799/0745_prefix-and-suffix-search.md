# 0745 — Prefix And Suffix Search

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ConstructorWordFilter(words []string) WordFilter`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Trie, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Trie** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #745: Prefix and Suffix Search
// https://leetcode.com/problems/prefix-and-suffix-search/
// Difficulty: Hard
//
// Algorithm: Trie with "suffix#prefix" key
// For each word at index i, insert all possible "suffix#prefix" combinations
// into a trie. Each node stores the maximum weight (index) seen.
// f(prefix, suffix) = search for "suffix#prefix" in the trie.

import (
	"fmt"
)

// WordFilter provides the f(prefix, suffix) query
type WordFilter struct {
	root *TrieNode
}

type TrieNode struct {
	children [27]*TrieNode // 26 letters + '#'
	weight   int
}

func ConstructorWordFilter(words []string) WordFilter {
	wf := WordFilter{root: &TrieNode{weight: -1}}
	for weight, word := range words {
		// Insert all suffix#prefix combinations
		// For each suffix length s (0..len(word)):
		//   suffix = word[len(word)-s:]
		//   key = suffix + "#" + word
		n := len(word)
		for s := 0; s <= n; s++ {
			suffix := word[n-s:]
			key := suffix + "#" + word
			wf.insert(key, weight)
		}
	}
	return wf
}

func (this *WordFilter) insert(key string, weight int) {
	node := this.root
	for _, ch := range key {
		idx := charIndex(byte(ch))
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{weight: -1}
		}
		node = node.children[idx]
		if weight > node.weight {
			node.weight = weight
		}
	}
}

func charIndex(b byte) int {
	if b == '#' {
		return 26
	}
	return int(b - 'a')
}

func (this *WordFilter) F(pref string, suff string) int {
	key := suff + "#" + pref
	node := this.root
	for _, ch := range key {
		idx := charIndex(byte(ch))
		if node.children[idx] == nil {
			return -1
		}
		node = node.children[idx]
	}
	return node.weight
}

func main() {
	// Example from problem
	words := []string{"apple"}
	wf := ConstructorWordFilter(words)
	r1 := wf.F("a", "e")
	fmt.Printf("f(%q, %q) = %d (expected: 0)\n", "a", "e", r1)

	r2 := wf.F("b", "")
	fmt.Printf("f(%q, %q) = %d (expected: -1)\n\n", "b", "", r2)

	// Test case 2: multiple words, check highest index
	words2 := []string{"abc", "abcd", "abcde"}
	wf2 := ConstructorWordFilter(words2)
	r3 := wf2.F("a", "e")
	fmt.Printf("f(%q, %q) = %d (expected: 2)\n", "a", "e", r3)

	r4 := wf2.F("ab", "cd")
	fmt.Printf("f(%q, %q) = %d (expected: 1)\n", "ab", "cd", r4)

	r5 := wf2.F("abc", "abc")
	fmt.Printf("f(%q, %q) = %d (expected: 0)\n\n", "abc", "abc", r5)

	// Test case 3: empty prefix
	r6 := wf2.F("", "e")
	fmt.Printf("f(%q, %q) = %d (expected: 2)\n", "", "e", r6)

	// Test case 4: empty suffix
	r7 := wf2.F("a", "")
	fmt.Printf("f(%q, %q) = %d (expected: 2)\n\n", "a", "", r7)

	// Test case 5: no match
	r8 := wf2.F("xyz", "abc")
	fmt.Printf("f(%q, %q) = %d (expected: -1)\n\n", "xyz", "abc", r8)

	// Test case 6: same prefix and suffix
	words3 := []string{"ab", "aab", "aaab"}
	wf3 := ConstructorWordFilter(words3)
	r9 := wf3.F("a", "b")
	fmt.Printf("f(%q, %q) = %d (expected: 2)\n", "a", "b", r9)

	// Test case 7: single characters
	words4 := []string{"a", "aa", "aaa"}
	wf4 := ConstructorWordFilter(words4)
	r10 := wf4.F("aa", "a")
	fmt.Printf("f(%q, %q) = %d\n", "aa", "a", r10)
}
```
