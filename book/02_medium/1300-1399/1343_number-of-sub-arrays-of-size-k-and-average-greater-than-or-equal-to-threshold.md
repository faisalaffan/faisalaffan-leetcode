# 1343 — Number Of Sub Arrays Of Size K And Average Greater Than Or Equal To Threshold

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numOfSubarrays(arr []int, k int, threshold int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n) where n = len(arr)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1343: Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4)) // 3

	// Test case 2
	fmt.Println(numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5)) // 6

	// Test case 3
	fmt.Println(numOfSubarrays([]int{1, 1, 1, 1, 1}, 1, 0)) // 5
}

// Time: O(n) where n = len(arr)
// Space: O(1)
func numOfSubarrays(arr []int, k int, threshold int) int {
	targetSum := k * threshold
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	count := 0
	if windowSum >= targetSum {
		count++
	}

	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum >= targetSum {
			count++
		}
	}

	return count
}
```
