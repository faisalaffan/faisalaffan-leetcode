# 1347 — Minimum Number Of Steps To Make Two Strings Anagram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minSteps(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = length of strings  
**Kompleksitas Ruang:** O(1) - fixed size array of 26

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1347: Minimum Number of Steps to Make Two Strings Anagram
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minSteps("bab", "aba")) // 1

	// Test case 2
	fmt.Println(minSteps("leetcode", "practice")) // 5

	// Test case 3
	fmt.Println(minSteps("anagram", "mangaar")) // 0
}

// Time: O(n) where n = length of strings
// Space: O(1) - fixed size array of 26
func minSteps(s string, t string) int {
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	steps := 0
	for _, f := range freq {
		if f > 0 {
			steps += f
		}
	}
	return steps
}
```
