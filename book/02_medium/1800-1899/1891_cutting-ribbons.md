# 1891 — Cutting Ribbons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxLength(ribbons []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log maxLen), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1891: Cutting Ribbons
// https://leetcode.com/problems/cutting-ribbons/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxLength([]int{9, 7, 5}, 3))
	fmt.Println(MaxLength([]int{7, 5, 9}, 4))
	fmt.Println(MaxLength([]int{5, 7, 9}, 22))
}

// Time: O(n log maxLen), Space: O(1)
func MaxLength(ribbons []int, k int) int {
	left, right := 1, 0
	for _, r := range ribbons {
		if r > right {
			right = r
		}
	}

	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		count := 0
		for _, r := range ribbons {
			count += r / mid
		}
		if count >= k {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
```
