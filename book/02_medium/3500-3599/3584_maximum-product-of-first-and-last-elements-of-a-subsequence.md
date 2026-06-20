# 3584 — Maximum Product Of First And Last Elements Of A Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumProductOfFirstAndLastElementsOfASubsequence(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3584: Maximum Product of First and Last Elements of a Subsequence
// https://leetcode.com/problems/maximum-product-of-first-and-last-elements-of-a-subsequence/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumProductOfFirstAndLastElementsOfASubsequence([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MaximumProductOfFirstAndLastElementsOfASubsequence([]int{-1, -2, -3}))
	// Test case 3
	fmt.Println("Test 3:", MaximumProductOfFirstAndLastElementsOfASubsequence([]int{5}))
}

func MaximumProductOfFirstAndLastElementsOfASubsequence(nums []int) int {
	if len(nums) < 2 {
		if len(nums) == 1 {
			return nums[0] * nums[0]
		}
		return 0
	}
	// The subsequence can be any; we can pick first and last element of any subsequence
	// For any subsequence, the first element is at index i and last at index j (i <= j)
	// We want max product of nums[i] * nums[j]
	maxProd := nums[0] * nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			prod := nums[i] * nums[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}
```
