# 3566 — Partition Array Into Two Equal Product Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PartitionArrayIntoTwoEqualProductSubsets(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3566: Partition Array into Two Equal Product Subsets
// https://leetcode.com/problems/partition-array-into-two-equal-product-subsets/
// Difficulty: Medium
// Complexity: O(2^n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionArrayIntoTwoEqualProductSubsets([]int{1, 2, 3, 6}))
	// Test case 2
	fmt.Println("Test 2:", PartitionArrayIntoTwoEqualProductSubsets([]int{2, 3, 4, 6}))
	// Test case 3
	fmt.Println("Test 3:", PartitionArrayIntoTwoEqualProductSubsets([]int{1, 1, 1}))
}

func PartitionArrayIntoTwoEqualProductSubsets(nums []int) bool {
	totalProduct := 1
	for _, v := range nums {
		totalProduct *= v
	}
	// We need to find subset whose product = sqrt(totalProduct)
	// sqrt(totalProduct) must be integer
	target := 1
	// Find largest square factor
	for i := 2; i*i <= totalProduct; i++ {
		for totalProduct%(i*i) == 0 {
			target *= i
			totalProduct /= (i * i)
		}
	}
	// Use subset sum approach but with multiplication
	return canPartition(nums, target, 0, 1)
}

func canPartition(nums []int, target int, idx int, product int) bool {
	if product == target {
		return true
	}
	if product > target || idx >= len(nums) {
		return false
	}
	return canPartition(nums, target, idx+1, product) ||
		canPartition(nums, target, idx+1, product*nums[idx])
}
```
