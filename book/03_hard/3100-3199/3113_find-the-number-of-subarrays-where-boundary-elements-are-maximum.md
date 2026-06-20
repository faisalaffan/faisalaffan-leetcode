# 3113 — Find The Number Of Subarrays Where Boundary Elements Are Maximum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubarrays(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3113: Find the Number of Subarrays Where Boundary Elements Are Maximum
// https://leetcode.com/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/
// Difficulty: Hard
// Time: O(n) | Space: O(n)
//
// A subarray is valid if its first and last elements are the maximum in the subarray.
// Use a monotonic decreasing stack. For each position i, count[i] = number of valid
// subarrays ending at i where nums[i] is the maximum. This extends from the previous
// same-value element's chain.

import (
	"fmt"
)

func numberOfSubarrays(nums []int) int64 {
	n := len(nums)
  // Alokasi slice integer
	stack := make([]int, 0)
  // Alokasi slice integer
	count := make([]int64, n)
	var result int64 = 0

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && nums[stack[len(stack)-1]] == nums[i] {
			count[i] = count[stack[len(stack)-1]] + 1
		} else {
			count[i] = 1
		}
		stack = append(stack, i)
		result += count[i]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubarrays([]int{1, 4, 3, 3, 2}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", numberOfSubarrays([]int{3, 3, 3}))
	// Expected: 6

	// Test case 3: single element
	fmt.Println("Test 3:", numberOfSubarrays([]int{1}))
	// Expected: 1

	// Test case 4: strictly increasing
	fmt.Println("Test 4:", numberOfSubarrays([]int{1, 2, 3, 4}))
	// Expected: 4

	// Test case 5: strictly decreasing
	fmt.Println("Test 5:", numberOfSubarrays([]int{4, 3, 2, 1}))
	// Expected: 4
}
```
