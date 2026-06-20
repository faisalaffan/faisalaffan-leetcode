# 0053 — Maximum Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSubArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #53: Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
// Difficulty: Medium

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currSum+nums[i] > nums[i] {
			currSum = currSum + nums[i]
		} else {
			currSum = nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}

func main() {
	// Test case 1
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	// Test case 2
	fmt.Println(maxSubArray([]int{1})) // 1

	// Test case 3
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8})) // 23
}

// Time: O(n) | Space: O(1)
```
