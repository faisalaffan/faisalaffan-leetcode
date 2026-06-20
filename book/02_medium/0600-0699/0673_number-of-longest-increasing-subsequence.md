# 0673 — Number Of Longest Increasing Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func findNumberOfLIS(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #673: Number of Longest Increasing Subsequence
// https://leetcode.com/problems/number-of-longest-increasing-subsequence/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

  // Alokasi slice integer
	length := make([]int, n)
  // Alokasi slice integer
	count := make([]int, n)
	maxLen := 0

	for i := 0; i < n; i++ {
		length[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j]+1 > length[i] {
					length[i] = length[j] + 1
					count[i] = count[j]
				} else if length[j]+1 == length[i] {
					count[i] += count[j]
				}
			}
		}
		if length[i] > maxLen {
			maxLen = length[i]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		if length[i] == maxLen {
			result += count[i]
		}
	}
	return result
}
```
