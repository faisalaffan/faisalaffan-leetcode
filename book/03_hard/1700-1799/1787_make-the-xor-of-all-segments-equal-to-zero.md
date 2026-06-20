# 1787 — Make The Xor Of All Segments Equal To Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minChanges(nums []int, k int) int
```

> **💡 Hint:** DP with grouping by index mod k.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1787: Make the XOR of All Segments Equal to Zero
// https://leetcode.com/problems/make-the-xor-of-all-segments-equal-to-zero/
// Difficulty: Hard
//
// Approach: DP with grouping by index mod k.
// For all windows of length k, their XOR must be equal. This means:
// nums[i] XOR nums[i+1] XOR ... XOR nums[i+k-1] == 0 for all i.
// This implies nums[i] == nums[i + k] for all i.
// So the array must be periodic with period k.
// Additionally, the XOR of the first k elements must be 0.
//
// Group nums by index mod k. For each group (0..k-1), we need to pick a
// value for all positions in that group. The XOR of the k chosen values
// must be 0.
//
// For each group g, count the frequency of each value.
// For each group, consider changing all elements to some value v.
// Cost for group g choosing v = size(g) - freq[g][v].
//
// DP over groups: dp[i][xor] = min cost for first i groups achieving XOR = xor.
// Result = dp[k][0].

import (
	"fmt"
	"math"
)

func minChanges(nums []int, k int) int {
	n := len(nums)

	// Group values by index mod k
  // Alokasi slice integer
	groups := make([]map[int]int, k)
	for i := 0; i < k; i++ {
		groups[i] = make(map[int]int)
	}
	for i, v := range nums {
		groups[i%k][v]++
	}

  // Alokasi slice integer
	groupSizes := make([]int, k)
	for g := 0; g < k; g++ {
		groupSizes[g] = len(groups[g]) // actually number of unique values is not the size; size = n/k rounded properly
	}
	// Actual group sizes (how many positions per group)
	for i := 0; i < n; i++ {
		_ = i % k // we'll compute directly
	}

	// Size of each group (number of positions)
  // Alokasi slice integer
	size := make([]int, k)
	for i := 0; i < n; i++ {
		size[i%k]++
	}

	const maxXor = 1024 // nums[i] < 1024 based on constraints (2^10)
	INF := math.MaxInt32

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, k+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, maxXor)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for g := 0; g < k; g++ {
		// Option 1: For each position in group, we can change to any value.
		// Compute globalBest = min over previous xor of dp[g][prevXor] + size[g]
		// (since we can change all elements in this group to some value that gives
		// best possible cost)
		globalBest := INF
		for x := 0; x < maxXor; x++ {
			if dp[g][x] < globalBest {
				globalBest = dp[g][x]
			}
		}

		// For each possible next xor state
		for nextXor := 0; nextXor < maxXor; nextXor++ {
			// Best via "change all in group to some value v that gives us global best"
			// Changing all in group costs size[g], and we can pick any value v.
			// This handles the case where we change to a value not in the group.
			best := globalBest + size[g]

			// Option 2: For each value present in this group, we can keep it
			for val, freq := range groups[g] {
				prevXor := nextXor ^ val
				if dp[g][prevXor]+size[g]-freq < best {
					best = dp[g][prevXor] + size[g] - freq
				}
			}

			dp[g+1][nextXor] = best
		}
	}

	return dp[k][0]
}

func main() {
	// Example test case
	fmt.Println("nums=[1,2,3,4,5,6],k=3 →", minChanges([]int{1, 2, 3, 4, 5, 6}, 3)) // Expected: 3

	// Additional tests
	fmt.Println("nums=[3,4,5,2,1,7,3,4,7],k=3 →", minChanges([]int{3, 4, 5, 2, 1, 7, 3, 4, 7}, 3))
	fmt.Println("nums=[1,2,3],k=1 →", minChanges([]int{1, 2, 3}, 1)) // Need XOR of whole array = 0, change 2

	// Edge case: already periodic
	fmt.Println("nums=[1,2,1,2],k=2 →", minChanges([]int{1, 2, 1, 2}, 2)) // Already: nums[0]=nums[2]=1, nums[1]=nums[3]=2, XOR=1^2=3≠0 so need 1 change
}
```
