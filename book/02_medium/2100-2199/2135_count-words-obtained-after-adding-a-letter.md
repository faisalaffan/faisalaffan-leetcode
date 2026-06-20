# 2135 — Count Words Obtained After Adding A Letter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func wordCount(startWords []string, targetWords []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Bitmask

**Kompleksitas Waktu:** O(n * L)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2135: Count Words Obtained After Adding a Letter
// https://leetcode.com/problems/count-words-obtained-after-adding-a-letter/
// Difficulty: Medium
// Time: O(n * L) | Space: O(n)

import "fmt"

func wordCount(startWords []string, targetWords []string) int {
	// Convert start words to bitmasks
  // Membuat map (HashMap) — pencarian O(1)
	startSet := make(map[int]bool)
	for _, w := range startWords {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		startSet[mask] = true
	}

	count := 0
	for _, w := range targetWords {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		// Try removing each character
		for _, c := range w {
			bit := 1 << (c - 'a')
			if startSet[mask^bit] {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", wordCount([]string{"ab", "a"}, []string{"abc", "abcd"}))
	// Expected: 1
}
```
