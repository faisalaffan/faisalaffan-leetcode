# 3571 — Find The Shortest Superstring Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func overlap(a, b string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n*m). Space: O(n+m).  
**Kompleksitas Ruang:** O(n+m).

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3571: Find the Shortest Superstring II
// https://leetcode.com/problems/find-the-shortest-superstring-ii/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(FindTheShortestSuperstringIi("abc", "bcd"))
	fmt.Println(FindTheShortestSuperstringIi("abc", "xyz"))
}

// overlap returns the length of overlapping suffix of a that matches prefix of b.
func overlap(a, b string) int {
	maxOvlp := 0
	maxLen := len(a)
	if len(b) < maxLen {
		maxLen = len(b)
	}
	for i := 1; i <= maxLen; i++ {
		if a[len(a)-i:] == b[:i] {
			maxOvlp = i
		}
	}
	return maxOvlp
}

// FindTheShortestSuperstringIi returns the shortest string that contains both s1 and s2 as substrings.
// Time: O(n*m). Space: O(n+m).
func FindTheShortestSuperstringIi(s1 string, s2 string) string {
	ovlp12 := overlap(s1, s2)
	ovlp21 := overlap(s2, s1)

	s1s2 := s1 + s2[ovlp12:]
	s2s1 := s2 + s1[ovlp21:]

	if len(s1s2) < len(s2s1) {
		return s1s2
	} else if len(s2s1) < len(s1s2) {
		return s2s1
	}
	// Same length, return lexicographically smaller
	if s1s2 < s2s1 {
		return s1s2
	}
	return s2s1
}
```
