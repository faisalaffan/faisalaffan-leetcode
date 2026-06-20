# 2996 — Smallest Missing Integer Greater Than Sequential Prefix Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func SmallestMissingIntegerGreaterThanSequentialPrefixSum(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2996: Smallest Missing Integer Greater Than Sequential Prefix Sum
// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: missingInteger
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{1, 2, 3, 2, 5}))        // 6
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{3, 4, 5, 1, 12, 14, 13})) // 15
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{4, 5, 6, 7, 8, 9, 10, 11})) // 60
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: missingInteger
func SmallestMissingIntegerGreaterThanSequentialPrefixSum(nums []int) int {
	// Find longest sequential prefix (each element = prev + 1)
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	// Find smallest missing >= sum
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}
	for seen[sum] {
		sum++
	}
	return sum
}
```
