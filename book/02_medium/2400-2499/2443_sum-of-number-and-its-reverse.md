# 2443 — Sum Of Number And Its Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func sumOfNumberAndReverse(num int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(num)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2443: Sum of Number and Its Reverse
// https://leetcode.com/problems/sum-of-number-and-its-reverse/
// Difficulty: Medium
// Time: O(num) | Space: O(1)
// Iterate from 0 to num, check if i + reverse(i) == num.

import "fmt"

func main() {
	fmt.Println(sumOfNumberAndReverse(443)) // true (241+142=443)
	fmt.Println(sumOfNumberAndReverse(63))  // false
	fmt.Println(sumOfNumberAndReverse(181)) // true (90+9=99? wait) (140+41=181)
}

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}

func reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}
```
