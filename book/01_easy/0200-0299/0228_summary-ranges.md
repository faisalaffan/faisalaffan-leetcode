# 0228 — Summary Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SummaryRanges(nums []int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #228: Summary Ranges
// https://leetcode.com/problems/summary-ranges/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

// Time: O(n) | Space: O(1) excluding output
func SummaryRanges(nums []int) []string {
	var res []string
	i := 0
	for i < len(nums) {
		start := nums[i]
		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}
		if start == nums[i] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		}
		i++
	}
	return res
}

func main() {
	fmt.Println(SummaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(SummaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(SummaryRanges([]int{}))
}
```
