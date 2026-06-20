# 3838 — Weighted Word Mapping

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func WeightedWordMapping(words []string, weights []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N * L)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3838: Weighted Word Mapping
// https://leetcode.com/problems/weighted-word-mapping/
// Difficulty: Easy

import "fmt"

func main() {
	weights1 := []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}
	fmt.Println(WeightedWordMapping([]string{"abcd", "def", "xyz"}, weights1))

  // Alokasi slice integer
	weights2 := make([]int, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range weights2 {
		weights2[i] = 1
	}
	fmt.Println(WeightedWordMapping([]string{"a", "b", "c"}, weights2))
}

// Time: O(N * L)
// Space: O(1)
func WeightedWordMapping(words []string, weights []int) string {
	result := make([]byte, len(words))
	for i, word := range words {
		total := 0
		for _, ch := range word {
			total += weights[ch-'a']
		}
		result[i] = byte('z' - total%26)
	}
	return string(result)
}
```
