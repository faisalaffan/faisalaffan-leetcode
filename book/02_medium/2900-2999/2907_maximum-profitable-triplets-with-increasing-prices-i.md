# 2907 — Maximum Profitable Triplets With Increasing Prices I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxProfit(prices []int, profits []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2907: Maximum Profitable Triplets With Increasing Prices I
// https://leetcode.com/problems/maximum-profitable-triplets-with-increasing-prices-i/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{10, 20, 30}, []int{1, 2, 3}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4}, []int{5, 4, 3, 2}))
	fmt.Println(maxProfit([]int{3, 2, 1}, []int{10, 20, 30}))
}

func maxProfit(prices []int, profits []int) int {
	n := len(prices)
	ans := -1
	for j, x := range profits {
		left, right := 0, 0
		for i := 0; i < j; i++ {
			if prices[i] < prices[j] {
				if profits[i] > left {
					left = profits[i]
				}
			}
		}
		for k := j + 1; k < n; k++ {
			if prices[j] < prices[k] {
				if profits[k] > right {
					right = profits[k]
				}
			}
		}
		if left > 0 && right > 0 {
			if left+x+right > ans {
				ans = left + x + right
			}
		}
	}
	return ans
}
```
