# 2268 — Minimum Number Of Keypresses

## Deskripsi

**Soal:** [2268. Minimum Number Of Keypresses](https://leetcode.com/problems/minimum-number-of-keypresses/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + 26 log 26)  
**Kompleksitas Ruang:** O(26)

**Algoritma:** —

**Fungsi Solusi:** `func minimumKeypresses(s string) int`

## Solusi Go

```go
package main

// LeetCode #2268: Minimum Number of Keypresses
// https://leetcode.com/problems/minimum-number-of-keypresses/
// Difficulty: Medium [Paid]
// Time: O(n + 26 log 26) | Space: O(26)

import (
	"fmt"
	"sort"
)

func minimumKeypresses(s string) int {
  // Membuat slice untuk menyimpan hasil
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	presses := 0
	for i, c := range count {
		if c == 0 {
			break
		}
		presses += c * (i/9 + 1)
	}
	return presses
}

func main() {
	// Test case 1
	fmt.Println(minimumKeypresses("apple"))
	// Expected: 5

	// Test case 2
	fmt.Println(minimumKeypresses("abcdefghijkl"))
	// Expected: 15

	// Test case 3
	fmt.Println(minimumKeypresses("aaaaaaa"))
	// Expected: 7
}
```
