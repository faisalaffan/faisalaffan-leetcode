# 2829 — Determine The Minimum Sum Of A K Avoiding Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DetermineTheMinimumSumOfAKAvoidingArray(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2829: Determine the Minimum Sum of a k-avoiding Array
// https://leetcode.com/problems/determine-the-minimum-sum-of-a-k-avoiding-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func DetermineTheMinimumSumOfAKAvoidingArray(n int, k int) int {
  // HashMap: O(1) lookup
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[k-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(DetermineTheMinimumSumOfAKAvoidingArray(5, 4))
	fmt.Println(DetermineTheMinimumSumOfAKAvoidingArray(3, 5))
}
```
