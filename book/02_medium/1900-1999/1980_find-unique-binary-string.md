# 1980 — Find Unique Binary String

## Deskripsi

**Soal:** [1980. Find Unique Binary String](https://leetcode.com/problems/find-unique-binary-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1980: Find Unique Binary String
// https://leetcode.com/problems/find-unique-binary-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindDifferentBinaryString([]string{"01", "10"}))
	fmt.Println(FindDifferentBinaryString([]string{"00", "01"}))
	fmt.Println(FindDifferentBinaryString([]string{"111", "011", "001"}))
}

// Time: O(n^2), Space: O(n)
func FindDifferentBinaryString(nums []string) string {
	n := len(nums)
  // Membuat map untuk pencarian O(1): key → value
	set := make(map[string]bool)
	for _, s := range nums {
		set[s] = true
	}

	// Generate candidates using Cantor diagonal argument
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = '0'
	}
	for {
		s := string(result)
		if !set[s] {
			return s
		}
		// Increment binary string
		j := n - 1
		for j >= 0 && result[j] == '1' {
			result[j] = '0'
			j--
		}
		if j < 0 {
			break
		}
		result[j] = '1'
	}
	return ""
}
```
