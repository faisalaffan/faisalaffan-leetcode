# 1512 — Number Of Good Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numIdenticalPairs(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1512: Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/
// Difficulty: Easy
//
// LeetCode submission: func numIdenticalPairs(nums []int) int

import "fmt"

func main() {
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	fmt.Println(NumberOfGoodPairs([]int{1, 1, 1, 1}))        // 6
	fmt.Println(NumberOfGoodPairs([]int{1, 2, 3}))           // 0
}

// Time: O(n), Space: O(n)
func NumberOfGoodPairs(nums []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v]
		freq[v]++
	}
	return count
}
```
