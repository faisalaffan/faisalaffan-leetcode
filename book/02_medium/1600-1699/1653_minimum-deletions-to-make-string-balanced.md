# 1653 — Minimum Deletions To Make String Balanced

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumDeletions(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1653: Minimum Deletions to Make String Balanced
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumDeletions("aababbab"))
	fmt.Println(MinimumDeletions("bbaaaaabb"))
	fmt.Println(MinimumDeletions("a"))
}

func MinimumDeletions(s string) int {
	// Time: O(N), Space: O(1)
	// Count 'a's on the right
	aCount := 0
	for _, ch := range s {
		if ch == 'a' {
			aCount++
		}
	}

	bCount := 0
	minDeletions := len(s)

	for _, ch := range s {
		if ch == 'a' {
			aCount--
		}

		// Deletions needed: remove all 'b's before this point + remove all 'a's after
		deletions := bCount + aCount
		if deletions < minDeletions {
			minDeletions = deletions
		}

		if ch == 'b' {
			bCount++
		}
	}

	return minDeletions
}
```
