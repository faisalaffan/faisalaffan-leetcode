# 3790 — Smallest All Ones Multiple

## Deskripsi

**Soal:** [3790. Smallest All Ones Multiple](https://leetcode.com/problems/smallest-all-ones-multiple/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(k)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func smallestAllOnesMultiple(k int) int`

## Solusi Go

```go
package main

// LeetCode #3790: Smallest All-Ones Multiple
// https://leetcode.com/problems/smallest-all-ones-multiple/
// Difficulty: Medium
// Time: O(k) | Space: O(1)

import "fmt"

func smallestAllOnesMultiple(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	rem := 0
	for n := 1; n <= k; n++ {
		rem = (rem*10 + 1) % k
		if rem == 0 {
			return n
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestAllOnesMultiple(3))
	fmt.Println(smallestAllOnesMultiple(7))
	fmt.Println(smallestAllOnesMultiple(2))
}
```
