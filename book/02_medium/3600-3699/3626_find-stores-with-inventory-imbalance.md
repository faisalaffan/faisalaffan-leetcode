# 3626 — Find Stores With Inventory Imbalance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindStoresWithInventoryImbalance(inventory []int, threshold int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3626: Find Stores with Inventory Imbalance
// https://leetcode.com/problems/find-stores-with-inventory-imbalance/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	inventory := []int{10, 5, 15, 3, 20}
	threshold := 10
	fmt.Println("Test 1:", FindStoresWithInventoryImbalance(inventory, threshold))
	// Test case 2
	inventory2 := []int{100, 50, 75}
	threshold2 := 20
	fmt.Println("Test 2:", FindStoresWithInventoryImbalance(inventory2, threshold2))
	// Test case 3
	inventory3 := []int{1, 1, 1}
	threshold3 := 0
	fmt.Println("Test 3:", FindStoresWithInventoryImbalance(inventory3, threshold3))
}

func FindStoresWithInventoryImbalance(inventory []int, threshold int) int {
	// Count stores where difference from average exceeds threshold
	if len(inventory) == 0 {
		return 0
	}
	sum := 0
	for _, v := range inventory {
		sum += v
	}
	avg := float64(sum) / float64(len(inventory))
	count := 0
	for _, v := range inventory {
		diff := float64(v) - avg
		if diff < 0 {
			diff = -diff
		}
		if diff > float64(threshold) {
			count++
		}
	}
	return count
}
```
