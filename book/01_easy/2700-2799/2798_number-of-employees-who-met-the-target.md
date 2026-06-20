# 2798 — Number Of Employees Who Met The Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfEmployeesWhoMetTheTarget(hours []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2798: Number of Employees Who Met the Target
// https://leetcode.com/problems/number-of-employees-who-met-the-target/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{5, 1, 4, 2, 2}, 6))
}

func NumberOfEmployeesWhoMetTheTarget(hours []int, target int) int {
	count := 0
	for _, h := range hours {
		if h >= target {
			count++
		}
	}
	return count
}
```
