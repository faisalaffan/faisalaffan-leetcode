# 1451 — Rearrange Words In A Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func arrangeWords(text string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) where n = number of words  
**Kompleksitas Ruang:** O(n) for storing words

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1451: Rearrange Words in a Sentence
// https://leetcode.com/problems/rearrange-words-in-a-sentence/
// Difficulty: Medium

import "fmt"
import "sort"
import "strings"
import "unicode"

func main() {
	// Test case 1
	fmt.Println(arrangeWords("Leetcode is cool")) // "Is cool leetcode"

	// Test case 2
	fmt.Println(arrangeWords("Keep calm and code on")) // "On and keep calm code"

	// Test case 3
	fmt.Println(arrangeWords("To be or not to be")) // "To be or to be not"
}

type wordInfo struct {
	word   string
	index  int
	length int
}

// Time: O(n log n) where n = number of words
// Space: O(n) for storing words
func arrangeWords(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	// Lowercase first word
	runes := []rune(words[0])
	runes[0] = unicode.ToLower(runes[0])
	words[0] = string(runes)

	infos := make([]wordInfo, len(words))
	for i, w := range words {
		infos[i] = wordInfo{w, i, len(w)}
	}

  // Custom sort dengan comparator
	sort.Slice(infos, func(i, j int) bool {
		if infos[i].length != infos[j].length {
			return infos[i].length < infos[j].length
		}
		return infos[i].index < infos[j].index
	})

	result := make([]string, len(infos))
	for i, info := range infos {
		result[i] = info.word
	}

	// Uppercase first letter
	firstWord := result[0]
	runes = []rune(firstWord)
	runes[0] = unicode.ToUpper(runes[0])
	result[0] = string(runes)

	return strings.Join(result, " ")
}
```
