# 2486 — Append Characters To String To Make Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func appendCharacters(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2486: Append Characters to String to Make Subsequence
// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/
// Difficulty: Medium
// Time: O(n + m) | Space: O(1)
// Two pointers: match as much of t in s as possible.

import "fmt"

func main() {
	fmt.Println(appendCharacters("coaching", "coding")) // 4
	fmt.Println(appendCharacters("abcde", "a"))         // 0
	fmt.Println(appendCharacters("z", "abcde"))         // 5
}

func appendCharacters(s string, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return len(t) - j
}
```
