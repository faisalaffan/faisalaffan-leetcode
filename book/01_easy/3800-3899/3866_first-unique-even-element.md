# 3866 — First Unique Even Element

## Deskripsi

**Soal:** [3866. First Unique Even Element](https://leetcode.com/problems/first-unique-even-element/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3866: First Unique Even Element
// https://leetcode.com/problems/first-unique-even-element/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FirstUniqueEvenElement([]int{3, 4, 2, 5, 4, 6}))
	fmt.Println(FirstUniqueEvenElement([]int{4, 4}))
	fmt.Println(FirstUniqueEvenElement([]int{1, 3, 5}))
}

// Time: O(n)
// Space: O(n)
func FirstUniqueEvenElement(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, v := range nums {
		if v%2 == 0 && freq[v] == 1 {
			return v
		}
	}
	return -1
}
```
