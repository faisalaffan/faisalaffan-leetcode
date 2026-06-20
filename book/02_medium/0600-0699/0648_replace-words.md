# 0648 — Replace Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func replaceWords(dictionary []string, sentence string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L) where n is number of words in sentence, L is avg word length  
**Kompleksitas Ruang:** O(d) where d is total characters in dictionary

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #648: Replace Words
// https://leetcode.com/problems/replace-words/
// Difficulty: Medium
// Time: O(n * L) where n is number of words in sentence, L is avg word length
// Space: O(d) where d is total characters in dictionary

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}

func replaceWords(dictionary []string, sentence string) string {
  // Membuat map (HashMap) — pencarian O(1)
	rootSet := make(map[string]bool)
	for _, root := range dictionary {
		rootSet[root] = true
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if rootSet[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
```
