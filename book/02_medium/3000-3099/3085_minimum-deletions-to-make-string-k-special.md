# 3085 — Minimum Deletions To Make String K Special

## Deskripsi

**Soal:** [3085. Minimum Deletions To Make String K Special](https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 26) = O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumDeletions(word string, k int) int`

## Solusi Go

```go
package main

// LeetCode #3085: Minimum Deletions to Make String K-Special
// https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/
// Difficulty: Medium
// Time: O(n * 26) = O(n) | Space: O(26) = O(1)

import "fmt"

func minimumDeletions(word string, k int) int {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
	for _, ch := range word {
		freq[ch-'a']++
	}

	ans := len(word)
	for _, minFreq := range freq {
		if minFreq == 0 {
			continue
		}
		ops := 0
		for _, f := range freq {
			if f < minFreq {
				ops += f
			} else if f > minFreq+k {
				ops += f - (minFreq + k)
			}
		}
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumDeletions("aabcaba", 2)) // Expected: 2
	fmt.Println(minimumDeletions("dabdcbdcdcd", 2)) // Expected: 2
	fmt.Println(minimumDeletions("aaabaaa", 2)) // Expected: 0
}
```
