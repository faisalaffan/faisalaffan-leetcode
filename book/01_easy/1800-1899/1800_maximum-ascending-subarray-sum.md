# 1800 — Maximum Ascending Subarray Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxAscendingSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1800: Maximum Ascending Subarray Sum
// https://leetcode.com/problems/maximum-ascending-subarray-sum/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaxAscendingSum(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(MaxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	fmt.Println(MaxAscendingSum([]int{10, 20, 30, 40, 50}))
	fmt.Println(MaxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}
```
