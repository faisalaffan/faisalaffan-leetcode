# 1060 — Missing Element In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func missingElement(nums []int, k int) int
```

> **💡 Hint:** Binary search. Number of missing elements up to index i = nums[i] - nums[0] - i.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1060: Missing Element in Sorted Array
// https://leetcode.com/problems/missing-element-in-sorted-array/
// Difficulty: Medium
//
// Approach: Binary search. Number of missing elements up to index i = nums[i] - nums[0] - i.
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 1))  // 5
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 3))  // 8
	fmt.Println(missingElement([]int{1, 2, 4}, 3))      // 6
}

func missingElement(nums []int, k int) int {
	n := len(nums)

	missingCount := func(idx int) int {
		return nums[idx] - nums[0] - idx
	}

	if missingCount(n-1) < k {
		return nums[n-1] + (k - missingCount(n-1))
	}

	left, right := 0, n-1
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if missingCount(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left-1] + (k - missingCount(left-1))
}
```
