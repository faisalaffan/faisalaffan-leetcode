# 3007 — Maximum Number That Sum Of The Prices Is Less Than Or Equal To K

## Deskripsi

**Soal:** [3007. Maximum Number That Sum Of The Prices Is Less Than Or Equal To K](https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log^2 n)  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3007: Maximum Number That Sum of the Prices Is Less Than or Equal to K
// https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(log^2 n) | Space: O(log n)

import "fmt"

func main() {
	fmt.Println(findMaximumNumber(9, 1))
	fmt.Println(findMaximumNumber(7, 2))
}

func findMaximumNumber(k int64, x int) int64 {
	var lo, hi int64 = 0, 1e18
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if countPrice(mid, x) <= k {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func countPrice(num int64, x int) int64 {
	var ans int64
	n := num
	for bit := x - 1; int64(1<<bit) <= n; bit += x {
		cycle := int64(1 << (bit + 1))
		full := n / cycle
		ans += full * int64(1<<bit)
		rem := n % cycle
		ans += max(0, rem-int64(1<<bit)+1)
	}
	return ans
}
```
