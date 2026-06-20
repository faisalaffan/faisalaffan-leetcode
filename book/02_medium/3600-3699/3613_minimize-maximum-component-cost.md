# 3613 — Minimize Maximum Component Cost

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimizeMaximumComponentCost(costs []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3613: Minimize Maximum Component Cost
// https://leetcode.com/problems/minimize-maximum-component-cost/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	costs := []int{1, 2, 3, 4, 5}
	k := 2
	fmt.Println("Test 1:", MinimizeMaximumComponentCost(costs, k))
	// Test case 2
	costs2 := []int{10, 20, 30}
	k2 := 1
	fmt.Println("Test 2:", MinimizeMaximumComponentCost(costs2, k2))
	// Test case 3
	costs3 := []int{5}
	k3 := 3
	fmt.Println("Test 3:", MinimizeMaximumComponentCost(costs3, k3))
}

func MinimizeMaximumComponentCost(costs []int, k int) int {
	if len(costs) == 0 {
		return 0
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(costs)
	// Binary search for minimal possible maximum component cost
	sum := 0
	for _, c := range costs {
		sum += c
	}
	left, right := costs[len(costs)-1], sum
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if canPartition(costs, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canPartition(costs []int, k, maxCost int) bool {
	components := 0
	sum := 0
	for _, c := range costs {
		if sum+c > maxCost {
			components++
			sum = c
		} else {
			sum += c
		}
	}
	if sum > 0 {
		components++
	}
	return components <= k
}
```
