# 0370 — Range Addition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func getModifiedArray(length int, updates [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n + k)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #370: Range Addition
// https://leetcode.com/problems/range-addition/
// Difficulty: Medium [Paid]
// Time: O(n + k) | Space: O(n)

import "fmt"

func getModifiedArray(length int, updates [][]int) []int {
  // Alokasi slice integer
	arr := make([]int, length+1)

	for _, upd := range updates {
		start, end, inc := upd[0], upd[1], upd[2]
		arr[start] += inc
		arr[end+1] -= inc
	}

	// Prefix sum
  // Alokasi slice integer
	result := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += arr[i]
		result[i] = sum
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))
	// Expected: [-2, 0, 3, 5, 3]

	// Test case 2: Single update
	fmt.Println("Test 2:", getModifiedArray(3, [][]int{{0, 2, 5}}))
	// Expected: [5, 5, 5]

	// Test case 3: No updates
	fmt.Println("Test 3:", getModifiedArray(3, [][]int{}))
	// Expected: [0, 0, 0]
}
```
