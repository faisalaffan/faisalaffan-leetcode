# 1835 — Find Xor Sum Of All Pairs Bitwise And

## Deskripsi

**Soal:** [1835. Find Xor Sum Of All Pairs Bitwise And](https://leetcode.com/problems/find-xor-sum-of-all-pairs-bitwise-and/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Bit Manipulation (operasi bitwise)

> **Ide Kunci:** Bit Manipulation.

## Solusi Go

```go
package main

// LeetCode #1835: Find XOR Sum of All Pairs Bitwise AND
// https://leetcode.com/problems/find-xor-sum-of-all-pairs-bitwise-and/
// Difficulty: Hard
//
// Approach: Bit Manipulation.
//   The XOR sum of all pairwise ANDs can be simplified:
//     For each bit position b:
//       Let c1 = count of arr1 elements with bit b set
//       Let c2 = count of arr2 elements with bit b set
//       Pairs with (arr1[i] & arr2[j]) having bit b set = c1 * c2
//       Bit b is set in final result iff c1 * c2 is odd
//     Since c1 * c2 is odd iff both c1 and c2 are odd, and parity of
//     count == XOR reduction:
//       result = (xor of all arr1) & (xor of all arr2)

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", xorAllNums([]int{1, 2, 3}, []int{5, 6, 7}))
	// Expected: 0
	// xor1 = 1^2^3 = 0, xor2 = 5^6^7 = 4, result = 0 & 4 = 0

	// Example 2
	fmt.Println("Example 2:", xorAllNums([]int{0, 1, 2}, []int{3, 4, 5}))
	// Expected: 4
	// xor1 = 0^1^2 = 3, xor2 = 3^4^5 = 2, result = 3 & 2 = 2? Let's compute:
	// 0&3 ^ 0&4 ^ 0&5 ^ 1&3 ^ 1&4 ^ 1&5 ^ 2&3 ^ 2&4 ^ 2&5
	// = 0 ^ 0 ^ 0 ^ 1 ^ 0 ^ 1 ^ 2 ^ 0 ^ 0 = 1^1^2 = 2
	// Hmm, expected might be different. Let's check with example.

	// Actually from LeetCode: [0,1,2] and [3,4,5]
	// All pairs:
	// 0&3=0, 0&4=0, 0&5=0
	// 1&3=1, 1&4=0, 1&5=1
	// 2&3=2, 2&4=0, 2&5=0
	// XOR sum = 0^0^0^1^0^1^2^0^0 = 2
	fmt.Println("  -> xor1=0^1^2=3, xor2=3^4^5=2, result=3&2=2")

	// Edge case: single elements each
	fmt.Println("Edge (single each):", xorAllNums([]int{1}, []int{2}))
	// Expected: 0 (1&2 = 0)

	// All zeros
	fmt.Println("Edge (zeros):", xorAllNums([]int{0, 0}, []int{0, 0}))
	// Expected: 0
}

func xorAllNums(arr1 []int, arr2 []int) int {
	xor1 := 0
	for _, v := range arr1 {
		xor1 ^= v
	}
	xor2 := 0
	for _, v := range arr2 {
		xor2 ^= v
	}
	return xor1 & xor2
}

// Stub kept for compatibility with the repo scaffold.
func FindXorSumOfAllPairsBitwiseAnd() any {
	return xorAllNums([]int{1, 2, 3}, []int{5, 6, 7})
}
```
