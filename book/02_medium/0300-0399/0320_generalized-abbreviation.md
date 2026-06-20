# 0320 — Generalized Abbreviation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func generateAbbreviations(word string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(n * 2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #320: Generalized Abbreviation
// https://leetcode.com/problems/generalized-abbreviation/
// Difficulty: Medium [Paid]
// Time: O(n * 2^n) | Space: O(n * 2^n)

import (
	"fmt"
	"strconv"
)

func generateAbbreviations(word string) []string {
	result := []string{}
	backtrack(word, 0, 0, "", &result)
	return result
}

func backtrack(word string, index int, count int, cur string, result *[]string) {
	if index == len(word) {
		if count > 0 {
			cur += strconv.Itoa(count)
		}
		*result = append(*result, cur)
		return
	}

	// Abbreviate current character (increase count)
	backtrack(word, index+1, count+1, cur, result)

	// Keep current character
	if count > 0 {
		cur += strconv.Itoa(count)
	}
	cur += string(word[index])
	backtrack(word, index+1, 0, cur, result)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", generateAbbreviations("word"))
	// Expected: ["word", "1ord", "w1rd", "2rd", "wo1d", "1o1d", "w2d", "3d", "wor1", "1or1", "w1r1", "2r1", "wo2", "1o2", "w3", "4"]

	// Test case 2: Empty string
	fmt.Println("Test 2:", generateAbbreviations(""))
	// Expected: [""]

	// Test case 3: Single character
	fmt.Println("Test 3:", generateAbbreviations("a"))
	// Expected: ["a", "1"]
}
```
