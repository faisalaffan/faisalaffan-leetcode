# 2169 — Count Operations To Obtain Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountOperationsToObtainZero(num1 int, num2 int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log max(num1, num2)), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2169: Count Operations to Obtain Zero
// https://leetcode.com/problems/count-operations-to-obtain-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountOperationsToObtainZero(2, 3))  // 3
	fmt.Println(CountOperationsToObtainZero(10, 10)) // 1
}

// Time: O(log max(num1, num2)), Space: O(1)
func CountOperationsToObtainZero(num1 int, num2 int) int {
	ops := 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		ops++
	}
	return ops
}
```
