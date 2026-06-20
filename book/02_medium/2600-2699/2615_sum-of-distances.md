# 2615 — Sum Of Distances

## Deskripsi

**Soal:** [2615. Sum Of Distances](https://leetcode.com/problems/sum-of-distances/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func distance(nums []int) []int64`

## Solusi Go

```go
package main

// LeetCode #2615: Sum of Distances
// https://leetcode.com/problems/sum-of-distances/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func distance(nums []int) []int64 {
	n := len(nums)
	// Group indices by value
  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[int][]int)
	for i, v := range nums {
		groups[v] = append(groups[v], i)
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int64, n)
	for _, indices := range groups {
		m := len(indices)
  // Membuat slice untuk menyimpan hasil
		prefix := make([]int64, m+1)
		for i, idx := range indices {
			prefix[i+1] = prefix[i] + int64(idx)
		}
		for i, idx := range indices {
			leftCount := int64(i)
			rightCount := int64(m - i - 1)
			leftSum := prefix[i]
			rightSum := prefix[m] - prefix[i+1]
			ans[idx] = int64(idx)*leftCount - leftSum + rightSum - int64(idx)*rightCount
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", distance([]int{1, 3, 1, 1, 2}))
	// Expected: [5,0,3,4,0]

	// Test case 2
	fmt.Println("Test 2:", distance([]int{0, 5, 3}))
	// Expected: [0,0,0]

	// Test case 3
	fmt.Println("Test 3:", distance([]int{1, 1, 1, 1}))
	// Expected: [6,4,4,6]
}
```
