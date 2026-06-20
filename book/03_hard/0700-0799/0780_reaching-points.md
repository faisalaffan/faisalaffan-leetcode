# 0780 — Reaching Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func reachingPoints(sx, sy, tx, ty int) bool
```

> **💡 Hint:** Reverse modulo

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #780: Reaching Points
// https://leetcode.com/problems/reaching-points/
// Difficulty: Hard
//
// From (x,y) can go to (x+y, y) or (x, x+y). Determine if (tx,ty)
// is reachable from (sx,sy).
//
// Approach: Reverse modulo
// Work backwards: if tx > ty, previous x = tx % ty (with adjustment
// for the starting constraint). Use modulo instead of repeated
// subtraction for efficiency.

import "fmt"

func main() {
	fmt.Println(reachingPoints(1, 1, 3, 5)) // true
	fmt.Println(reachingPoints(1, 1, 2, 2)) // false
	fmt.Println(reachingPoints(1, 1, 1, 1)) // true
	fmt.Println(reachingPoints(3, 3, 12, 9)) // true
	fmt.Println(reachingPoints(1, 2, 3, 5)) // true
}

func reachingPoints(sx, sy, tx, ty int) bool {
	for tx > sx && ty > sy {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}

	// Now one coordinate equals the start value
	if tx == sx {
		// Need (ty - sy) >= 0 and divisible by sx
		return (ty-sy) >= 0 && (ty-sy)%sx == 0
	}
	if ty == sy {
		return (tx-sx) >= 0 && (tx-sx)%sy == 0
	}

	return false
}
```
