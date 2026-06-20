# 2824 — Count Pairs Whose Sum Is Less Than Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPairsWhoseSumIsLessThanTarget(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2824: Count Pairs Whose Sum is Less than Target
// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}

func CountPairsWhoseSumIsLessThanTarget(nums []int, target int) int {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}
```
