# 0953 — Verifying An Alien Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isAlienSorted(words []string, order string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #953: Verifying an Alien Dictionary
// https://leetcode.com/problems/verifying-an-alien-dictionary/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")) // false
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz")) // false
}

// isAlienSorted checks if words are sorted in the alien language order.
// Time: O(n * m). Space: O(1).
func isAlienSorted(words []string, order string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	orderMap := make(map[byte]int)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		if !isLess(words[i-1], words[i], orderMap) {
			return false
		}
	}
	return true
}

func isLess(a, b string, order map[byte]int) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if order[a[i]] < order[b[i]] {
			return true
		}
		if order[a[i]] > order[b[i]] {
			return false
		}
	}
	return len(a) <= len(b)
}
```
