# 1554 — Strings Differ By One Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func DifferByOne(dict []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N*M^2) where N = len(dict), M = string length  
**Kompleksitas Ruang:** O(N*M)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1554: Strings Differ by One Character
// https://leetcode.com/problems/strings-differ-by-one-character/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(DifferByOne([]string{"abcd", "acbd", "aacd"}))
	fmt.Println(DifferByOne([]string{"ab", "cd", "yz"}))
	fmt.Println(DifferByOne([]string{"abcd", "cccc", "abxd", "abzd"}))
}

func DifferByOne(dict []string) bool {
	// Time: O(N*M^2) where N = len(dict), M = string length
	// Space: O(N*M)
	// Use rolling hash: for each position, check if any two strings
	// become identical when that position is skipped.

	n := len(dict)
	if n < 2 {
		return false
	}
	m := len(dict[0])

	for skipIdx := 0; skipIdx < m; skipIdx++ {
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[string]bool)
		for i := 0; i < n; i++ {
			// Create string without char at skipIdx
			key := dict[i][:skipIdx] + dict[i][skipIdx+1:]
			if seen[key] {
				return true
			}
			seen[key] = true
		}
	}

	return false
}
```
