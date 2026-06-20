# 1065 — Index Pairs Of A String

## Deskripsi

**Soal:** [1065. Index Pairs Of A String](https://leetcode.com/problems/index-pairs-of-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2 + m*k)  
**Kompleksitas Ruang:** O(m*k) for trie

**Algoritma:** Trie (pohon awalan)

## Solusi Go

```go
package main

// LeetCode #1065: Index Pairs of a String
// https://leetcode.com/problems/index-pairs-of-a-string/
// Difficulty: Easy [Paid]
// Time: O(n^2 + m*k) | Space: O(m*k) for trie

import "fmt"

func main() {
	fmt.Println(indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	// [[1,5],[3,7],[10,13],[10,18]]
	fmt.Println(indexPairs("ababa", []string{"aba", "ab"}))
	// [[0,1],[0,2],[2,3],[2,4]]
}

// LeetCode submission: indexPairs
func indexPairs(text string, words []string) [][]int {
  // Membuat map untuk pencarian O(1): key → value
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}
	var ans [][]int
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(text); i++ {
		for j := i; j < len(text); j++ {
			if wordSet[text[i:j+1]] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
```
