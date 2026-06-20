# 0754 — Reach A Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func reachNumber(target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(sqrt(target))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #754: Reach a Number
// https://leetcode.com/problems/reach-a-number/
// Difficulty: Medium
// Time: O(sqrt(target))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(reachNumber(3))
	fmt.Println(reachNumber(2))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	sum := 0
	steps := 0

	for sum < target || (sum-target)%2 != 0 {
		steps++
		sum += steps
	}

	return steps
}
```
