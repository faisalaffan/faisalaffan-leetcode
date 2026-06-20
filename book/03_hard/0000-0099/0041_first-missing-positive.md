# 0041 — First Missing Positive

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func firstMissingPositive(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #41: First Missing Positive
// https://leetcode.com/problems/first-missing-positive/
// Difficulty: Hard

import "fmt"

// firstMissingPositive finds the smallest missing positive integer.
// Uses cyclic sort (O(1) extra space by modifying the input array).
//
// Complexity: O(n) time, O(1) space
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Cyclic sort: place each number at its correct index (num-1)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// Find the first index where the value is not (index+1)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test 1: [1,2,0] ->", firstMissingPositive([]int{1, 2, 0}))          // 3
	fmt.Println("Test 2: [3,4,-1,1] ->", firstMissingPositive([]int{3, 4, -1, 1}))   // 2
	fmt.Println("Test 3: [7,8,9,11,12] ->", firstMissingPositive([]int{7, 8, 9, 11, 12})) // 1
	fmt.Println("Test 4: [1,2,3] ->", firstMissingPositive([]int{1, 2, 3}))          // 4
	fmt.Println("Test 5: [] ->", firstMissingPositive([]int{}))                       // 1
}
```
