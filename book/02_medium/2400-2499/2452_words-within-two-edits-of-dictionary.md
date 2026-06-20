# 2452 — Words Within Two Edits Of Dictionary

## Deskripsi

**Soal:** [2452. Words Within Two Edits Of Dictionary](https://leetcode.com/problems/words-within-two-edits-of-dictionary/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m * L)  
**Kompleksitas Ruang:** O(1) where L = word length

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2452: Words Within Two Edits of Dictionary
// https://leetcode.com/problems/words-within-two-edits-of-dictionary/
// Difficulty: Medium
// Time: O(n * m * L) | Space: O(1) where L = word length
// For each query word, check Hamming distance to each dictionary word.

import "fmt"

func main() {
	fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"})) // ["word","note","wood"]
	fmt.Println(twoEditWords([]string{"yes"}, []string{"not"}))                                            // []
}

func twoEditWords(queries []string, dictionary []string) []string {
  // Membuat slice untuk menyimpan hasil
	ans := make([]string, 0)
	for _, q := range queries {
		for _, d := range dictionary {
			if hammingDist(q, d) <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}

func hammingDist(a, b string) int {
	dist := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
```
