# 2788 — Split Strings By Separator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func SplitStringsBySeparator(words []string, separator byte) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2788: Split Strings by Separator
// https://leetcode.com/problems/split-strings-by-separator/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SplitStringsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
	fmt.Println(SplitStringsBySeparator([]string{"$easy$", "$problem$"}, '$'))
}

func SplitStringsBySeparator(words []string, separator byte) []string {
	result := []string{}
	for _, w := range words {
		parts := strings.Split(w, string(separator))
		for _, p := range parts {
			if p != "" {
				result = append(result, p)
			}
		}
	}
	return result
}
```
