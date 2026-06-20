# 1647 — Minimum Deletions To Make Character Frequencies Unique

## Deskripsi

**Soal:** [1647. Minimum Deletions To Make Character Frequencies Unique](https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1647: Minimum Deletions to Make Character Frequencies Unique
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDeletions("aab"))
	fmt.Println(MinDeletions("aaabbbcc"))
	fmt.Println(MinDeletions("ceabaacb"))
}

func MinDeletions(s string) int {
	// Time: O(N), Space: O(1)
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

  // Membuat map untuk pencarian O(1): key → value
	used := make(map[int]bool)
	deletions := 0

	for _, f := range freq {
		for f > 0 && used[f] {
			f--
			deletions++
		}
		if f > 0 {
			used[f] = true
		}
	}

	return deletions
}
```
