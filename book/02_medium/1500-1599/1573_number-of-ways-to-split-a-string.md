# 1573 — Number Of Ways To Split A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumWays(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1573: Number of Ways to Split a String
// https://leetcode.com/problems/number-of-ways-to-split-a-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumWays("10101"))
	fmt.Println(NumWays("1001"))
	fmt.Println(NumWays("0000"))
}

func NumWays(s string) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	// Count total ones
	totalOnes := 0
	for _, ch := range s {
		if ch == '1' {
			totalOnes++
		}
	}

	if totalOnes == 0 {
		// All zeros: need to choose 2 cut positions out of n-1 gaps
		// C(n-1, 2) = (n-1)*(n-2)/2
		n := len(s)
		return ((n - 1) * (n - 2) / 2) % mod
	}

	if totalOnes%3 != 0 {
		return 0
	}

	onesPerPart := totalOnes / 3
	count := 0
	firstWays := 0
	secondWays := 0

	for _, ch := range s {
		if ch == '1' {
			count++
		}

		if count == onesPerPart {
			firstWays++
		} else if count == 2*onesPerPart {
			secondWays++
		}
	}

	return (firstWays * secondWays) % mod
}
```
