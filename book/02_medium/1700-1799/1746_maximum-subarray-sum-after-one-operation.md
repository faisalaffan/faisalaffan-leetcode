# 1746 — Maximum Subarray Sum After One Operation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSumAfterOperation(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1746: Maximum Subarray Sum After One Operation
// https://leetcode.com/problems/maximum-subarray-sum-after-one-operation/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func maxSumAfterOperation(nums []int) int {
	// dp0: max subarray sum without using operation
	// dp1: max subarray sum with exactly one operation used
	dp0 := 0
	dp1 := 0
	maxSum := nums[0] * nums[0]

	for _, v := range nums {
		// Either start new or extend
		newDp0 := max(v, dp0+v)
		newDp1 := max(v*v, dp0+v*v, dp1+v)

		dp0 = newDp0
		dp1 = newDp1
		maxSum = max(maxSum, dp1)
	}
	return maxSum
}

func max(nums ...int) int {
	result := nums[0]
	for _, v := range nums[1:] {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Println(maxSumAfterOperation([]int{2, -1, -4, -3})) // Expected: 17
	fmt.Println(maxSumAfterOperation([]int{1, -2, 3, 4}))   // Expected: 19 ([3,4] with 4^2=16: 3+16=19)

	fmt.Println(maxSumAfterOperation([]int{-1, -1, -1})) // Expected: 1 (replace -1 with 1)
}
```
