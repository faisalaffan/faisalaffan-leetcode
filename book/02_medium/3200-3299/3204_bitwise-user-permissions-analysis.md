# 3204 — Bitwise User Permissions Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func bitwiseUserPermissions(permissions [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3204: Bitwise User Permissions Analysis
// https://leetcode.com/problems/bitwise-user-permissions-analysis/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func bitwiseUserPermissions(permissions [][]int) int {
	if len(permissions) == 0 {
		return 0
	}

	combined := permissions[0][1]
	for i := 1; i < len(permissions); i++ {
		combined |= permissions[i][1]
	}
	return combined
}

func main() {
	fmt.Println(bitwiseUserPermissions([][]int{{1, 1}, {2, 2}, {3, 4}})) // Expected: 7 (1|2|4)
	fmt.Println(bitwiseUserPermissions([][]int{{1, 8}, {2, 3}}))         // Expected: 11 (8|3)
}
```
