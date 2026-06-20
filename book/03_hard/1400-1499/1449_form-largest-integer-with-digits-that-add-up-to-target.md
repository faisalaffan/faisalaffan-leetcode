# 1449 — Form Largest Integer With Digits That Add Up To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestNumber(cost []int, target int) string
```

> **💡 Hint:** DP to find maximum length for each cost, then reconstruct

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1449: Form Largest Integer With Digits That Add up to Target
// https://leetcode.com/problems/form-largest-integer-with-digits-that-add-up-to-target/
// Difficulty: Hard
//
// Given an array cost where cost[i] is the cost of digit (i+1), and a target,
// find the largest integer that can be formed with total cost equal to target.
// Digits can be used multiple times.
//
// Approach: DP to find maximum length for each cost, then reconstruct
// the largest number by trying digits from 9 down to 1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
	// Example 2
	fmt.Println(largestNumber([]int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12))
	// Edge: impossible
	fmt.Println(largestNumber([]int{2, 4, 6, 2, 4, 6, 4, 4, 4}, 1))
	// Edge: single digit
	fmt.Println(largestNumber([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3))
}

func largestNumber(cost []int, target int) string {
	const inf = 1 << 30
  // Membuat matriks/slice 2D untuk DP
	f := make([][]int, 10)
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, 10)
  // Range loop: iterasi dengan indeks + nilai
	for i := range f {
		f[i] = make([]int, target+1)
		g[i] = make([]int, target+1)
		for j := range f[i] {
			f[i][j] = -inf
		}
	}
	f[0][0] = 0
	for i := 1; i <= 9; i++ {
		c := cost[i-1]
		for j := 0; j <= target; j++ {
			if j < c || f[i][j-c]+1 < f[i-1][j] {
				f[i][j] = f[i-1][j]
				g[i][j] = j
			} else {
				f[i][j] = f[i][j-c] + 1
				g[i][j] = j - c
			}
		}
	}
	if f[9][target] < 0 {
		return "0"
	}
	ans := []byte{}
	for i, j := 9, target; i > 0; {
		if g[i][j] == j {
			i--
		} else {
			ans = append(ans, '0'+byte(i))
			j = g[i][j]
		}
	}
	return string(ans)
}
```
