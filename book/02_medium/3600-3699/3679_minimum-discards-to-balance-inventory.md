# 3679 — Minimum Discards To Balance Inventory

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDiscardsToBalanceInventory(arrivals []int, w int, m int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(max(arrivals))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3679: Minimum Discards to Balance Inventory
// https://leetcode.com/problems/minimum-discards-to-balance-inventory/
// Difficulty: Medium
// Time: O(n) | Space: O(max(arrivals))

import "fmt"

func minimumDiscardsToBalanceInventory(arrivals []int, w int, m int) int {
	maxVal := 0
	for _, v := range arrivals {
		if v > maxVal {
			maxVal = v
		}
	}

  // Alokasi slice integer
	cnt := make([]int, maxVal+1)
	ans := 0

	for i, x := range arrivals {
		if cnt[x] == m {
			arrivals[i] = 0
			ans++
		} else {
			cnt[x]++
		}

		left := i + 1 - w
		if left >= 0 {
			cnt[arrivals[left]]--
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 2, 3, 3, 3, 4}, 3, 2))
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 1, 1, 2, 2}, 2, 2))
	fmt.Println(minimumDiscardsToBalanceInventory([]int{1, 2, 3}, 3, 1))
}
```
