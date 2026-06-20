# 3118 — Friday Purchase Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func fridayPurchase(purchases [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3118: Friday Purchase III
// https://leetcode.com/problems/friday-purchase-iii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func fridayPurchase(purchases [][]int) int64 {
	total := int64(0)
	for _, p := range purchases {
		// p[0] = day of week (5 = Friday), p[1] = amount
		if p[0] == 5 {
			total += int64(p[1])
		}
	}
	return total
}

func main() {
	fmt.Println(fridayPurchase([][]int{{5, 100}, {1, 50}, {5, 200}, {3, 75}})) // Expected: 300
	fmt.Println(fridayPurchase([][]int{{2, 50}, {3, 100}}))                     // Expected: 0
}
```
