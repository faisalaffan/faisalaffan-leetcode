# 3483 — Unique 3 Digit Even Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func UniqueThreeDigitEvenNumbers(digits []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^3). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3483: Unique 3-Digit Even Numbers
// https://leetcode.com/problems/unique-3-digit-even-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(UniqueThreeDigitEvenNumbers([]int{1, 2, 3, 4}))
	fmt.Println(UniqueThreeDigitEvenNumbers([]int{0, 2, 2}))
}

// UniqueThreeDigitEvenNumbers counts unique 3-digit even numbers that can be formed from digits (no leading zero).
// Time: O(n^3). Space: O(n).
func UniqueThreeDigitEvenNumbers(digits []int) int {
  // HashMap: O(1) lookup
	used := make(map[int]bool)
	n := len(digits)
	for i := 0; i < n; i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if num%2 == 0 {
					used[num] = true
				}
			}
		}
	}
	return len(used)
}
```
