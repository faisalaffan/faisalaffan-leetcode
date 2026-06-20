# 2207 — Maximize Number Of Subsequences In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSubsequenceCount(text string, pattern string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2207: Maximize Number of Subsequences in a String
// https://leetcode.com/problems/maximize-number-of-subsequences-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSubsequenceCount(text string, pattern string) int64 {
	var count int64 = 0
	var first, second int64 = 0, 0

	for _, ch := range text {
		if byte(ch) == pattern[1] {
			count += first
			second++
		}
		if byte(ch) == pattern[0] {
			first++
		}
	}

	if first > second {
		count += first
	} else {
		count += second
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(maximumSubsequenceCount("abdcdbc", "ac"))
	// Expected: 4

	// Test case 2
	fmt.Println(maximumSubsequenceCount("aabb", "ab"))
	// Expected: 6

	// Test case 3
	fmt.Println(maximumSubsequenceCount("fwymvreuftzgrcrxczjacqovduqaiig", "yy"))
	// Expected: 2
}
```
