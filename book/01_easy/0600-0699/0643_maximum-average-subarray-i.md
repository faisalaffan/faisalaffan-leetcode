# 0643 — Maximum Average Subarray I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMaxAverage(nums []int, k int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #643: Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // 12.75
	fmt.Println(findMaxAverage([]int{5}, 1))                     // 5.0
	fmt.Println(findMaxAverage([]int{-1}, 1))                    // -1.0
}

// findMaxAverage finds a contiguous subarray of length k that has the maximum average value.
// Time: O(n). Space: O(1).
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
```
