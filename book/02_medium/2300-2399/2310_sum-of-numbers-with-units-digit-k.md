# 2310 — Sum Of Numbers With Units Digit K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minimumNumbers(num int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2310: Sum of Numbers With Units Digit K
// https://leetcode.com/problems/sum-of-numbers-with-units-digit-k/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		if (i*k)%10 == num%10 && i*k <= num {
			return i
		}
	}
	return -1
}

func main() {
	// Test case 1
	fmt.Println(minimumNumbers(58, 9))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumNumbers(37, 2))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumNumbers(0, 7))
	// Expected: 0
}
```
