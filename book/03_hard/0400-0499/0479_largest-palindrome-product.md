# 0479 — Largest Palindrome Product

## Deskripsi

**Soal:** [0479. Largest Palindrome Product](https://leetcode.com/problems/largest-palindrome-product/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Enumerate palindromes in descending order by building the first

## Solusi Go

```go
package main

// LeetCode #479: Largest Palindrome Product
// https://leetcode.com/problems/largest-palindrome-product/
// Difficulty: Hard
// Approach: Enumerate palindromes in descending order by building the first
// half and mirroring it. Check if the palindrome has a factor in the n-digit range.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("479 - Largest Palindrome Product")

	for n := 1; n <= 8; n++ {
		fmt.Printf("n=%d -> %d\n", n, largestPalindrome(n))
	}
	// n=1->9, n=2->987, n=3->123, n=4->597, n=5->677, n=6->1218, n=7->877, n=8->475
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	maxNum := int(math.Pow10(n)) - 1
	minNum := int(math.Pow10(n - 1))

	// Enumerate palindromes by iterating over the first half
	// Build palindrome: half + reverse(half) = an even-digit palindrome
	for half := maxNum; half >= 1; half-- {
		// Build the palindrome
		pal := int64(half)
		for temp := half; temp > 0; temp /= 10 {
			pal = pal*10 + int64(temp%10)
		}

		// Check if this palindrome is a product of two n-digit numbers
		// Only need to check factors up to sqrt(pal)
		for factor := int64(maxNum); factor*factor >= pal; factor-- {
			if pal%factor == 0 {
				other := pal / factor
				if other >= int64(minNum) && other <= int64(maxNum) {
					return int(pal % 1337)
				}
			}
		}
	}

	return 0
}
```
