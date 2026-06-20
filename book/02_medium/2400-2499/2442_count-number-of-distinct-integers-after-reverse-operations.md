# 2442 — Count Number Of Distinct Integers After Reverse Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countDistinctIntegers(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * log(max))  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2442: Count Number of Distinct Integers After Reverse Operations
// https://leetcode.com/problems/count-number-of-distinct-integers-after-reverse-operations/
// Difficulty: Medium
// Time: O(n * log(max)) | Space: O(n)
// For each num, add num and its reverse to a set.

import "fmt"

func main() {
	fmt.Println(countDistinctIntegers([]int{1, 13, 10, 12, 31})) // 6
	fmt.Println(countDistinctIntegers([]int{2, 2, 2}))           // 1
}

func countDistinctIntegers(nums []int) int {
  // HashMap: O(1) lookup
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
		set[reverse(v)] = true
	}
	return len(set)
}

func reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}
```
