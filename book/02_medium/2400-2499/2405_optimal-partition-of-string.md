# 2405 — Optimal Partition Of String

## Deskripsi

**Soal:** [2405. Optimal Partition Of String](https://leetcode.com/problems/optimal-partition-of-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Greedy (pemilihan optimal lokal)

## Solusi Go

```go
package main

// LeetCode #2405: Optimal Partition of String
// https://leetcode.com/problems/optimal-partition-of-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new partition when duplicate char found.

import "fmt"

func main() {
	fmt.Println(partitionString("abacaba")) // 4
	fmt.Println(partitionString("ssssss"))  // 6
}

func partitionString(s string) int {
	ans := 1
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[byte]bool)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			ans++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return ans
}
```
