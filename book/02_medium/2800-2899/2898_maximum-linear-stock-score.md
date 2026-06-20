# 2898 — Maximum Linear Stock Score

## Deskripsi

**Soal:** [2898. Maximum Linear Stock Score](https://leetcode.com/problems/maximum-linear-stock-score/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func MaximumLinearStockScore(prices []int) int64`

## Solusi Go

```go
package main

// LeetCode #2898: Maximum Linear Stock Score
// https://leetcode.com/problems/maximum-linear-stock-score/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumLinearStockScore(prices []int) int64 {
	// For each stock, score = sum of prices where prices[i] - i is same
  // Membuat map untuk pencarian O(1): key → value
	score := make(map[int]int64)
	var best int64

	for i, p := range prices {
		key := p - i
		score[key] += int64(p)
		if score[key] > best {
			best = score[key]
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumLinearStockScore([]int{1, 2, 3, 4}))
	fmt.Println(MaximumLinearStockScore([]int{2, 1, 3}))
}
```
