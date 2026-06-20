# 3032 — Count Numbers With Unique Digits Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountNumbersWithUniqueDigitsIi(a int, b int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O((b-a) * log(b))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3032: Count Numbers With Unique Digits II
// https://leetcode.com/problems/count-numbers-with-unique-digits-ii/
// Difficulty: Easy [Paid]
//
// Note: This is a premium problem. In Go, we implement the equivalent logic.

import "fmt"

func main() {
	// LeetCode name: numberCount
	fmt.Println(CountNumbersWithUniqueDigitsIi(1, 20))  // 19
	fmt.Println(CountNumbersWithUniqueDigitsIi(9, 19))  // 10
}

// Time: O((b-a) * log(b)) | Space: O(1)
// LeetCode submission name: numberCount
func CountNumbersWithUniqueDigitsIi(a int, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		if hasUniqueDigits(i) {
			count++
		}
	}
	return count
}

func hasUniqueDigits(n int) bool {
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	for n > 0 {
		digit := n % 10
		if seen[digit] {
			return false
		}
		seen[digit] = true
		n /= 10
	}
	return true
}
```
