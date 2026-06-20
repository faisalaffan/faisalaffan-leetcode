# 0290 — Word Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func WordPattern(pattern string, s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #290: Word Pattern
// https://leetcode.com/problems/word-pattern/
// Difficulty: Easy

import "strings"
import "fmt"

// Time: O(n) | Space: O(n)
func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
  // Membuat map (HashMap) — pencarian O(1)
	p2w := make(map[byte]string)
  // Membuat map (HashMap) — pencarian O(1)
	w2p := make(map[string]byte)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]
		if mappedW, ok := p2w[p]; ok && mappedW != w {
			return false
		}
		if mappedP, ok := w2p[w]; ok && mappedP != p {
			return false
		}
		p2w[p] = w
		w2p[w] = p
	}
	return true
}

func main() {
	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
}
```
