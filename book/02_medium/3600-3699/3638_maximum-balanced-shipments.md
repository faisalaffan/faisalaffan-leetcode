# 3638 — Maximum Balanced Shipments

## Deskripsi

**Soal:** [3638. Maximum Balanced Shipments](https://leetcode.com/problems/maximum-balanced-shipments/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxBalancedShipments(weight []int) int`

## Solusi Go

```go
package main

// LeetCode #3638: Maximum Balanced Shipments
// https://leetcode.com/problems/maximum-balanced-shipments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxBalancedShipments(weight []int) int {
	ans := 0
	mx := 0
	for _, x := range weight {
		if x > mx {
			mx = x
		}
		if x < mx {
			ans++
			mx = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(maxBalancedShipments([]int{2, 5, 1, 4, 3}))
	fmt.Println(maxBalancedShipments([]int{4, 4}))
	fmt.Println(maxBalancedShipments([]int{1, 3, 2, 4, 5, 2}))
}
```
