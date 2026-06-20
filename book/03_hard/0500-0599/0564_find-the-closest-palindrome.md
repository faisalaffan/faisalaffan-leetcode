# 0564 — Find The Closest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func nearestPalindromic(n string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #564: Find the Closest Palindrome
// https://leetcode.com/problems/find-the-closest-palindrome/
// Difficulty: Hard

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nearestPalindromic("123")) // Expected: "121"
}

func nearestPalindromic(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	if num <= 10 {
		return strconv.FormatInt(num-1, 10)
	}
	if num == 11 {
		return "9"
	}

	length := len(n)
	halfStr := n[:(length+1)/2]
	half, _ := strconv.ParseInt(halfStr, 10, 64)

	candidates := []int64{}

	// Generate candidates by mirroring
	for _, d := range []int64{-1, 0, 1} {
		h := half + d
		s := strconv.FormatInt(h, 10)
		var pal string
		if length%2 == 0 {
			pal = s + reverse(s)
		} else {
			pal = s + reverse(s[:len(s)-1])
		}
		if val, err := strconv.ParseInt(pal, 10, 64); err == nil {
			candidates = append(candidates, val)
		}
	}

	// Edge cases: 10^(len-1)-1 (9, 99, 999...) and 10^len+1 (101, 1001...)
	small, _ := strconv.ParseInt("9"+string(make([]byte, length-1)), 10, 64)
	large, _ := strconv.ParseInt("1"+string(make([]byte, length))+"1", 10, 64)
	candidates = append(candidates, small)
	candidates = append(candidates, large)

	var best int64 = 0
	for _, c := range candidates {
		if c == num {
			continue
		}
		diff := c - num
		if diff < 0 {
			diff = -diff
		}
		bestDiff := best - num
		if bestDiff < 0 {
			bestDiff = -bestDiff
		}
		if best == 0 || diff < bestDiff || (diff == bestDiff && c < best) {
			best = c
		}
	}
	return strconv.FormatInt(best, 10)
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

```
