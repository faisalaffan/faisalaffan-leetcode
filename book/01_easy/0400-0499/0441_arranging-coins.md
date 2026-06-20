# 0441 — Arranging Coins

## Deskripsi

**Soal:** [0441. Arranging Coins](https://leetcode.com/problems/arranging-coins/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ArrangingCoins(n int) int`

## Solusi Go

```go
package main

// LeetCode #441: Arranging Coins
// https://leetcode.com/problems/arranging-coins/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func ArrangingCoins(n int) int {
	lo, hi := 1, n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		} else if sum < n {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi
}

func main() {
	fmt.Println(ArrangingCoins(5))
	fmt.Println(ArrangingCoins(8))
	fmt.Println(ArrangingCoins(1))
}
```
