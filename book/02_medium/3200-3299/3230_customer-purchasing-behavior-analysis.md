# 3230 — Customer Purchasing Behavior Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func customerPurchasingBehavior(purchases [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3230: Customer Purchasing Behavior Analysis
// https://leetcode.com/problems/customer-purchasing-behavior-analysis/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func customerPurchasingBehavior(purchases [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[int]int)
	for _, p := range purchases {
		counts[p[0]]++
	}

	var customers []int
	for id := range counts {
		customers = append(customers, id)
	}
  // Custom sort dengan comparator
	sort.Slice(customers, func(i, j int) bool {
		if counts[customers[i]] != counts[customers[j]] {
			return counts[customers[i]] > counts[customers[j]]
		}
		return customers[i] < customers[j]
	})

  // Alokasi slice integer
	ans := make([]int, len(customers))
	for i, id := range customers {
		ans[i] = id
	}
	return ans
}

func main() {
	fmt.Println(customerPurchasingBehavior([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 75}})) // Expected: [1 2 3]
	fmt.Println(customerPurchasingBehavior([][]int{{1, 10}, {2, 20}, {2, 30}}))             // Expected: [2 1]
}
```
