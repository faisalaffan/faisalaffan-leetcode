# 3093 — Longest Common Suffix Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func stringIndices(wordsContainer []string, wordsQuery []string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Trie, Prefix Sum

**Waktu:** O((N + Q) * M)  |  **Ruang:** O(N * M)

> 🎓 **Fresh Grad Tips:** Kuasai **Trie** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3093: Longest Common Suffix Queries
// https://leetcode.com/problems/longest-common-suffix-queries/
// Difficulty: Hard
// Time: O((N + Q) * M) | Space: O(N * M)
//
// Approach: Trie over reversed words
// Insert each wordContainer (reversed) into a trie. At each node, store
// the index and length of the shortest container word that shares this suffix.
// For each query, traverse the reversed query in the trie. The last reachable
// node gives the shortest container word matching the longest possible suffix.
// If no suffix matches, fall back to the root (shortest overall word).

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	length   int // shortest container word length passing through this node
	idx      int // index of that shortest word
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &TrieNode{length: 1 << 30, idx: -1}

	// Insert each container word into trie (reversed)
	for i, w := range wordsContainer {
		node := root
		// Update root with shortest word
		if len(w) < node.length {
			node.length = len(w)
			node.idx = i
		}
		// Traverse reversed word
		for k := len(w) - 1; k >= 0; k-- {
			c := int(w[k] - 'a')
			if node.children[c] == nil {
				node.children[c] = &TrieNode{length: 1 << 30, idx: -1}
			}
			node = node.children[c]
			if len(w) < node.length {
				node.length = len(w)
				node.idx = i
			}
		}
	}

	// Query each word
  // Alokasi slice
	ans := make([]int, len(wordsQuery))
	for qi, q := range wordsQuery {
		node := root
		for k := len(q) - 1; k >= 0; k-- {
			c := int(q[k] - 'a')
			if node.children[c] == nil {
				break
			}
			node = node.children[c]
		}
		ans[qi] = node.idx
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println("Test 1:", stringIndices(
		[]string{"abcd", "bcd", "xbcd"},
		[]string{"bcd", "ab", "cd"},
	))
	// Expected: [1, 1, 1]
	// "bcd" is shortest container sharing suffix "bcd" → idx 1
	// "ab" suffix "b" matches "bcd" suffix "b" → idx 1
	// "cd" suffix "cd" matches "bcd" suffix "cd" → idx 1

	// Example 2
	fmt.Println("Test 2:", stringIndices(
		[]string{"abcd", "bcd", "xbcd"},
		[]string{"cd", "bcd", "xyz"},
	))
	// Expected: [1, 1, 1]
	// "xyz" has no suffix match, fallback to root (shortest overall = "bcd" len 3 vs "abcd" len 4)

	// Example 3
	fmt.Println("Test 3:", stringIndices(
		[]string{"aaaa", "a"},
		[]string{"a"},
	))
	// Expected: [1]
	// "a" suffix "a": both match. Shortest = "a" (len 1) at idx 1

	// Single container
	fmt.Println("Test 4:", stringIndices(
		[]string{"hello"},
		[]string{"lo", "o", "xyz"},
	))
	// "lo" suffix "lo": matches "hello" → idx 0 (only option anyway)
	// "o" suffix "o": matches "hello" → idx 0
	// "xyz": no match, fallback to root → idx 0

	// Multiple containers, query matching none
	fmt.Println("Test 5:", stringIndices(
		[]string{"cat", "bat", "rat"},
		[]string{"dog"},
	))
	// "dog": no suffix match, fallback to root (shortest = "cat" or "bat" or "rat", all len 3, first = idx 0)

	// Prefix longer than container
	fmt.Println("Test 6:", stringIndices(
		[]string{"ab", "a"},
		[]string{"abcde"},
	))
	// "abcde" suffix "abde": matches "ab" first then runs out
	// Last reachable node is the one for "ab" suffix → idx 0 (len 2)
	// Compare: root has "a" (len 1), but we don't stop at root
	// Traverse: 'e','d','c','b','a' → 'a' exists (from "ab" reversed "ba")
	// Wait, "ab" reversed is "ba". Trie: 'b', 'a'.
	// Query "abcde" reversed: 'e','d','c','b','a'
	// 'e': no child → break. Node is root → idx -1 → but we expect 0.
	// Actually root has idx=0 because shortest word "a" has length 1 < root's initial large length.
	// Wait: "ab" len 2, "a" len 1. Root: first idx=0 (shortest so far "ab" len 2), then idx=1 (shortest "a" len 1).
	// So root.idx=1.
	// Query "abcde" reversed: 'e' has no child. Break. Return root.idx = 1.
	// Expected: 1 (container "a")

	// Same prefix, different lengths
	fmt.Println("Test 7:", stringIndices(
		[]string{"abcdef", "abc"},
		[]string{"def"},
	))
	// "def" reversed: 'f','e','d'
	// Trie has "abcdef" reversed: 'f','e','d','c','b','a'
	// 'f' → matches "abcdef". Node stores len 6.
	// 'e' → matches. Node stores len 6 (still, "abcdef" is only one).
	// 'd' → matches. Node stores len 6.
	// No more query chars. Current node has idx=0 (abcdef).
	// Answer: 0
}
```
