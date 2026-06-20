# 2416 — Sum Of Prefix Scores Of Strings

## Deskripsi

**Soal:** [2416. Sum Of Prefix Scores Of Strings](https://leetcode.com/problems/sum-of-prefix-scores-of-strings/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Trie (pohon awalan)

## Solusi Go

```go
package main

// LeetCode #2416: Sum of Prefix Scores of Strings
// https://leetcode.com/problems/sum-of-prefix-scores-of-strings/
// Difficulty: Hard
//
// Trie with visit count. Insert each word, counting how many words pass through
// each prefix node. Then for each word, sum the counts along its prefix path.
// Time O(N * L) | Space O(N * L) where N = len(words), L = avg word length.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
	// Example 2
	fmt.Println(sumPrefixScores([]string{"abcd"}))
	// Edge: single char
	fmt.Println(sumPrefixScores([]string{"a", "a", "a"}))
}

type trieNode struct {
	children [26]*trieNode
	count    int
}

func sumPrefixScores(words []string) []int {
	root := &trieNode{}

	// Insert all words into trie, incrementing count at each node
	for _, w := range words {
		cur := root
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if cur.children[idx] == nil {
				cur.children[idx] = &trieNode{}
			}
			cur = cur.children[idx]
			cur.count++
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(words))
	for wi, w := range words {
		cur := root
		total := 0
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			cur = cur.children[idx]
			total += cur.count
		}
		ans[wi] = total
	}
	return ans
}
```
