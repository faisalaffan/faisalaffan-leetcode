# 2108 — Find First Palindromic String In The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindFirstPalindromicStringInTheArray(words []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2108: Find First Palindromic String in the Array
// https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"abc", "car", "ada", "racecar", "cool"})) // "ada"
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"notapalindrome", "racecar"}))             // "racecar"
	fmt.Println(FindFirstPalindromicStringInTheArray([]string{"def", "ghi"}))                             // ""
}

// Time: O(n * m), Space: O(1)
func FindFirstPalindromicStringInTheArray(words []string) string {
	for _, w := range words {
		if isPalindrome(w) {
			return w
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```
