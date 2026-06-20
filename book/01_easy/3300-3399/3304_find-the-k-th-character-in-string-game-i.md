# 3304 — Find The K Th Character In String Game I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheKThCharacterInStringGameI(k int) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log k). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3304: Find the K-th Character in String Game I
// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(string(FindTheKThCharacterInStringGameI(5)))
	fmt.Println(string(FindTheKThCharacterInStringGameI(10)))
	fmt.Println(string(FindTheKThCharacterInStringGameI(1)))
}

// FindTheKThCharacterInStringGameI returns the k-th character in the string game.
// Start with "a", each step append the incremented version of the current string.
// Time: O(log k). Space: O(1).
func FindTheKThCharacterInStringGameI(k int) byte {
	// The string doubles in length each step. k is 1-indexed.
	// The character at position k depends on how many times we wrap around.
	k-- // 0-indexed
	count := 0
	for k > 0 {
		// Find the largest power of 2 <= k
		msb := 1
		for msb*2 <= k {
			msb *= 2
		}
		k -= msb
		count++
	}
	return byte('a' + byte(count%26))
}
```
