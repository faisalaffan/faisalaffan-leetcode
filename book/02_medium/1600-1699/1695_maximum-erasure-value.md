# 1695 — Maximum Erasure Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumUniqueSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1695: Maximum Erasure Value
// https://leetcode.com/problems/maximum-erasure-value/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumUniqueSubarray(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	lastPos := make(map[int]int)
	maxSum := 0
	currentSum := 0
	left := 0

	for right, num := range nums {
		if pos, ok := lastPos[num]; ok && pos >= left {
			// Remove elements from left to pos
			for left <= pos {
				currentSum -= nums[left]
				left++
			}
		}
		currentSum += num
		lastPos[num] = right
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))   // Expected: 17
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) // Expected: 8
	fmt.Println(maximumUniqueSubarray([]int{1})) // Expected: 1
}
```
