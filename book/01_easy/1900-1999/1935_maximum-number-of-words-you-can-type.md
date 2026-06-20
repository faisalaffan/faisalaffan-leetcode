# 1935 — Maximum Number Of Words You Can Type

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumNumberOfWordsYouCanType(text string, brokenLetters string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m), Space: O(k) where k = len(brokenLetters)  
**Kompleksitas Ruang:** O(k) where k = len(brokenLetters)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1935: Maximum Number of Words You Can Type
// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsYouCanType("hello world", "ad"))                   // 1
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "e"))                      // 0
	fmt.Println(MaximumNumberOfWordsYouCanType("leet code", "lt"))                     // 1
}

// Time: O(n + m), Space: O(k) where k = len(brokenLetters)
func MaximumNumberOfWordsYouCanType(text string, brokenLetters string) int {
  // Membuat map (HashMap) — pencarian O(1)
	broken := make(map[byte]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(brokenLetters); i++ {
		broken[brokenLetters[i]] = true
	}

	words := strings.Fields(text)
	count := 0
	for _, word := range words {
		canType := true
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(word); i++ {
			if broken[word[i]] {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}
	return count
}
```
