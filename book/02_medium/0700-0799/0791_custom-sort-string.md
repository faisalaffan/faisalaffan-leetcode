# 0791 — Custom Sort String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func customSortString(order string, s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #791: Custom Sort String
// https://leetcode.com/problems/custom-sort-string/
// Difficulty: Medium
// Time: O(n + m)
// Space: O(1)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("bcafg", "abcd"))
}

func customSortString(order string, s string) string {
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	var result strings.Builder
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(order); i++ {
		c := order[i]
		for freq[c-'a'] > 0 {
			result.WriteByte(c)
			freq[c-'a']--
		}
	}

	for i := 0; i < 26; i++ {
		for freq[i] > 0 {
			result.WriteByte(byte(i + 'a'))
			freq[i]--
		}
	}

	return result.String()
}
```
