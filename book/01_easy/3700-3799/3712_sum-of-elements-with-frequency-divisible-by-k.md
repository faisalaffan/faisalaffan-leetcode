# 3712 — Sum Of Elements With Frequency Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SumOfElementsWithFrequencyDivisibleByK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3712: Sum of Elements With Frequency Divisible by K
// https://leetcode.com/problems/sum-of-elements-with-frequency-divisible-by-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{1, 2, 2, 3, 3, 3, 3, 4}, 2))
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{4, 4, 4, 1, 2, 3}, 3))
}

// Time: O(n)
// Space: O(1)
func SumOfElementsWithFrequencyDivisibleByK(nums []int, k int) int {
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}

	sum := 0
	for v := 1; v <= 100; v++ {
		if cnt[v]%k == 0 {
			sum += v * cnt[v]
		}
	}
	return sum
}
```
