# 3479 — Fruits Into Baskets Iii

## Deskripsi

**Soal:** [3479. Fruits Into Baskets Iii](https://leetcode.com/problems/fruits-into-baskets-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3479: Fruits Into Baskets III
// https://leetcode.com/problems/fruits-into-baskets-iii/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FruitsIntoBasketsIii([]int{1, 2, 1}))
	// Test case 2
	fmt.Println("Test 2:", FruitsIntoBasketsIii([]int{0, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", FruitsIntoBasketsIii([]int{1, 2, 3, 2, 2}))
}

func FruitsIntoBasketsIii(fruits []int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	left := 0
	maxLen := 0
	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++
		for len(freq) > 2 {
			freq[fruits[left]]--
			if freq[fruits[left]] == 0 {
				delete(freq, fruits[left])
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
