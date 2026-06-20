# 3158 — Find The Xor Of Numbers Which Appear Twice

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheXorOfNumbersWhichAppearTwice(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3158: Find the XOR of Numbers Which Appear Twice
// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: duplicateNumbersXOR
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 2, 1})) // 3
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 3}))    // 0
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: duplicateNumbersXOR
func FindTheXorOfNumbersWhichAppearTwice(nums []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	xor := 0
	for v, f := range freq {
		if f == 2 {
			xor ^= v
		}
	}
	return xor
}
```
