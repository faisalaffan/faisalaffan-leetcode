# 0386 — Lexicographical Numbers

## Deskripsi

**Soal:** [0386. Lexicographical Numbers](https://leetcode.com/problems/lexicographical-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (excluding output)

**Algoritma:** —

**Fungsi Solusi:** `func lexicalOrder(n int) []int`

## Solusi Go

```go
package main

// LeetCode #386: Lexicographical Numbers
// https://leetcode.com/problems/lexicographical-numbers/
// Difficulty: Medium
// Time: O(n) | Space: O(1) (excluding output)

import "fmt"

func lexicalOrder(n int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, n)
	cur := 1

	for len(result) < n {
		result = append(result, cur)

		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur+1 > n {
				cur /= 10
			}
			cur++
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lexicalOrder(13))
	// Expected: [1,10,11,12,13,2,3,4,5,6,7,8,9]

	// Test case 2
	fmt.Println("Test 2:", lexicalOrder(2))
	// Expected: [1,2]

	// Test case 3
	fmt.Println("Test 3:", lexicalOrder(25))
	// Expected: [1,10,11,12,13,14,15,16,17,18,19,2,20,21,22,23,24,25,3,4,5,6,7,8,9]
}
```
