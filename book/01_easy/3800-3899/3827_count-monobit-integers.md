# 3827 — Count Monobit Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountMonobitIntegers(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3827: Count Monobit Integers
// https://leetcode.com/problems/count-monobit-integers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountMonobitIntegers(1))
	fmt.Println(CountMonobitIntegers(4))
	fmt.Println(CountMonobitIntegers(0))
}

// Time: O(log n)
// Space: O(1)
func CountMonobitIntegers(n int) int {
	ans := 1 // 0 is monobit (all zeros)
	x := 1   // 2^1 - 1 = 1 (all ones)
	for x <= n {
		ans++
		x = x*2 + 1
	}
	return ans
}
```
