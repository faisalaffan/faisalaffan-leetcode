# 1031 — Maximum Sum Of Two Non Overlapping Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int
```

> **💡 Hint:** DP with prefix sums. Consider both orders (L then M, M then L)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
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
