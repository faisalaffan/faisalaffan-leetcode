# 3789 — Minimum Cost To Acquire Required Items

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCostToAcquireRequiredItems(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3789: Minimum Cost to Acquire Required Items
// https://leetcode.com/problems/minimum-cost-to-acquire-required-items/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumCostToAcquireRequiredItems(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	a := int64(need1)*int64(cost1) + int64(need2)*int64(cost2)
	b := int64(costBoth) * int64(max(need1, need2))
	mn := min(need1, need2)
	c := int64(costBoth)*int64(mn) + int64(need1-mn)*int64(cost1) + int64(need2-mn)*int64(cost2)

	ans := a
	if b < ans {
		ans = b
	}
	if c < ans {
		ans = c
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumCostToAcquireRequiredItems(3, 2, 1, 3, 2))
	fmt.Println(minimumCostToAcquireRequiredItems(5, 4, 15, 2, 3))
	fmt.Println(minimumCostToAcquireRequiredItems(10, 10, 5, 5, 5))
}
```
