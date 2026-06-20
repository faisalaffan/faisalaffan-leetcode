# 3597 — Partition String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func PartitionString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3597: Partition String
// https://leetcode.com/problems/partition-string/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionString("abac"))
	// Test case 2
	fmt.Println("Test 2:", PartitionString("aaaa"))
	// Test case 3
	fmt.Println("Test 3:", PartitionString("abc"))
}

func PartitionString(s string) int {
	// Partition into substrings with unique characters
	// Use greedy: start new partition when duplicate found
	count := 1
  // HashMap: O(1) lookup
	seen := make(map[byte]bool)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			count++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return count
}
```
