# 1662 — Check If Two String Arrays Are Equivalent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrayStringsAreEqual(word1 []string, word2 []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n) where n is total characters  
**Kompleksitas Ruang:** O(n) where n is total characters

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1662: Check If Two String Arrays Are Equivalent
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
// Difficulty: Easy

import "fmt"
import "strings"

// Time: O(n), Space: O(n) where n is total characters
func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func main() {
	fmt.Println(ArrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(ArrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(ArrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
```
