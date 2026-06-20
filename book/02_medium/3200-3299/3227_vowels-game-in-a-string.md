# 3227 — Vowels Game In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func doesAliceWin(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3227: Vowels Game in a String
// https://leetcode.com/problems/vowels-game-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesAliceWin(s string) bool {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(doesAliceWin("leetcoder")) // Expected: true
	fmt.Println(doesAliceWin("bbcd"))       // Expected: false
}
```
