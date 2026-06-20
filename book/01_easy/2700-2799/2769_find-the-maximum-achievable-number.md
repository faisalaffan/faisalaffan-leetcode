# 2769 — Find The Maximum Achievable Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheMaximumAchievableNumber(num int, t int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2769: Find the Maximum Achievable Number
// https://leetcode.com/problems/find-the-maximum-achievable-number/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumAchievableNumber(4, 1))
	fmt.Println(FindTheMaximumAchievableNumber(3, 2))
}

func FindTheMaximumAchievableNumber(num int, t int) int {
	return num + 2*t
}
```
