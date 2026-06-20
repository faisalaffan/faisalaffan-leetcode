# 1477 — Find Two Non Overlapping Sub Arrays Each With Target Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minSumOfLengths(arr []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(n) where n = len(arr)  
**Kompleksitas Ruang:** O(n) for prefix minimum array

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1477: Find Two Non-overlapping Sub-arrays Each With Target Sum
// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3)) // 2

	// Test case 2
	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7)) // 2

	// Test case 3
	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6)) // -1

	// Test case 4
	fmt.Println(minSumOfLengths([]int{1, 1, 1, 2, 1, 1}, 3)) // 3
}

// Time: O(n) where n = len(arr)
// Space: O(n) for prefix minimum array
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	// left[i] = minimum length of subarray with sum = target ending at or before i
  // Alokasi slice integer
	left := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range left {
		left[i] = math.MaxInt32
	}

	prefixSum := 0
  // Membuat map (HashMap) — pencarian O(1)
	sumMap := make(map[int]int)
	sumMap[0] = -1
	bestLeft := math.MaxInt32

	for i := 0; i < n; i++ {
		prefixSum += arr[i]
		if j, ok := sumMap[prefixSum-target]; ok {
			length := i - j
			if length < bestLeft {
				bestLeft = length
			}
			if i > 0 && left[i-1] < bestLeft {
				bestLeft = left[i-1]
			}
		}
		left[i] = bestLeft
		sumMap[prefixSum] = i
	}

	// right[i] = minimum length of subarray with sum = target starting at or after i
  // Alokasi slice integer
	right := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range right {
		right[i] = math.MaxInt32
	}

	suffixSum := 0
	sumMap = make(map[int]int)
	sumMap[0] = n
	bestRight := math.MaxInt32

	for i := n - 1; i >= 0; i-- {
		suffixSum += arr[i]
		if j, ok := sumMap[suffixSum-target]; ok {
			length := j - i
			if length < bestRight {
				bestRight = length
			}
			if i < n-1 && right[i+1] < bestRight {
				bestRight = right[i+1]
			}
		}
		right[i] = bestRight
		sumMap[suffixSum] = i
	}

	result := math.MaxInt32
	for i := 0; i < n-1; i++ {
		if left[i] != math.MaxInt32 && right[i+1] != math.MaxInt32 {
			total := left[i] + right[i+1]
			if total < result {
				result = total
			}
		}
	}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}
```
