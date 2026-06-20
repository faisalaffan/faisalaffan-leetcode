# 1405 — Longest Happy String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestDiverseString(a int, b int, c int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(a+b+c) - building the result string  
**Kompleksitas Ruang:** O(1) - constant extra space

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1405: Longest Happy String
// https://leetcode.com/problems/longest-happy-string/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(longestDiverseString(1, 1, 7)) // "ccaccbcc" or "ccbccacc"

	// Test case 2
	fmt.Println(longestDiverseString(7, 1, 0)) // "aabaa"

	// Test case 3
	fmt.Println(longestDiverseString(0, 8, 11)) // "ccbccbbccbbccbbccbc"
}

type charCount struct {
	count int
	char  byte
}

// Time: O(a+b+c) - building the result string
// Space: O(1) - constant extra space
func longestDiverseString(a int, b int, c int) string {
	pairs := []charCount{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	result := make([]byte, 0, a+b+c)

	for {
		// Sort by remaining count descending
  // Custom sort dengan comparator
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].count > pairs[j].count
		})

		placed := false
		for i := 0; i < 3; i++ {
			if pairs[i].count == 0 {
				break
			}
			n := len(result)
			// Check if we can place this character
			if n >= 2 && result[n-1] == pairs[i].char && result[n-2] == pairs[i].char {
				continue
			}
			result = append(result, pairs[i].char)
			pairs[i].count--
			placed = true
			break
		}

		if !placed {
			break
		}
	}

	return string(result)
}
```
