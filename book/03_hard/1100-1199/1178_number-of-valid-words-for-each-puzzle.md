# 1178 — Number Of Valid Words For Each Puzzle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func wordMask(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Bitmask, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1178: Number of Valid Words for Each Puzzle
// https://leetcode.com/problems/number-of-valid-words-for-each-puzzle/
// Difficulty: Hard

import "fmt"

// wordMask converts a word to a bitmask of its letters (26 bits).
func wordMask(s string) int {
	mask := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	// Count words by mask
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, w := range words {
		m := wordMask(w)
		// Optimization: skip words with > 2^7 bits (not subset of any puzzle with 7 letters)
		// Actually just count normally
		freq[m]++
	}

  // Alokasi slice integer
	result := make([]int, len(puzzles))
	for i, p := range puzzles {
		pmask := wordMask(p)
		firstBit := 1 << (p[0] - 'a')
		count := 0

		// Enumerate all subsets of pmask that include the first letter
		sub := pmask
		for sub > 0 {
			if (sub & firstBit) != 0 {
				count += freq[sub]
			}
			sub = (sub - 1) & pmask
		}
		// Also check empty subset? No, word must have at least one letter (first letter)
		// Actually the subset loop above only goes through non-zero subsets of pmask
		// that include firstBit. But we already covered that via the loop.
		result[i] = count
	}
	return result
}

func main() {
	// Test case
	words := []string{"aaaa", "asas", "able", "ability", "actt", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
	fmt.Println(findNumOfValidWords(words, puzzles))
	// Expected: [1, 1, 3, 2, 4, 0]

	// Test case 2: empty
	fmt.Println(findNumOfValidWords([]string{}, []string{"a"})) // [0]
}
```
