# 1915 — Number Of Wonderful Substrings

## Deskripsi

**Soal:** [1915. Number Of Wonderful Substrings](https://leetcode.com/problems/number-of-wonderful-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 10) = O(n), Space: O(2^10) = O(1)  
**Kompleksitas Ruang:** O(2^10) = O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1915: Number of Wonderful Substrings
// https://leetcode.com/problems/number-of-wonderful-substrings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WonderfulSubstrings("aba"))
	fmt.Println(WonderfulSubstrings("aabb"))
	fmt.Println(WonderfulSubstrings("he"))
}

// Time: O(n * 10) = O(n), Space: O(2^10) = O(1)
func WonderfulSubstrings(word string) int64 {
	// mask represents parity of each of 10 letters
  // Membuat slice untuk menyimpan hasil
	count := make([]int64, 1024) // 2^10 possible masks
	count[0] = 1
	mask := 0
	var result int64 = 0

	for _, c := range word {
		mask ^= 1 << (c - 'a')

		// Count substrings where all letters have even count
		result += count[mask]

		// Count substrings where exactly one letter has odd count
		for i := 0; i < 10; i++ {
			result += count[mask^(1<<i)]
		}

		count[mask]++
	}
	return result
}
```
