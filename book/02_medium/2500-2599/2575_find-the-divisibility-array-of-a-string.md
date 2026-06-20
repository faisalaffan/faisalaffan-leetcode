# 2575 — Find The Divisibility Array Of A String

## Deskripsi

**Soal:** [2575. Find The Divisibility Array Of A String](https://leetcode.com/problems/find-the-divisibility-array-of-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func divisibilityArray(word string, m int) []int`

## Solusi Go

```go
package main

// LeetCode #2575: Find the Divisibility Array of a String
// https://leetcode.com/problems/find-the-divisibility-array-of-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func divisibilityArray(word string, m int) []int {
	n := len(word)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
	rem := 0
	for i := 0; i < n; i++ {
		rem = (rem*10 + int(word[i]-'0')) % m
		if rem == 0 {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", divisibilityArray("998244353", 3))
	// Expected: [1,1,0,0,0,1,1,0,0]

	// Test case 2
	fmt.Println("Test 2:", divisibilityArray("1010", 10))
	// Expected: [0,1,0,1]

	// Test case 3
	fmt.Println("Test 3:", divisibilityArray("10", 5))
	// Expected: [0,0]
}
```
