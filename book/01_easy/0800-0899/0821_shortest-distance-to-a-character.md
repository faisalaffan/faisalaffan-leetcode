# 0821 — Shortest Distance To A Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestToChar(s string, c byte) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n). Space: O(1) excluding output.  
**Kompleksitas Ruang:** O(1) excluding output.

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #821: Shortest Distance to a Character
// https://leetcode.com/problems/shortest-distance-to-a-character/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e')) // [3,2,1,0,1,0,0,1,2,2,1,0]
	fmt.Println(shortestToChar("aaab", 'b'))          // [3,2,1,0]
}

// shortestToChar returns the shortest distance from each character to the target character c.
// Time: O(n). Space: O(1) excluding output.
func shortestToChar(s string, c byte) []int {
	n := len(s)
  // Alokasi slice integer
	result := make([]int, n)
	// Initialize with large value
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = n
	}

	// Left to right
	last := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			last = i
		}
		result[i] = min(result[i], i-last)
	}

	// Right to left
	last = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			last = i
		}
		result[i] = min(result[i], last-i)
	}
	return result
}
```
