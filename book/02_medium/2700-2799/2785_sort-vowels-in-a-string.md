# 2785 — Sort Vowels In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortVowelsInAString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2785: Sort Vowels in a String
// https://leetcode.com/problems/sort-vowels-in-a-string/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func SortVowelsInAString(s string) string {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	vowels := make([]byte, 0)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

  // Custom sort dengan comparator
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	result := []byte(s)
	vi := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			result[i] = vowels[vi]
			vi++
		}
	}

	return string(result)
}

func main() {
	fmt.Println(SortVowelsInAString("lEetcOde"))
	fmt.Println(SortVowelsInAString("lYmpH"))
}
```
