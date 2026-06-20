# 1859 — Sorting The Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortSentence(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1859: Sorting the Sentence
// https://leetcode.com/problems/sorting-the-sentence/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(n)
func SortSentence(s string) string {
	words := strings.Split(s, " ")
	result := make([]string, len(words))
	for _, w := range words {
		pos := int(w[len(w)-1] - '0') - 1
		result[pos] = w[:len(w)-1]
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(SortSentence("is2 sentence4 This1 a3"))
	fmt.Println(SortSentence("Myself2 Me1 I4 and3"))
}
```
