# 3657 — Find Loyal Customers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3657: Find Loyal Customers
// https://leetcode.com/problems/find-loyal-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int {
  // Membuat map (HashMap) — pencarian O(1)
	customerTotals := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	customerCounts := make(map[int]int)

	for _, p := range purchases {
		custID, amount := p[0], p[1]
		customerTotals[custID] += amount
		customerCounts[custID]++
	}

	var loyal []int
	for id := range customerTotals {
		if customerCounts[id] >= minPurchases && float64(customerTotals[id]) >= minAmount {
			loyal = append(loyal, id)
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(loyal)
	return loyal
}

func main() {
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 300}, {2, 150}, {1, 50}}, 2, 200))
	fmt.Println(findLoyalCustomers([][]int{{1, 50}, {1, 50}}, 2, 100))
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 200}}, 2, 100))
}
```
