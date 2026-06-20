# 0991 — Broken Calculator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func brokenCalc(startValue int, target int) int
```

> **💡 Hint:** Work backwards from target to startValue

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log target)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #991: Broken Calculator
// https://leetcode.com/problems/broken-calculator/
// Difficulty: Medium
//
// Approach: Work backwards from target to startValue
//   - If target is even, divide by 2 (reverse of multiply by 2)
//   - If target is odd, add 1 (reverse of subtract 1)
// Time: O(log target)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))  // 2
	fmt.Println(brokenCalc(5, 8))  // 2
	fmt.Println(brokenCalc(3, 10)) // 3
}

func brokenCalc(startValue int, target int) int {
	ops := 0
	for target > startValue {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		ops++
	}
	return ops + (startValue - target)
}
```
