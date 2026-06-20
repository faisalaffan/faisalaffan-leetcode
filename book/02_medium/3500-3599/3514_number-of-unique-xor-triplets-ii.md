# 3514 — Number Of Unique Xor Triplets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfUniqueXorTripletsIi(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3514: Number of Unique XOR Triplets II
// https://leetcode.com/problems/number-of-unique-xor-triplets-ii/
// Difficulty: Medium
// Complexity: O(n^2) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", NumberOfUniqueXorTripletsIi([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", NumberOfUniqueXorTripletsIi([]int{1, 1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", NumberOfUniqueXorTripletsIi([]int{0, 0, 0}))
}

func NumberOfUniqueXorTripletsIi(nums []int) int {
	// XOR of any pair, then XOR with third element
  // HashMap: O(1) lookup
	pairXors := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairXors[nums[i]^nums[j]] = true
		}
	}

  // HashMap: O(1) lookup
	unique := make(map[int]bool)
	for px := range pairXors {
		for _, v := range nums {
			unique[px^v] = true
		}
	}
	return len(unique)
}
```
