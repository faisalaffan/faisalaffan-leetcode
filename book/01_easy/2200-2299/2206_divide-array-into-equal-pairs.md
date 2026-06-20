# 2206 — Divide Array Into Equal Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DivideArrayIntoEqualPairs(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2206: Divide Array Into Equal Pairs
// https://leetcode.com/problems/divide-array-into-equal-pairs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideArrayIntoEqualPairs([]int{3, 2, 3, 2, 2, 2})) // true
	fmt.Println(DivideArrayIntoEqualPairs([]int{1, 2, 3, 4}))       // false
}

// Time: O(n), Space: O(n)
func DivideArrayIntoEqualPairs(nums []int) bool {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for _, c := range freq {
		if c%2 != 0 {
			return false
		}
	}
	return true
}
```
