# 3867 — Sum Of Gcd Of Formed Pairs

## Deskripsi

**Soal:** [3867. Sum Of Gcd Of Formed Pairs](https://leetcode.com/problems/sum-of-gcd-of-formed-pairs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log M)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func gcd(a, b int) int`

> **Ide Kunci:** Build prefixGcd array where prefixGcd[i] = gcd(nums[i], max(nums[0..i])).

## Solusi Go

```go
package main

// LeetCode #3867: Sum of GCD of Formed Pairs
// https://leetcode.com/problems/sum-of-gcd-of-formed-pairs/
// Difficulty: Medium
// Time: O(N log M) | Space: O(N)
// Approach: Build prefixGcd array where prefixGcd[i] = gcd(nums[i], max(nums[0..i])).
// Sort, pair smallest with largest, sum gcd of each pair.

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func SumOfGcdOfFormedPairs(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefixGcd := make([]int, n)
	mx := 0
	for i, v := range nums {
		if v > mx {
			mx = v
		}
		prefixGcd[i] = gcd(v, mx)
	}

	sort.Ints(prefixGcd)

	ans := 0
	for i := 0; i < n/2; i++ {
		ans += gcd(prefixGcd[i], prefixGcd[n-1-i])
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(SumOfGcdOfFormedPairs([]int{2, 6, 4})) // Expected: 2

	// Example 2
	fmt.Println(SumOfGcdOfFormedPairs([]int{3, 6, 2, 8})) // Expected: 5

	// Extra
	fmt.Println(SumOfGcdOfFormedPairs([]int{1})) // Expected: 0
}
```
