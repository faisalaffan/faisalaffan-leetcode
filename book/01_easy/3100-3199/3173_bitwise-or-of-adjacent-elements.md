# 3173 — Bitwise Or Of Adjacent Elements

## Deskripsi

**Soal:** [3173. Bitwise Or Of Adjacent Elements](https://leetcode.com/problems/bitwise-or-of-adjacent-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3173: Bitwise OR of Adjacent Elements
// https://leetcode.com/problems/bitwise-or-of-adjacent-elements/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// LeetCode name: orArray
	fmt.Println(BitwiseOrOfAdjacentElements([]int{1, 2, 3, 4})) // [3, 3, 7]
	fmt.Println(BitwiseOrOfAdjacentElements([]int{5, 1, 6}))     // [5, 7]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: orArray
func BitwiseOrOfAdjacentElements(nums []int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(nums)-1)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-1; i++ {
		result[i] = nums[i] | nums[i+1]
	}
	return result
}
```
