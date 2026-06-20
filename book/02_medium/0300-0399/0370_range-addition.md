# 0370 — Range Addition

## Deskripsi

**Soal:** [0370. Range Addition](https://leetcode.com/problems/range-addition/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + k)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func getModifiedArray(length int, updates [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #370: Range Addition
// https://leetcode.com/problems/range-addition/
// Difficulty: Medium [Paid]
// Time: O(n + k) | Space: O(n)

import "fmt"

func getModifiedArray(length int, updates [][]int) []int {
  // Membuat slice untuk menyimpan hasil
	arr := make([]int, length+1)

	for _, upd := range updates {
		start, end, inc := upd[0], upd[1], upd[2]
		arr[start] += inc
		arr[end+1] -= inc
	}

	// Prefix sum
  // Membuat slice untuk menyimpan hasil
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
