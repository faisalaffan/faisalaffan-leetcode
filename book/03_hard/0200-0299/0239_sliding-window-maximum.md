# 0239 — Sliding Window Maximum

## Deskripsi

**Soal:** [0239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

**Fungsi Solusi:** `func maxSlidingWindow(nums []int, k int) []int`

## Solusi Go

```go
package main

// LeetCode #239: Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/
// Difficulty: Hard

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

  // Membuat slice untuk menyimpan hasil
	deque := make([]int, 0) // stores indices
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, len(nums)-k+1)

	for i, num := range nums {
		// Remove indices outside the window (from front)
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove from back while current num is larger (maintain decreasing order)
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// First valid window starts at index k-1
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
```
