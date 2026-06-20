# 3118 — Friday Purchase Iii

## Deskripsi

**Soal:** [3118. Friday Purchase Iii](https://leetcode.com/problems/friday-purchase-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func fridayPurchase(purchases [][]int) int64`

## Solusi Go

```go
package main

// LeetCode #3118: Friday Purchase III
// https://leetcode.com/problems/friday-purchase-iii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func fridayPurchase(purchases [][]int) int64 {
	total := int64(0)
	for _, p := range purchases {
		// p[0] = day of week (5 = Friday), p[1] = amount
		if p[0] == 5 {
			total += int64(p[1])
		}
	}
	return total
}

func main() {
	fmt.Println(fridayPurchase([][]int{{5, 100}, {1, 50}, {5, 200}, {3, 75}})) // Expected: 300
	fmt.Println(fridayPurchase([][]int{{2, 50}, {3, 100}}))                     // Expected: 0
}
```
