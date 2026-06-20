# 1546 — Maximum Number Of Non Overlapping Subarrays With Sum Equals Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxNonOverlapping(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1546: Maximum Number of Non-Overlapping Subarrays With Sum Equals Target
// https://leetcode.com/problems/maximum-number-of-non-overlapping-subarrays-with-sum-equals-target/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxNonOverlapping([]int{1, 1, 1, 1, 1}, 2))
	fmt.Println(MaxNonOverlapping([]int{-1, 3, 5, 1, 4, 2, -9}, 6))
	fmt.Println(MaxNonOverlapping([]int{-2, 6, 6, 3, 5, 4, 1, 2, 8}, 10))
}

func MaxNonOverlapping(nums []int, target int) int {
	// Time: O(N), Space: O(N)
	// Greedy: use prefix sums map to find earliest non-overlapping subarray
	prefixSum := 0
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]int)
	seen[0] = -1 // prefix sum of empty array
	result := 0
	lastEnd := -1 // last used subarray end index

	for i, num := range nums {
		prefixSum += num
		if prevEnd, exists := seen[prefixSum-target]; exists && prevEnd >= lastEnd {
			result++
			lastEnd = i
		}
		seen[prefixSum] = i
	}

	return result
}
```
