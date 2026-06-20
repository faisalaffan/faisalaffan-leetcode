# 1768 — Merge Strings Alternately

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeAlternately(word1 string, word2 string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m), Space: O(n+m)  
**Kompleksitas Ruang:** O(n+m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1768: Merge Strings Alternately
// https://leetcode.com/problems/merge-strings-alternately/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n+m)
func MergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		result = append(result, word1[i], word2[j])
		i++
		j++
	}
	result = append(result, word1[i:]...)
	result = append(result, word2[j:]...)
	return string(result)
}

func main() {
	fmt.Println(MergeAlternately("abc", "pqr"))
	fmt.Println(MergeAlternately("ab", "pqrs"))
	fmt.Println(MergeAlternately("abcd", "pq"))
}
```
