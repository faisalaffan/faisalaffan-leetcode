# 0060 — Permutation Sequence

## Deskripsi

**Soal:** [0060. Permutation Sequence](https://leetcode.com/problems/permutation-sequence/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #60: Permutation Sequence
// https://leetcode.com/problems/permutation-sequence/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("60. Permutation Sequence")
	fmt.Println("n=3, k=3:", getPermutation(3, 3), "(expected 213)")
	fmt.Println("n=4, k=9:", getPermutation(4, 9), "(expected 2314)")
	fmt.Println("n=3, k=1:", getPermutation(3, 1), "(expected 123)")
}

func getPermutation(n int, k int) string {
	fact := 1
  // Membuat slice untuk menyimpan hasil
	nums := make([]byte, 0, n)
	for i := 1; i <= n; i++ {
		fact *= i
		nums = append(nums, byte('0'+i))
	}

	k-- // convert to 0-indexed
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		fact /= (n - i)
		idx := k / fact
		result = append(result, nums[idx])
		nums = append(nums[:idx], nums[idx+1:]...)
		k %= fact
	}

	return string(result)
}
```
