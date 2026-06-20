# 2222 — Number Of Ways To Select Buildings

## Deskripsi

**Soal:** [2222. Number Of Ways To Select Buildings](https://leetcode.com/problems/number-of-ways-to-select-buildings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfWays(s string) int64`

## Solusi Go

```go
package main

// LeetCode #2222: Number of Ways to Select Buildings
// https://leetcode.com/problems/number-of-ways-to-select-buildings/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfWays(s string) int64 {
	var total0, total1 int64 = 0, 0
	for _, ch := range s {
		if ch == '0' {
			total0++
		} else {
			total1++
		}
	}

	var ways, prefix0, prefix1 int64 = 0, 0, 0
	for _, ch := range s {
		if ch == '0' {
			ways += prefix1 * (total1 - prefix1)
			prefix0++
		} else {
			ways += prefix0 * (total0 - prefix0)
			prefix1++
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(numberOfWays("001101"))
	// Expected: 6

	// Test case 2
	fmt.Println(numberOfWays("11100"))
	// Expected: 0

	// Test case 3
	fmt.Println(numberOfWays("0001100100"))
	// Expected: 12
}
```
