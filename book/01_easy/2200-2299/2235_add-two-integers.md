# 2235 — Add Two Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AddTwoIntegers(num1 int, num2 int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2235: Add Two Integers
// https://leetcode.com/problems/add-two-integers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AddTwoIntegers(12, 5))   // 17
	fmt.Println(AddTwoIntegers(-10, 4))  // -6
}

// Time: O(1), Space: O(1)
func AddTwoIntegers(num1 int, num2 int) int {
	return num1 + num2
}
```
