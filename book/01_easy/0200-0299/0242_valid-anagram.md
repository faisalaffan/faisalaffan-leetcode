# 0242 — Valid Anagram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func IsAnagram(s string, t string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (fixed 26 chars)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed 26 chars)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	return count == [26]int{}
}

func main() {
	fmt.Println(IsAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram("rat", "car"))
}
```
