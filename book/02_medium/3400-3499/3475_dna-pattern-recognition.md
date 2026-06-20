# 3475 — Dna Pattern Recognition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func DnaPatternRecognition(dna string, pattern string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3475: DNA Pattern Recognition
// https://leetcode.com/problems/dna-pattern-recognition/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", DnaPatternRecognition("ACGTACGT", "ACGT"))
	// Test case 2
	fmt.Println("Test 2:", DnaPatternRecognition("AAAA", "AA"))
	// Test case 3
	fmt.Println("Test 3:", DnaPatternRecognition("ACGT", "TGCA"))
}

func DnaPatternRecognition(dna string, pattern string) int {
	// Count occurrences of pattern in DNA string
	if len(pattern) == 0 {
		return 0
	}
	count := 0
	for i := 0; i <= len(dna)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			if dna[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}
```
