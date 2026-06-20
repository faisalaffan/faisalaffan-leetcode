# 3738 — Longest Non Decreasing Subarray After Replacing At Most One Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestNonDecreasingSubarrayAfterReplacingAtMostOneElement(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3738: Longest Non-Decreasing Subarray After Replacing at Most One Element
// https://leetcode.com/problems/longest-non-decreasing-subarray-after-replacing-at-most-one-element/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestNonDecreasingSubarrayAfterReplacingAtMostOneElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

  // Alokasi slice integer
	left := make([]int, n)
  // Alokasi slice integer
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1

	for i := 1; i < n; i++ {
		if nums[i] >= nums[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		if i == 0 {
			if 1+right[i+1] > ans {
				ans = 1 + right[i+1]
			}
		} else if i == n-1 {
			if 1+left[i-1] > ans {
				ans = 1 + left[i-1]
			}
		} else if nums[i-1] <= nums[i+1] {
			cand := left[i-1] + 1 + right[i+1]
			if cand > ans {
				ans = cand
			}
		} else {
			cand := left[i-1]
			if right[i+1] > cand {
				cand = right[i+1]
			}
			cand++
			if cand > ans {
				ans = cand
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(longestNonDecreasingSubarrayAfterReplacingAtMostOneElement([]int{1, 2, 3, 1, 2}))
	fmt.Println(longestNonDecreasingSubarrayAfterReplacingAtMostOneElement([]int{5, 4, 3, 2, 1}))
	fmt.Println(longestNonDecreasingSubarrayAfterReplacingAtMostOneElement([]int{1, 2, 3, 4}))
}
```
