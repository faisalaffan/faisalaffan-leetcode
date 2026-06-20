# 2177 — Find Three Consecutive Integers That Sum To A Given Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func sumOfThree(num int64) []int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2177: Find Three Consecutive Integers That Sum to a Given Number
// https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	x := num / 3
	return []int64{x - 1, x, x + 1}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sumOfThree(33))
	// Expected: [10, 11, 12]

	// Test case 2
	fmt.Println("Test 2:", sumOfThree(4))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", sumOfThree(0))
	// Expected: [-1, 0, 1]
}
```
