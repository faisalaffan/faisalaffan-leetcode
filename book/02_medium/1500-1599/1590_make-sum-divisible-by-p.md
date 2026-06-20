# 1590 — Make Sum Divisible By P

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinSubarray(nums []int, p int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1590: Make Sum Divisible by P
// https://leetcode.com/problems/make-sum-divisible-by-p/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(MinSubarray([]int{6, 3, 5, 2}, 9))
	fmt.Println(MinSubarray([]int{1, 2, 3}, 7))
}

func MinSubarray(nums []int, p int) int {
	// Time: O(N), Space: O(N)
	n := len(nums)

	// Total sum modulo p
	totalSum := 0
	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum // the remainder we need to remove
	if target == 0 {
		return 0
	}

	// Map from prefix sum modulo p to index
  // Membuat map (HashMap) — pencarian O(1)
	prefixMap := make(map[int]int)
	prefixMap[0] = -1
	prefixSum := 0
	minLen := n

	for i, num := range nums {
		prefixSum = (prefixSum + num) % p
		// We need prefixSum - prefixSum[j] ≡ target (mod p)
		// => prefixSum[j] ≡ prefixSum - target (mod p)
		needed := (prefixSum - target + p) % p
		if j, exists := prefixMap[needed]; exists {
			if i-j < minLen {
				minLen = i - j
			}
		}
		prefixMap[prefixSum] = i
	}

	if minLen == n {
		return -1
	}
	return minLen
}
```
