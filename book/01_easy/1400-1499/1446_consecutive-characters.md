# 1446 — Consecutive Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPower(s string) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1446: Consecutive Characters
// https://leetcode.com/problems/consecutive-characters/
// Difficulty: Easy
//
// LeetCode submission: func maxPower(s string) int

import "fmt"

func main() {
	fmt.Println(ConsecutiveCharacters("leetcode")) // 2
	fmt.Println(ConsecutiveCharacters("abbcccddddeeeeedcba")) // 5
	fmt.Println(ConsecutiveCharacters("triplepillooooow")) // 5
}

// Time: O(n), Space: O(1)
func ConsecutiveCharacters(s string) int {
	ans, cur := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
			if cur > ans {
				ans = cur
			}
		} else {
			cur = 1
		}
	}
	return ans
}
```
