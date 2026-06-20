# 3133 — Minimum Array End

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minEnd(n int, x int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3133: Minimum Array End
// https://leetcode.com/problems/minimum-array-end/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func minEnd(n int, x int) int64 {
	n64 := int64(n - 1)
	x64 := int64(x)
	ans := int64(0)

	bitPos := 0
	for n64 > 0 || x64 > 0 {
		if x64&1 == 1 {
			ans |= (int64(1) << bitPos)
		} else {
			ans |= ((n64 & 1) << bitPos)
			n64 >>= 1
		}
		x64 >>= 1
		bitPos++
	}
	return ans
}

func main() {
	fmt.Println(minEnd(3, 4))  // Expected: 6
	fmt.Println(minEnd(2, 7))  // Expected: 15
	fmt.Println(minEnd(1, 5))  // Expected: 5
}
```
