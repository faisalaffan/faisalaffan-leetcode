# 1938 — Maximum Genetic Difference Query

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func newTrie() *Trie`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Trie

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1938: Maximum Genetic Difference Query
// https://leetcode.com/problems/maximum-genetic-difference-query/
// Difficulty: Hard
// Trie + DFS offline. Binary trie stores node values along current path.
// Answer each query by finding max XOR with any ancestor value.

import "fmt"

const bitLen = 18 // 2^18 = 262144, covers values up to 200k

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

func (t *Trie) insert(x int) {
	node := t.root
	for b := bitLen; b >= 0; b-- {
		bit := (x >> b) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

func (t *Trie) remove(x int) {
	node := t.root
	for b := bitLen; b >= 0; b-- {
		bit := (x >> b) & 1
		child := node.children[bit]
		child.count--
		if child.count == 0 {
			node.children[bit] = nil
			return
		}
		node = child
	}
}

func (t *Trie) maxXor(x int) int {
	node := t.root
	ans := 0
	for b := bitLen; b >= 0; b-- {
		bit := (x >> b) & 1
		want := 1 - bit
		if node.children[want] != nil {
			ans |= (1 << b)
			node = node.children[want]
		} else if node.children[bit] != nil {
			node = node.children[bit]
		} else {
			break
		}
	}
	return ans
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n := len(parents)
	// Build adjacency and find root
  // Matriks 2D
	children := make([][]int, n)
	var root int
	for i, p := range parents {
		if p == -1 {
			root = i
		} else {
			children[p] = append(children[p], i)
		}
	}

	// Group queries by node
  // Matriks 2D
	qByNode := make([][][2]int, n)
	for i, q := range queries {
		node, val := q[0], q[1]
		qByNode[node] = append(qByNode[node], [2]int{val, i})
	}

  // Alokasi slice
	ans := make([]int, len(queries))
	trie := newTrie()

	var dfs func(u int)
	dfs = func(u int) {
		trie.insert(u)
		for _, q := range qByNode[u] {
			val, idx := q[0], q[1]
			ans[idx] = trie.maxXor(val)
		}
		for _, v := range children[u] {
			dfs(v)
		}
		trie.remove(u)
	}

	dfs(root)
	return ans
}

func main() {
	// Example 1
	parents := []int{-1, 0, 1, 1}
	queries := [][]int{{0, 2}, {3, 2}, {2, 5}}
	result := maxGeneticDifference(parents, queries)
	fmt.Println(result) // Expected: [2, 3, 7]

	// Example 2
	parents2 := []int{-1, 0, 0, 1, 1, 2, 2}
	queries2 := [][]int{{0, 10}, {6, 7}}
	result2 := maxGeneticDifference(parents2, queries2)
	fmt.Println(result2)
}
```
