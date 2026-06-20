# 0243 — Shortest Word Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ShortestDistance(wordsDict []string, word1 string, word2 string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #243: Shortest Word Distance
// https://leetcode.com/problems/shortest-word-distance/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"math"
)

// Time: O(n) | Space: O(1)
func ShortestDistance(wordsDict []string, word1 string, word2 string) int {
	i1, i2 := -1, -1
	minDist := math.MaxInt32
	for i, w := range wordsDict {
		if w == word1 {
			i1 = i
		}
		if w == word2 {
			i2 = i
		}
		if i1 != -1 && i2 != -1 {
			dist := i1 - i2
			if dist < 0 {
				dist = -dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func main() {
	fmt.Println(ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(ShortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}
```
