# 3835 — Count Subarrays With Cost Less Than Or Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSubarraysWithCostLessThanOrEqualToK(nums []int, k int) int
```

> **💡 Hint:** Sliding window with two monotonic deques for max and min.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3835: Count Subarrays With Cost Less Than or Equal to K
// https://leetcode.com/problems/count-subarrays-with-cost-less-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Sliding window with two monotonic deques for max and min.
// cost = (max - min) * length, expand right, shrink left when cost > k.

import "fmt"

func CountSubarraysWithCostLessThanOrEqualToK(nums []int, k int) int {
	n := len(nums)
	ans := 0
	left := 0

	// Monotonic deques for max (decreasing) and min (increasing)
  // Alokasi slice integer
	maxQ := make([]int, 0) // indices, values decreasing
  // Alokasi slice integer
	minQ := make([]int, 0) // indices, values increasing

	for right := 0; right < n; right++ {
		// Add nums[right] to max deque
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= nums[right] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// Add nums[right] to min deque
		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= nums[right] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)

		// Shrink window while cost > k
		for left <= right {
			curMin := nums[minQ[0]]
			curMax := nums[maxQ[0]]
			cost := (curMax - curMin) * (right - left + 1)
			if cost <= k {
				break
			}
			// Remove left from deques if at front
			if len(maxQ) > 0 && maxQ[0] == left {
				maxQ = maxQ[1:]
			}
			if len(minQ) > 0 && minQ[0] == left {
				minQ = minQ[1:]
			}
			left++
		}

		ans += right - left + 1
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountSubarraysWithCostLessThanOrEqualToK([]int{1, 3, 2}, 4)) // Expected: 5

	// Example 2
	fmt.Println(CountSubarraysWithCostLessThanOrEqualToK([]int{5, 5, 5, 5}, 0)) // Expected: 10

	// Example 3
	fmt.Println(CountSubarraysWithCostLessThanOrEqualToK([]int{1, 2, 3}, 0)) // Expected: 3
}
```
