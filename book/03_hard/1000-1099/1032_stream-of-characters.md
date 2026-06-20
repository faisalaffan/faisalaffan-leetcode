# 1032 — Stream Of Characters

## Deskripsi

**Soal:** [1032. Stream Of Characters](https://leetcode.com/problems/stream-of-characters/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Trie (pohon awalan)

> **Ide Kunci:** Trie of reversed words.

## Solusi Go

```go
package main

// LeetCode #1032: Stream of Characters
// https://leetcode.com/problems/stream-of-characters/
// Difficulty: Hard
//
// Approach: Trie of reversed words.
//   Build a trie from the reversed version of each word.
//   On each query, accumulate characters and walk the trie in reverse order.
//   If we ever reach a node that marks the end of a word, return true.

import "fmt"

func main() {
	// Example: StreamChecker({"cd","f","kl"})
	// queries: a,b,c,d,e,f,g,h,i,j,k,l -> all false until d (true), f (true), k (false), l (true)
	sc := Constructor([]string{"cd", "f", "kl"})
	fmt.Println(sc.Query('a')) // false
	fmt.Println(sc.Query('b')) // false
	fmt.Println(sc.Query('c')) // false
	fmt.Println(sc.Query('d')) // true
	fmt.Println(sc.Query('e')) // false
	fmt.Println(sc.Query('f')) // true
	fmt.Println(sc.Query('g')) // false
	fmt.Println(sc.Query('h')) // false
	fmt.Println(sc.Query('i')) // false
	fmt.Println(sc.Query('j')) // false
	fmt.Println(sc.Query('k')) // false
	fmt.Println(sc.Query('l')) // true

	fmt.Println("---")

	sc2 := Constructor([]string{"ab", "ba", "aaab", "abab", "baa"})
	fmt.Println(sc2.Query('a')) // false
	fmt.Println(sc2.Query('a')) // false
	fmt.Println(sc2.Query('a')) // false
	fmt.Println(sc2.Query('a')) // false
	fmt.Println(sc2.Query('a')) // false
	fmt.Println(sc2.Query('b')) // true
}

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type StreamChecker struct {
	root *TrieNode
	buf  []byte
}

func Constructor(words []string) StreamChecker {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		// Insert reversed word
		for i := len(w) - 1; i >= 0; i-- {
			idx := w[i] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.isEnd = true
	}
	return StreamChecker{root: root, buf: make([]byte, 0)}
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.buf = append(sc.buf, letter)
	node := sc.root
	// Walk the trie from the end of the buffer
	for i := len(sc.buf) - 1; i >= 0; i-- {
		idx := sc.buf[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
		if node.isEnd {
			return true
		}
	}
	return false
}
```
