# 2195 — Append K Integers With Minimal Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimalKSum(nums []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2195: Append K Integers With Minimal Sum
// https://leetcode.com/problems/append-k-integers-with-minimal-sum/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	var sum int64 = 0
	prev := 0

	for _, num := range nums {
		if num == prev {
			continue
		}
		if num > prev+1 {
			gap := num - prev - 1
			if gap >= k {
				// arithmetic series: (prev+1) + (prev+2) + ... + (prev+k)
				first := int64(prev) + 1
				last := int64(prev) + int64(k)
				sum += (first + last) * int64(k) / 2
				k = 0
				break
			}
			first := int64(prev) + 1
			last := int64(num) - 1
			sum += (first + last) * int64(gap) / 2
			k -= gap
		}
		prev = num
		if k == 0 {
			break
		}
	}

	if k > 0 {
		first := int64(prev) + 1
		last := int64(prev) + int64(k)
		sum += (first + last) * int64(k) / 2
	}

	return sum
}

func main() {
	// Test case 1
	fmt.Println(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	// Expected: 5

	// Test case 2
	fmt.Println(minimalKSum([]int{5, 6}, 6))
	// Expected: 25
}
```
