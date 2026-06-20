# 0187 — Repeated Dna Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func findRepeatedDnaSequences(s string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #187: Repeated DNA Sequences
// https://leetcode.com/problems/repeated-dna-sequences/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

  // HashMap: O(1) lookup
	seen := make(map[string]int)
	result := []string{}

	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		seen[sub]++
		if seen[sub] == 2 {
			result = append(result, sub)
		}
	}

	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
```
