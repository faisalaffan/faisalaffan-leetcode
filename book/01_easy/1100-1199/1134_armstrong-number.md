# 1134 — Armstrong Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func isArmstrong(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1134: Armstrong Number
// https://leetcode.com/problems/armstrong-number/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isArmstrong(153))  // true
	fmt.Println(isArmstrong(123))  // false
	fmt.Println(isArmstrong(1))    // true
}

// LeetCode submission: isArmstrong
func isArmstrong(n int) bool {
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}
	sum := 0
	for x := n; x > 0; x /= 10 {
		d := x % 10
		p := 1
		for i := 0; i < digits; i++ {
			p *= d
		}
		sum += p
	}
	return sum == n
}
```
