# 0283 — Move Zeroes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MoveZeroes(nums []int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #283: Move Zeroes
// https://leetcode.com/problems/move-zeroes/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MoveZeroes(nums []int) {
	lastNonZero := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

func main() {
	n1 := []int{0, 1, 0, 3, 12}
	MoveZeroes(n1)
	fmt.Println(n1)
	n2 := []int{0}
	MoveZeroes(n2)
	fmt.Println(n2)
}
```
