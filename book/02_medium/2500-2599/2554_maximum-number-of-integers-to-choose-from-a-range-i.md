# 2554 — Maximum Number Of Integers To Choose From A Range I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxCount(banned []int, n int, maxSum int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2554: Maximum Number of Integers to Choose From a Range I
// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func maxCount(banned []int, n int, maxSum int) int {
  // HashMap: O(1) lookup
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	count := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if bannedSet[i] {
			continue
		}
		if sum+i > maxSum {
			break
		}
		sum += i
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxCount([]int{1, 6, 5}, 5, 6))
	// Expected: 2 (choose 2,3)

	// Test case 2
	fmt.Println("Test 2:", maxCount([]int{1, 2, 3, 4, 5, 6, 7}, 8, 1))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxCount([]int{11}, 7, 50))
	// Expected: 7 (choose 1..7)
}
```
