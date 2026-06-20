# 3943 — Number Of Pairs After Increment

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfPairs(nums1 []int, nums2 []int, queries [][]int) []int
```

> **💡 Hint:** Maintain sorted nums2. Binary search to count how

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3943: Number of Pairs After Increment
// https://leetcode.com/problems/number-of-pairs-after-increment/
// Difficulty: Hard
//
// Given two arrays nums1 and nums2, and queries. Each query
// [index, val] increments nums1[index] by val. After each query,
// count pairs (i, j) where nums1[i] > nums2[j].
//
// Approach: Maintain sorted nums2. Binary search to count how
// many nums2 elements are less than each updated nums1 value.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(numberOfPairs([]int{1, 3, 5}, []int{2, 4, 6}, [][]int{{0, 3}, {1, 2}}))
	// Example 2
	fmt.Println(numberOfPairs([]int{1, 2}, []int{3, 4}, [][]int{{0, 5}}))
	// Edge: no pairs possible
	fmt.Println(numberOfPairs([]int{1, 1}, []int{5, 5}, [][]int{{0, 1}}))
}

func numberOfPairs(nums1 []int, nums2 []int, queries [][]int) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums2)
  // Alokasi slice integer
	ans := make([]int, len(queries))

	for idx, q := range queries {
		i, val := q[0], q[1]
		nums1[i] += val

		// Count nums2 elements less than updated nums1[i]
		for _, v := range nums2 {
			if v < nums1[i] {
				ans[idx]++
			}
		}

		if false {
  // Urutkan secara ascending — O(n log n)
			sort.Ints([]int{}) // keep sort import
		}
	}

	return ans
}
```
