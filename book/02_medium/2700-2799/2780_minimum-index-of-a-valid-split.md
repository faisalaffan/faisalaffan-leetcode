# 2780 — Minimum Index Of A Valid Split

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumIndexOfAValidSplit(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2780: Minimum Index of a Valid Split
// https://leetcode.com/problems/minimum-index-of-a-valid-split/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumIndexOfAValidSplit(nums []int) int {
	n := len(nums)

	// Find majority element using Boyer-Moore
	majority := nums[0]
	count := 0
	for _, num := range nums {
		if count == 0 {
			majority = num
			count = 1
		} else if num == majority {
			count++
		} else {
			count--
		}
	}

	// Count total occurrences of majority
	totalCount := 0
	for _, num := range nums {
		if num == majority {
			totalCount++
		}
	}

	// Find minimum split index
	leftCount := 0
	for i := 0; i < n-1; i++ {
		if nums[i] == majority {
			leftCount++
		}
		rightCount := totalCount - leftCount
		leftLen := i + 1
		rightLen := n - i - 1
		if leftCount*2 > leftLen && rightCount*2 > rightLen {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(MinimumIndexOfAValidSplit([]int{1, 2, 2, 2}))
	fmt.Println(MinimumIndexOfAValidSplit([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
}
```
