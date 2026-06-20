# 3884 — First Matching Character From Both Ends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstMatchingCharacterFromBothEnds(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3884: First Matching Character From Both Ends
// https://leetcode.com/problems/first-matching-character-from-both-ends/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FirstMatchingCharacterFromBothEnds("abcacbd"))
	fmt.Println(FirstMatchingCharacterFromBothEnds("abc"))
	fmt.Println(FirstMatchingCharacterFromBothEnds("abcdab"))
}

// Time: O(n)
// Space: O(1)
func FirstMatchingCharacterFromBothEnds(s string) int {
	n := len(s)
	for i := 0; i <= n/2; i++ {
		if s[i] == s[n-1-i] {
			return i
		}
	}
	return -1
}
```
