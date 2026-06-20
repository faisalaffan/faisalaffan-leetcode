# 0187 — Repeated Dna Sequences

## Deskripsi

**Soal:** [0187. Repeated Dna Sequences](https://leetcode.com/problems/repeated-dna-sequences/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findRepeatedDnaSequences(s string) []string`

## Solusi Go

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

  // Membuat map untuk pencarian O(1): key → value
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
