# 1247 — Minimum Swaps To Make Strings Equal

## Deskripsi

**Soal:** [1247. Minimum Swaps To Make Strings Equal](https://leetcode.com/problems/minimum-swaps-to-make-strings-equal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumSwap(s1 string, s2 string) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1247: Minimum Swaps to Make Strings Equal
// https://leetcode.com/problems/minimum-swaps-to-make-strings-equal/
// Difficulty: Medium

// Count positions where s1[i] != s2[i].
// If odd count -> not possible.
// Count patterns: x_y (s1=x, s2=y) and y_x.
// Answer = x_y/2 + y_x/2 + (x_y%2)*2

// Time: O(n)
// Space: O(1)

func minimumSwap(s1 string, s2 string) int {
	n := len(s1)
	if n != len(s2) {
		return -1
	}

	xy, yx := 0, 0
	for i := 0; i < n; i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			xy++
		} else if s1[i] == 'y' && s2[i] == 'x' {
			yx++
		}
	}

	if (xy+yx)%2 == 1 {
		return -1
	}

	return xy/2 + yx/2 + (xy%2)*2
}

func main() {
	fmt.Printf("%d (expected: 1)\n", minimumSwap("xy", "yx"))
	fmt.Printf("%d (expected: 2)\n", minimumSwap("xx", "yy"))
	fmt.Printf("%d (expected: -1)\n", minimumSwap("xy", "xx"))
}
```
