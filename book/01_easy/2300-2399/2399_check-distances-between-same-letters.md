# 2399 — Check Distances Between Same Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckDistancesBetweenSameLetters(s string, distance []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2399: Check Distances Between Same Letters
// https://leetcode.com/problems/check-distances-between-same-letters/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckDistancesBetweenSameLetters("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) // true
	fmt.Println(CheckDistancesBetweenSameLetters("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))    // false
}

func CheckDistancesBetweenSameLetters(s string, distance []int) bool {
	first := [26]int{}
	for i := 0; i < 26; i++ {
		first[i] = -1
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if first[idx] == -1 {
			first[idx] = i
		} else if i-first[idx]-1 != distance[idx] {
			return false
		}
	}
	return true
}
```
