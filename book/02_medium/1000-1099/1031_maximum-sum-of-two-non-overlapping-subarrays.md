# 1031 — Maximum Sum Of Two Non Overlapping Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1031: Maximum Sum of Two Non-Overlapping Subarrays
// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/
// Difficulty: Medium
//
// Approach: DP with prefix sums. Consider both orders (L then M, M then L)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2)) // 20
	fmt.Println(maxSumTwoNoOverlap([]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2)) // 29
	fmt.Println(maxSumTwoNoOverlap([]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3)) // 31
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
  // Alokasi slice
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Try (firstLen before secondLen) and (secondLen before firstLen)
	result := 0
	// Case 1: firstLen comes first
	leftMax := 0
	for i := firstLen; i <= n-secondLen; i++ {
		leftSum := prefix[i] - prefix[i-firstLen]
		if leftSum > leftMax {
			leftMax = leftSum
		}
		rightSum := prefix[i+secondLen] - prefix[i]
		if leftMax+rightSum > result {
			result = leftMax + rightSum
		}
	}

	// Case 2: secondLen comes first
	leftMax = 0
	for i := secondLen; i <= n-firstLen; i++ {
		leftSum := prefix[i] - prefix[i-secondLen]
		if leftSum > leftMax {
			leftMax = leftSum
		}
		rightSum := prefix[i+firstLen] - prefix[i]
		if leftMax+rightSum > result {
			result = leftMax + rightSum
		}
	}

	return result
}
```
