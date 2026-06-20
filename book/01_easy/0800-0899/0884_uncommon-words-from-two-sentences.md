# 0884 — Uncommon Words From Two Sentences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func uncommonFromSentences(s1 string, s2 string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m). Space: O(n + m).  
**Kompleksitas Ruang:** O(n + m).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #884: Uncommon Words from Two Sentences
// https://leetcode.com/problems/uncommon-words-from-two-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour")) // [sweet sour]
	fmt.Println(uncommonFromSentences("apple apple", "banana"))                     // [banana]
}

// uncommonFromSentences returns all uncommon words across two sentences.
// Time: O(n + m). Space: O(n + m).
func uncommonFromSentences(s1 string, s2 string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[string]int)
	for _, w := range strings.Fields(s1) {
		count[w]++
	}
	for _, w := range strings.Fields(s2) {
		count[w]++
	}
	result := make([]string, 0)
	for w, c := range count {
		if c == 1 {
			result = append(result, w)
		}
	}
	return result
}
```
