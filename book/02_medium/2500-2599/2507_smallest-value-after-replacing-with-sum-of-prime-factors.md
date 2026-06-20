# 2507 — Smallest Value After Replacing With Sum Of Prime Factors

## Deskripsi

**Soal:** [2507. Smallest Value After Replacing With Sum Of Prime Factors](https://leetcode.com/problems/smallest-value-after-replacing-with-sum-of-prime-factors/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(n) * iterations)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2507: Smallest Value After Replacing With Sum of Prime Factors
// https://leetcode.com/problems/smallest-value-after-replacing-with-sum-of-prime-factors/
// Difficulty: Medium
// Time: O(sqrt(n) * iterations) | Space: O(1)
// Replace n with sum of its prime factors (with multiplicity) until stable.

import "fmt"

func main() {
	fmt.Println(smallestValue(15)) // 5 (15=3*5 -> 8=2*2*2 -> 6=2*3 -> 5=prime)
	fmt.Println(smallestValue(4))  // 4 (4=2*2 -> 4, stable)
}

func smallestValue(n int) int {
	for {
		sum := primeFactorSum(n)
		if sum == n {
			return n
		}
		n = sum
	}
}

func primeFactorSum(n int) int {
	sum := 0
	// Factor 2
	for n%2 == 0 {
		sum += 2
		n /= 2
	}
	// Odd factors
	for f := 3; f*f <= n; f += 2 {
		for n%f == 0 {
			sum += f
			n /= f
		}
	}
	if n > 1 {
		sum += n
	}
	return sum
}
```
