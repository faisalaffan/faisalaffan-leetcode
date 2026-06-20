# 2739 — Total Distance Traveled

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func TotalDistanceTraveled(mainTank int, additionalTank int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2739: Total Distance Traveled
// https://leetcode.com/problems/total-distance-traveled/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalDistanceTraveled(5, 10))
	fmt.Println(TotalDistanceTraveled(1, 2))
}

func TotalDistanceTraveled(mainTank int, additionalTank int) int {
	total := 0
	for mainTank > 0 {
		if mainTank >= 5 {
			mainTank -= 5
			total += 50
			if additionalTank > 0 {
				additionalTank--
				mainTank++
			}
		} else {
			total += mainTank * 10
			mainTank = 0
		}
	}
	return total
}
```
