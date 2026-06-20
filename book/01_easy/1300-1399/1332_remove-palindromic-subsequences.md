# 1332 — Remove Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removePalindromeSub(s string) int

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

// LeetCode #1332: Remove Palindromic Subsequences
// https://leetcode.com/problems/remove-palindromic-subsequences/
// Difficulty: Easy
//
// LeetCode submission: func removePalindromeSub(s string) int

import "fmt"

func main() {
	fmt.Println(RemovePalindromicSubsequences("ababa"))  // 1 (already palindrome)
	fmt.Println(RemovePalindromicSubsequences("abb"))    // 2
	fmt.Println(RemovePalindromicSubsequences("baabb"))  // 2
}

// Time: O(n), Space: O(1)
func RemovePalindromicSubsequences(s string) int {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}
```
