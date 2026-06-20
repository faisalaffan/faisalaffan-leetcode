# 3344 — Maximum Sized Array

## Deskripsi

**Soal:** [3344. Maximum Sized Array](https://leetcode.com/problems/maximum-sized-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(mx^2) Space: O(mx)  
**Kompleksitas Ruang:** O(mx)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3344: Maximum Sized Array
// https://leetcode.com/problems/maximum-sized-array/
// Difficulty: Medium [Paid]
// Time: O(mx^2) Space: O(mx)

import "fmt"

func main() {
	fmt.Println(maxSizedArray(1))  // 1
	fmt.Println(maxSizedArray(10)) // 2
	fmt.Println(maxSizedArray(0))  // 1
}

func maxSizedArray(s int) int {
	// Sum_{i=0}^{n-1} i * Sum_{j=0}^{n-1} Sum_{k=0}^{n-1} (j|k)
	// = (n-1)*n/2 * orSum(n)
	// Upper bound: sum ~ (n-1)^5/4, s <= 10^15 => n <= ~1330
	mx := 1335

	// Precompute orSum[n] = sum of (j|k) for all 0 <= j,k < n
  // Membuat slice untuk menyimpan hasil
	orSum := make([]int64, mx+1)
	for n := 1; n <= mx; n++ {
		orSum[n] = orSum[n-1]
		j := n - 1
		for k := 0; k < j; k++ {
			val := int64(j | k)
			orSum[n] += 2 * val
		}
		orSum[n] += int64(j | j)
	}

	ans := 1
	for n := 1; n <= mx; n++ {
		sumI := int64(n-1) * int64(n) / 2
		total := orSum[n] * sumI
		if total <= int64(s) {
			ans = n
		} else {
			break
		}
	}
	return ans
}
```
