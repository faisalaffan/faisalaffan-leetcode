# 3145 — Find Products Of Elements Of Big Array

## Deskripsi

**Soal:** [3145. Find Products Of Elements Of Big Array](https://leetcode.com/problems/find-products-of-elements-of-big-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func cnt1(num int) int`

## Solusi Go

```go
package main

// LeetCode #3145: Find Products of Elements of Big Array
// https://leetcode.com/problems/find-products-of-elements-of-big-array/
// Difficulty: Hard
//
// The "powerful array" of x is the shortest sorted array of powers of two that
// sum to x (i.e., set bits of x). The "big array" is the concatenation of
// powerful arrays for all positive integers: [1, 2, 1, 2, 4, 1, 4, ...].
// Each query [from, to, mod] asks for the product of big array elements in that
// range modulo mod. Since every element is a power of 2, product = 2^(exponent_sum).
// Use binary search + bit counting to compute prefix exponent sums.

import (
	"fmt"
	"math/big"
	"sort"
)

// cnt1 returns total count of set bits in numbers 1..num (i.e., big array prefix length)
func cnt1(num int) int {
	if num <= 0 {
		return 0
	}
	res := 0
	for i := 0; 1<<uint(i) <= num; i++ {
		cycle := 1 << uint(i+1)
		cur := (num + 1) % cycle
		res += ((num + 1) / cycle) * (1 << uint(i))
		if cur > (1 << uint(i)) {
			res += cur - (1 << uint(i))
		}
	}
	return res
}

// acc0 returns total sum of bit-position exponents for set bits in 1..num
func acc0(num int) int {
	if num <= 0 {
		return 0
	}
	res := 0
	for i := 0; 1<<uint(i) <= num; i++ {
		cycle := 1 << uint(i+1)
		cur := (num + 1) % cycle
		res += ((num + 1) / cycle) * (1 << uint(i)) * i
		if cur > (1 << uint(i)) {
			res += (cur - (1 << uint(i))) * i
		}
	}
	return res
}

// prefixExpSum returns total exponent sum for big array elements [0..bound-1]
func prefixExpSum(bound int) int {
	if bound <= 0 {
		return 0
	}
	target := sort.Search(bound, func(n int) bool {
		return cnt1(n) >= bound
	})

	prevCnt := cnt1(target - 1)
	rest := bound - prevCnt
	expSum := acc0(target - 1)

	for i := 0; rest > 0; i++ {
		if target&(1<<uint(i)) != 0 {
			expSum += i
			rest--
		}
	}
	return expSum
}

func findProductsOfElementsOfBigArray(queries [][]int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for idx, q := range queries {
		from, to, mod := q[0], q[1], q[2]
		exp := prefixExpSum(to+1) - prefixExpSum(from)
		ans[idx] = int(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(exp)), big.NewInt(int64(mod))).Int64())
	}
	return ans
}

func main() {
	// Test case 1
	queries := [][]int{
		{1, 3, 1000000007},
		{5, 7, 1000000007},
	}
	fmt.Println("Test 1:", findProductsOfElementsOfBigArray(queries))

	// Test case 2: single element query
	queries2 := [][]int{{0, 0, 1000000007}}
	fmt.Println("Test 2:", findProductsOfElementsOfBigArray(queries2))
	// Expected: 1 (2^0 = 1)

	// Test case 3
	queries3 := [][]int{{2, 5, 1000000007}}
	fmt.Println("Test 3:", findProductsOfElementsOfBigArray(queries3))
}
```
