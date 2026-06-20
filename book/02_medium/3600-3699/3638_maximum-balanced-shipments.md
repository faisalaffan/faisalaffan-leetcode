# 3638 — Maximum Balanced Shipments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxBalancedShipments(weight []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
