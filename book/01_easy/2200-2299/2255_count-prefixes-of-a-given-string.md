# 2255 — Count Prefixes Of A Given String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPrefixesOfAGivenString(words []string, s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2255: Count Prefixes of a Given String
// https://leetcode.com/problems/count-prefixes-of-a-given-string/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))           // 3
	fmt.Println(CountPrefixesOfAGivenString([]string{"a", "a"}, "aa"))                                      // 2
	fmt.Println(CountPrefixesOfAGivenString([]string{"feh", "w", "w", "l", "w", "o", "w", "o", "w"}, "w")) // 0
}

func CountPrefixesOfAGivenString(words []string, s string) int {
	count := 0
	for _, w := range words {
		if len(w) <= len(s) && s[:len(w)] == w {
			count++
		}
	}
	return count
}
```
