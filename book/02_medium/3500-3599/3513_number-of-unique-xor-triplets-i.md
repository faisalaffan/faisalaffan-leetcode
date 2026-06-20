# 3513 — Number Of Unique Xor Triplets I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfUniqueXorTripletsI(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3513: Number of Unique XOR Triplets I
// https://leetcode.com/problems/number-of-unique-xor-triplets-i/
// Difficulty: Medium
// Complexity: O(n^3) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", NumberOfUniqueXorTripletsI([]int{1, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", NumberOfUniqueXorTripletsI([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", NumberOfUniqueXorTripletsI([]int{5, 6, 7, 8}))
}

func NumberOfUniqueXorTripletsI(nums []int) int {
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				xor := nums[i] ^ nums[j] ^ nums[k]
				seen[xor] = true
			}
		}
	}
	return len(seen)
}
```
