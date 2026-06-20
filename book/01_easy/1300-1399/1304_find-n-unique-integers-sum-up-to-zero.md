# 1304 — Find N Unique Integers Sum Up To Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumZero(n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1304: Find N Unique Integers Sum up to Zero
// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
// Difficulty: Easy
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(sumZero(5)) // [0,1,2,3,-6] or similar
	fmt.Println(sumZero(3)) // [0,1,-1]
	fmt.Println(sumZero(1)) // [0]
}

// LeetCode submission: sumZero
func sumZero(n int) []int {
  // Alokasi slice integer
	ans := make([]int, n)
	if n == 1 {
		return ans // [0]
	}
	half := n / 2
	for i := 0; i < half; i++ {
		ans[i] = i + 1
		ans[i+half] = -(i + 1)
	}
	if n%2 == 1 {
		ans[n-1] = 0
	}
	return ans
}
```
