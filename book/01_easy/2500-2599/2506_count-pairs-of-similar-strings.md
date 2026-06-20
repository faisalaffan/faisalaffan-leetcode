# 2506 — Count Pairs Of Similar Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func charMask(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2506: Count Pairs Of Similar Strings
// https://leetcode.com/problems/count-pairs-of-similar-strings/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(CountPairsOfSimilarStrings([]string{"aba", "aabb", "abcd", "bac", "aabc"})) // 2
	fmt.Println(CountPairsOfSimilarStrings([]string{"aabb", "ab", "ba"}))                    // 3
}

func charMask(s string) int {
	mask := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

func CountPairsOfSimilarStrings(words []string) int {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(words); i++ {
		maskI := charMask(words[i])
		for j := i + 1; j < len(words); j++ {
			if maskI == charMask(words[j]) {
				count++
			}
		}
	}
	return count
}
```
