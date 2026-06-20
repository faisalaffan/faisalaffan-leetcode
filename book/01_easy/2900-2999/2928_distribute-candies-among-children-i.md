# 2928 — Distribute Candies Among Children I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DistributeCandiesAmongChildrenI(n int, limit int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(limit^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2928: Distribute Candies Among Children I
// https://leetcode.com/problems/distribute-candies-among-children-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: distributeCandies
	fmt.Println(DistributeCandiesAmongChildrenI(5, 2)) // 3
	fmt.Println(DistributeCandiesAmongChildrenI(3, 3)) // 10
}

// Time: O(limit^2) | Space: O(1)
// LeetCode submission name: distributeCandies
func DistributeCandiesAmongChildrenI(n int, limit int) int {
	ways := 0
	for a := 0; a <= limit && a <= n; a++ {
		for b := 0; b <= limit && a+b <= n; b++ {
			c := n - a - b
			if c <= limit {
				ways++
			}
		}
	}
	return ways
}
```
