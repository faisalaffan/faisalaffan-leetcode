# 2457 — Minimum Addition To Make Integer Beautiful

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func makeIntegerBeautiful(n int64, target int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n * log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2457: Minimum Addition to Make Integer Beautiful
// https://leetcode.com/problems/minimum-addition-to-make-integer-beautiful/
// Difficulty: Medium
// Time: O(log n * log n) | Space: O(log n)
// Try rounding up to higher digit positions until digit sum <= target.

import "fmt"

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))   // 4 (16+4=20, digit sum 2 <= 6)
	fmt.Println(makeIntegerBeautiful(467, 6))  // 33 (467+33=500, digit sum 5 <= 6)
	fmt.Println(makeIntegerBeautiful(1, 1))    // 0
}

func makeIntegerBeautiful(n int64, target int) int64 {
	if digitSum(n) <= target {
		return 0
	}

	pow10 := int64(10)
	for {
		next := ((n / pow10) + 1) * pow10
		if digitSum(next) <= target {
			return next - n
		}
		pow10 *= 10
	}
}

func digitSum(n int64) int {
	sum := 0
	for n > 0 {
		sum += int(n % 10)
		n /= 10
	}
	return sum
}
```
