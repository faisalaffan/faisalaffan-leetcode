# 3220 — Odd And Even Transactions

## Deskripsi

**Soal:** [3220. Odd And Even Transactions](https://leetcode.com/problems/odd-and-even-transactions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func oddAndEvenTransactions(transactions [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #3220: Odd and Even Transactions
// https://leetcode.com/problems/odd-and-even-transactions/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func oddAndEvenTransactions(transactions [][]int) []int {
	oddSum, evenSum := 0, 0
	for _, t := range transactions {
		amount := t[1]
		if amount%2 == 0 {
			evenSum += amount
		} else {
			oddSum += amount
		}
	}
	return []int{oddSum, evenSum}
}

func main() {
	fmt.Println(oddAndEvenTransactions([][]int{{1, 10}, {2, 15}, {3, 20}})) // Expected: [25 30]
	fmt.Println(oddAndEvenTransactions([][]int{{1, 1}, {2, 2}, {3, 3}}))    // Expected: [4 2]
}
```
