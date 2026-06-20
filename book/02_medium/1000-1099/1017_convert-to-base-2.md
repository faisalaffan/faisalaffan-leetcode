# 1017 — Convert To Base 2

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func baseNeg2(n int) string
```

> **💡 Hint:** Repeated division by -2. Handle negative remainder.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1017: Convert to Base -2
// https://leetcode.com/problems/convert-to-base-2/
// Difficulty: Medium
//
// Approach: Repeated division by -2. Handle negative remainder.
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(baseNeg2(2))  // "110"
	fmt.Println(baseNeg2(3))  // "111"
	fmt.Println(baseNeg2(4))  // "100"
}

func baseNeg2(n int) string {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return "0"
	}

	result := ""
	for n != 0 {
		remainder := n % -2
		n /= -2
		if remainder < 0 {
			remainder += 2
			n++
		}
		result = string(rune('0'+remainder)) + result
	}

	return result
}
```
