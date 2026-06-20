# 3751 — Total Waviness Of Numbers In Range I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func totalWavinessOfNumbersInRangeI(num1 int, num2 int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N * D)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3751: Total Waviness of Numbers in Range I
// https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/
// Difficulty: Medium
// Time: O(N * D) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func totalWavinessOfNumbersInRangeI(num1 int, num2 int) int {
	ans := 0
	for num := num1; num <= num2; num++ {
		s := strconv.Itoa(num)
		for i := 1; i < len(s)-1; i++ {
			if (s[i] > s[i-1] && s[i] > s[i+1]) || (s[i] < s[i-1] && s[i] < s[i+1]) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(totalWavinessOfNumbersInRangeI(1, 100))
	fmt.Println(totalWavinessOfNumbersInRangeI(100, 200))
	fmt.Println(totalWavinessOfNumbersInRangeI(1000, 1050))
}
```
