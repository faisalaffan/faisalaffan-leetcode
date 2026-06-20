# 1652 — Defuse The Bomb

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Decrypt(code []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*  
**Kompleksitas Ruang:** O(n) (or O(1) excluding output)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1652: Defuse the Bomb
// https://leetcode.com/problems/defuse-the-bomb/
// Difficulty: Easy

import "fmt"

// Time: O(n*|k|), Space: O(n) (or O(1) excluding output)
func Decrypt(code []int, k int) []int {
	n := len(code)
  // Alokasi slice integer
	result := make([]int, n)
	if k == 0 {
		return result
	}
	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
		} else {
			for j := 1; j <= -k; j++ {
				sum += code[(i-j+n)%n]
			}
		}
		result[i] = sum
	}
	return result
}

func main() {
	fmt.Println(Decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(Decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(Decrypt([]int{2, 4, 9, 3}, -2))
}
```
