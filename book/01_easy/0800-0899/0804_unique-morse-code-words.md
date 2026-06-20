# 0804 — Unique Morse Code Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func uniqueMorseRepresentations(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m) where n = len(words), m = avg len. Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #804: Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/
// Difficulty: Easy

import "fmt"

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})) // 2
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))                        // 1
}

// uniqueMorseRepresentations counts unique Morse code transformations of words.
// Time: O(n * m) where n = len(words), m = avg len. Space: O(n).
func uniqueMorseRepresentations(words []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[string]bool)
	for _, word := range words {
		var code string
		for _, c := range word {
			code += morse[c-'a']
		}
		set[code] = true
	}
	return len(set)
}
```
