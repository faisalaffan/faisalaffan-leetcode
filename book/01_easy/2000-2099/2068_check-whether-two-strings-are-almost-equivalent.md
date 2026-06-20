# 2068 — Check Whether Two Strings Are Almost Equivalent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckWhetherTwoStringsAreAlmostEquivalent(word1 string, word2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2068: Check Whether Two Strings are Almost Equivalent
// https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("aaaa", "bccb"))   // false
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("abcdeef", "abaaacc")) // true
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("cccddabba", "babababab")) // true
}

// Time: O(n), Space: O(1)
func CheckWhetherTwoStringsAreAlmostEquivalent(word1 string, word2 string) bool {
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word1); i++ {
		freq[word1[i]-'a']++
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word2); i++ {
		freq[word2[i]-'a']--
	}
	for _, v := range freq {
		if v < 0 {
			v = -v
		}
		if v > 3 {
			return false
		}
	}
	return true
}
```
