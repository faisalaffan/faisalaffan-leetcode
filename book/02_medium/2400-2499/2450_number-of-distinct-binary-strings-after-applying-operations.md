# 2450 — Number Of Distinct Binary Strings After Applying Operations

## Deskripsi

**Soal:** [2450. Number Of Distinct Binary Strings After Applying Operations](https://leetcode.com/problems/number-of-distinct-binary-strings-after-applying-operations/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2450: Number of Distinct Binary Strings After Applying Operations
// https://leetcode.com/problems/number-of-distinct-binary-strings-after-applying-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Count distinct binary strings reachable by inverting any k-length substring.
// This is equivalent to 2^(number_of_free_variables).

import "fmt"

func main() {
	fmt.Println(distinctBinaryStrings("110", 2)) // 4
	fmt.Println(distinctBinaryStrings("10110", 5)) // 2
}

const MOD = 1000000007

func distinctBinaryStrings(s string, k int) int {
	n := len(s)
	// Number of reachable strings = 2^(n-k+1) if k > 0
	// Because each of the first n-k+1 bits can be independently flipped
	if k > n {
		return 1
	}
	pow := 1
	for i := 0; i < n-k+1; i++ {
		pow = (pow * 2) % MOD
	}
	return pow
}
```
