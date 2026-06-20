# 2114 — Maximum Number Of Words Found In Sentences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumNumberOfWordsFoundInSentences(sentences []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2114: Maximum Number of Words Found in Sentences
// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})) // 6
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"please wait", "continue to fight", "continue to win"}))                              // 3
}

// Time: O(n * m), Space: O(1)
func MaximumNumberOfWordsFoundInSentences(sentences []string) int {
	maxWords := 0
	for _, s := range sentences {
		count := strings.Count(s, " ") + 1
		if count > maxWords {
			maxWords = count
		}
	}
	return maxWords
}
```
