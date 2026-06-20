# 2150 — Find All Lonely Numbers In The Array

## Deskripsi

**Soal:** [2150. Find All Lonely Numbers In The Array](https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findLonely(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #2150: Find All Lonely Numbers in the Array
// https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findLonely(nums []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	result := []int{}
	for _, v := range nums {
		if freq[v] == 1 && freq[v-1] == 0 && freq[v+1] == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findLonely([]int{10, 6, 5, 8}))
	// Expected: [10, 8]

	// Test case 2
	fmt.Println("Test 2:", findLonely([]int{1, 3, 5, 3}))
	// Expected: [1, 5]
}
```
