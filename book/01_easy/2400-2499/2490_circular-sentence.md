# 2490 — Circular Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CircularSentence(sentence string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2490: Circular Sentence
// https://leetcode.com/problems/circular-sentence/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CircularSentence("leetcode exercises sound delightful")) // true
	fmt.Println(CircularSentence("eetcode"))                            // true
	fmt.Println(CircularSentence("Leetcode is cool"))                   // false
}

func CircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
  // Linear scan O(n)
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
```
