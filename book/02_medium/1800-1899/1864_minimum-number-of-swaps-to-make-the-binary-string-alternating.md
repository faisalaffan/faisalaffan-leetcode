# 1864 — Minimum Number Of Swaps To Make The Binary String Alternating

## Deskripsi

**Soal:** [1864. Minimum Number Of Swaps To Make The Binary String Alternating](https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1864: Minimum Number of Swaps to Make the Binary String Alternating
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwaps("111000"))
	fmt.Println(MinSwaps("010"))
	fmt.Println(MinSwaps("1110"))
}

// Time: O(n), Space: O(1)
func MinSwaps(s string) int {
	n := len(s)
	ones := 0
	zeros := 0
	for _, c := range s {
		if c == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if abs(ones-zeros) > 1 {
		return -1
	}

	// Count mismatches when starting with '0' and starting with '1'
	mismatch0 := 0 // pattern: 010101...
	mismatch1 := 0 // pattern: 101010...
	for i, c := range s {
		if i%2 == 0 {
			if c == '1' {
				mismatch0++
			} else {
				mismatch1++
			}
		} else {
			if c == '0' {
				mismatch0++
			} else {
				mismatch1++
			}
		}
	}

	// Each swap fixes 2 mismatches
	if n%2 == 0 {
		return min(mismatch0/2, mismatch1/2)
	}
	// For odd length, only one pattern is valid
	if zeros > ones {
		return mismatch0 / 2
	}
	return mismatch1 / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
