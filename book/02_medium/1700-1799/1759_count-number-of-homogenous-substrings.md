# 1759 — Count Number Of Homogenous Substrings

## Deskripsi

**Soal:** [1759. Count Number Of Homogenous Substrings](https://leetcode.com/problems/count-number-of-homogenous-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countHomogenous(s string) int`

## Solusi Go

```go
package main

// LeetCode #1759: Count Number of Homogenous Substrings
// https://leetcode.com/problems/count-number-of-homogenous-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func countHomogenous(s string) int {
	result := 0
	count := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		result = (result + count) % mod
	}
	return result
}

func main() {
	fmt.Println(countHomogenous("abbcccaa")) // Expected: 13
	fmt.Println(countHomogenous("xy"))        // Expected: 2
	fmt.Println(countHomogenous("zzzzz"))     // Expected: 15
}
```
