# 3521 — Find Product Recommendation Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindProductRecommendationPairs(products []int, target int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3521: Find Product Recommendation Pairs
// https://leetcode.com/problems/find-product-recommendation-pairs/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindProductRecommendationPairs([]int{1, 2, 3, 4, 5}, 5))
	// Test case 2
	fmt.Println("Test 2:", FindProductRecommendationPairs([]int{1, 1, 1, 1}, 2))
	// Test case 3
	fmt.Println("Test 3:", FindProductRecommendationPairs([]int{1, 2, 3}, 7))
}

func FindProductRecommendationPairs(products []int, target int) [][]int {
	// Find pairs that sum to target
	var result [][]int
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, p := range products {
		complement := target - p
		if seen[complement] {
			result = append(result, []int{complement, p})
		}
		seen[p] = true
	}
	if result == nil {
		return [][]int{}
	}
	return result
}
```
