# 1793 — Maximum Score Of A Good Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumScore(nums []int, k int) int
```

> **💡 Hint:** expand outward from k, maintaining the running minimum.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1793: Maximum Score of a Good Subarray
// https://leetcode.com/problems/maximum-score-of-a-good-subarray/
// Difficulty: Hard
//
// Given array nums and index k, find the maximum score of any "good" subarray
// that contains index k. Score = min(subarray) * length(subarray).
//
// Approach: expand outward from k, maintaining the running minimum.
// At each step, expand in the direction with the larger next value to keep
// the minimum as high as possible.

import "fmt"

func main() {
	// Example 1: nums = [1,4,3,7,4,5], k = 3 => 15 (subarray [4,3,7,4,5] min=3, len=5)
	fmt.Println(maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
	// Example 2: nums = [5,5,4,5,4,1,1,1], k = 0 => 20 (subarray [5,5,4,5,4] min=4, len=5)
	fmt.Println(maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0))
	// Single element
	fmt.Println(maximumScore([]int{5}, 0))
	// Two elements
	fmt.Println(maximumScore([]int{2, 1}, 1))
	// All same
	fmt.Println(maximumScore([]int{3, 3, 3, 3, 3}, 2))
	// Decreasing
	fmt.Println(maximumScore([]int{10, 9, 8, 7, 6}, 2))
	// Random
	fmt.Println(maximumScore([]int{6569, 9667, 3148, 7698, 1622, 6272, 4522, 2757, 5270, 9955}, 2))
}

func maximumScore(nums []int, k int) int {
	left, right := k, k
	minVal := nums[k]
	ans := nums[k]
	n := len(nums)

	for left > 0 || right < n-1 {
		if left == 0 {
			right++
		} else if right == n-1 {
			left--
		} else if nums[left-1] >= nums[right+1] {
			left--
		} else {
			right++
		}
		if nums[left] < minVal {
			minVal = nums[left]
		}
		if nums[right] < minVal {
			minVal = nums[right]
		}
		score := minVal * (right - left + 1)
		if score > ans {
			ans = score
		}
	}
	return ans
}
```
