# 3903 — Smallest Stable Index I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestStableIndexI(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3903: Smallest Stable Index I
// https://leetcode.com/problems/smallest-stable-index-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestStableIndexI([]int{5, 0, 1, 4}, 3))
	fmt.Println(SmallestStableIndexI([]int{3, 2, 1}, 1))
	fmt.Println(SmallestStableIndexI([]int{0}, 0))
}

// Time: O(n)
// Space: O(n)
func SmallestStableIndexI(nums []int, k int) int {
	n := len(nums)

  // Alokasi slice integer
	prefixMax := make([]int, n)
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

  // Alokasi slice integer
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if prefixMax[i]-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}
```
