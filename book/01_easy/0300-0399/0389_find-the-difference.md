# 0389 — Find The Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheDifference(s, t string) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #389: Find the Difference
// https://leetcode.com/problems/find-the-difference/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindTheDifference(s, t string) byte {
	var diff byte
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		diff ^= s[i]
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(t); i++ {
		diff ^= t[i]
	}
	return diff
}

func main() {
	fmt.Printf("%c\n", FindTheDifference("abcd", "abcde"))
	fmt.Printf("%c\n", FindTheDifference("", "y"))
	fmt.Printf("%c\n", FindTheDifference("a", "aa"))
}
```
