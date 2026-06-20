# 2185 — Counting Words With A Given Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountingWordsWithAGivenPrefix(words []string, pref string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2185: Counting Words With a Given Prefix
// https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"pay", "attention", "practice", "attend"}, "at")) // 2
	fmt.Println(CountingWordsWithAGivenPrefix([]string{"leetcode", "win", "loops", "success"}, "code"))   // 0
}

// Time: O(n * m), Space: O(1)
func CountingWordsWithAGivenPrefix(words []string, pref string) int {
	count := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			count++
		}
	}
	return count
}
```
