# 3597 — Partition String

## Deskripsi

**Soal:** [3597. Partition String](https://leetcode.com/problems/partition-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[byte]bool)
  // Loop standar: indeks 0 sampai n-1
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
