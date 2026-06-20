# 1780 — Check If Number Is A Sum Of Powers Of Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func checkPowersOfThree(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log_3 n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1780: Check if Number is a Sum of Powers of Three
// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
// Difficulty: Medium
// Time: O(log_3 n), Space: O(1)

import "fmt"

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	fmt.Println(checkPowersOfThree(12)) // Expected: true (3^2 + 3^1 = 9 + 3)
	fmt.Println(checkPowersOfThree(91)) // Expected: true (3^4 + 3^2 + 3^0 = 81 + 9 + 1)
	fmt.Println(checkPowersOfThree(21)) // Expected: false (21 = 2*9 + 3, cannot use 2)
}
```
