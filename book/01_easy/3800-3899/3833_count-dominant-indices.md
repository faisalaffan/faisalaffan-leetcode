# 3833 — Count Dominant Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountDominantIndices(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3833: Count Dominant Indices
// https://leetcode.com/problems/count-dominant-indices/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountDominantIndices([]int{5, 4, 3}))
	fmt.Println(CountDominantIndices([]int{4, 1, 2}))
	fmt.Println(CountDominantIndices([]int{1}))
}

// Time: O(n)
// Space: O(n)
func CountDominantIndices(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
  // Alokasi slice integer
	suffixSum := make([]int, n)
	suffixSum[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + nums[i]
	}
	count := 0
	for i := 0; i < n-1; i++ {
		rightSum := suffixSum[i+1]
		rightCount := n - 1 - i
		if nums[i]*rightCount > rightSum {
			count++
		}
	}
	return count
}
```
