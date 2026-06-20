# 0660 — Remove 9

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func newInteger(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #660: Remove 9
// https://leetcode.com/problems/remove-9/
// Difficulty: Hard [Paid]

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 10},
		{10, 11},
		{11, 12},
		{12, 13},
		{13, 14},
		{14, 15},
		{15, 16},
		{16, 17},
		{17, 18},
		{18, 20},
		{80, 88},
		{81, 100},
		{100, 121},
		{500, 615},
		{1000, 1331},
	}

	for _, tc := range testCases {
		got := newInteger(tc.n)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: newInteger(%d) = %d (want %d)\n", status, tc.n, got, tc.want)
	}
}

func newInteger(n int) int {
	// Convert n to base-9, interpret result as base-10 number
	result := 0
	multiplier := 1

	for n > 0 {
		result += (n % 9) * multiplier
		n /= 9
		multiplier *= 10
	}

	return result
}
```
