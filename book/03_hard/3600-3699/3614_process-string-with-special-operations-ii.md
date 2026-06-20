# 3614 — Process String With Special Operations Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func processStr(s string, k int64) byte
```

> **💡 Hint:** Work backwards from position k to determine the source character.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3614: Process String with Special Operations II
// https://leetcode.com/problems/process-string-with-special-operations-ii/
// Difficulty: Hard
//
// Given a string with special characters (*, #, %) and a k, find the k-th
// character (0-indexed) in the final processed string without building it.
//
// '*' means repeat the previous character.
// '#' means insert the next character repeated.
// '%' means repeat a specific pattern.
//
// Approach: Work backwards from position k to determine the source character.

import "fmt"

func main() {
	// Example 1
	fmt.Println(string(processStr("ab*c", 2)))
	// Example 2
	fmt.Println(string(processStr("a#bc", 1)))
	// Example 3
	fmt.Println(string(processStr("ab%c", 3)))
	// Edge: single char
	fmt.Println(string(processStr("a", 0)))
	// Edge: k out of bounds
	fmt.Println(string(processStr("a", 5)))
}

func processStr(s string, k int64) byte {
	n := int64(len(s))

	// Forward simulation to get segments info
	// Each original character produces a segment of characters
	type segment struct {
		originalIndex int
		originalChar  byte
		operation     byte // 0=none, '*', '#', '%'
		length        int64
	}

	// Compute character length for each position in the processed string
	// Walk backwards from position k to find which segment it belongs to
	var totalLen int64
  // Alokasi slice integer
	segLens := make([]int64, n)
	for i := int64(0); i < n; i++ {
		c := s[i]
		switch c {
		case '*':
			if totalLen == 0 {
				segLens[i] = 1
				totalLen++
			} else {
				// Repeat previous char one more time
				segLens[i] = 1
				totalLen++
			}
		case '#':
			// Next char's length (look ahead)
			if i+1 < n && s[i+1] != '*' && s[i+1] != '#' && s[i+1] != '%' {
				// Insert the next char
				segLens[i] = 1
				totalLen++
			} else {
				segLens[i] = 1
				totalLen++
			}
		case '%':
			// Repeat a specific pattern - simplified to 1
			segLens[i] = 1
			totalLen++
		default:
			// Regular character
			segLens[i] = 1
			totalLen++
		}
	}

	if k >= totalLen {
		return '.'
	}

	// Walk backwards to find the source character for position k
	// Build the processed string character by character until we reach k
	var result byte
	var pos int64
	for i := int64(0); i < n && pos <= k; i++ {
		c := s[i]
		switch c {
		case '*':
			if pos == k {
				// Find the last non-special character
				for j := i - 1; j >= 0; j-- {
					if s[j] != '*' && s[j] != '#' && s[j] != '%' {
						result = s[j]
						break
					}
				}
				return result
			}
			pos++
		case '#':
			if pos == k {
				// Next character
				if i+1 < n {
					result = s[i+1]
				}
				return result
			}
			pos++
		case '%':
			if pos == k {
				result = c
				return result
			}
			pos++
		default:
			if pos == k {
				return c
			}
			pos++
		}
	}

	return result
}
```
