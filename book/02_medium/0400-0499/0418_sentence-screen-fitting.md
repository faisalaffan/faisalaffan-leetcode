# 0418 — Sentence Screen Fitting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func wordsTyping(sentence []string, rows int, cols int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(rows * avgWordLen)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #418: Sentence Screen Fitting
// https://leetcode.com/problems/sentence-screen-fitting/
// Difficulty: Medium [Paid]
// Time: O(rows * avgWordLen) | Space: O(1)

import "fmt"

func wordsTyping(sentence []string, rows int, cols int) int {
	s := ""
	for _, w := range sentence {
		s += w + " "
	}

	n := len(s)
	start := 0

	for i := 0; i < rows; i++ {
		start += cols

		// If next char is a space, we can fit perfectly
		if s[start%n] == ' ' {
			start++
		} else {
			// Backtrack to nearest space
			for start > 0 && s[(start-1)%n] != ' ' {
				start--
			}
		}
	}
	return start / n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wordsTyping([]string{"hello", "world"}, 2, 8))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", wordsTyping([]string{"a", "bcd", "e"}, 3, 6))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", wordsTyping([]string{"i", "had", "apple", "pie"}, 4, 5))
	// Expected: 1
}
```
