# 0898 — Bitwise Ors Of Subarrays

## Deskripsi

**Soal:** [0898. Bitwise Ors Of Subarrays](https://leetcode.com/problems/bitwise-ors-of-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * log(max))  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #898: Bitwise ORs of Subarrays
// https://leetcode.com/problems/bitwise-ors-of-subarrays/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BitwiseOrsOfSubarrays([]int{0}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 1, 2}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 2, 4}))
}

// Time: O(n * log(max)) | Space: O(n)
func BitwiseOrsOfSubarrays(arr []int) int {
  // Membuat map untuk pencarian O(1): key → value
	set := make(map[int]bool)

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
		for j := i - 1; j >= 0; j-- {
			if arr[i]|arr[j] == arr[j] {
				break
			}
			arr[j] |= arr[i]
			set[arr[j]] = true
		}
	}

	return len(set)
}
```
