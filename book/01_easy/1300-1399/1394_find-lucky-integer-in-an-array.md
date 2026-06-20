# 1394 — Find Lucky Integer In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func findLucky(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1394: Find Lucky Integer in an Array
// https://leetcode.com/problems/find-lucky-integer-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func findLucky(arr []int) int

import "fmt"

func main() {
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 3, 4}))       // 2
	fmt.Println(FindLuckyIntegerInAnArray([]int{1, 2, 2, 3, 3, 3})) // 3
	fmt.Println(FindLuckyIntegerInAnArray([]int{2, 2, 2, 3, 3}))    // -1
}

// Time: O(n), Space: O(n)
func FindLuckyIntegerInAnArray(arr []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int, len(arr))
	for _, v := range arr {
		freq[v]++
	}
	ans := -1
	for k, v := range freq {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}
```
