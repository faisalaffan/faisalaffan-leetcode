# 1181 — Before And After Puzzle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func beforeAndAfterPuzzles(phrases []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** O(n^2 * L)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1181: Before and After Puzzle
// https://leetcode.com/problems/before-and-after-puzzle/
// Difficulty: Medium [Paid]

// Given list of phrases. Merge phrase i and j if last word of i
// equals first word of j. Result: i + j[firstWordLen:].
// Return sorted unique results.

// Time: O(n^2 * L)
// Space: O(n^2)

func beforeAndAfterPuzzles(phrases []string) []string {
	n := len(phrases)
	firstWords := make([]string, n)
	lastWords := make([]string, n)
  // Membuat matriks/slice 2D untuk DP
	words := make([][]string, n)

	for i, p := range phrases {
		words[i] = strings.Fields(p)
		firstWords[i] = words[i][0]
		lastWords[i] = words[i][len(words[i])-1]
	}

  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)
	result := make([]string, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if lastWords[i] == firstWords[j] {
				merged := phrases[i] + phrases[j][len(firstWords[j]):]
				if !seen[merged] {
					seen[merged] = true
					result = append(result, merged)
				}
			}
		}
	}

	sort.Strings(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [writing code rocks])\n",
		beforeAndAfterPuzzles([]string{"writing code", "code rocks"}))

	fmt.Printf("%v (expected: [a d a b c d])\n",
		beforeAndAfterPuzzles([]string{"a b", "b c", "c d"}))
}
```
