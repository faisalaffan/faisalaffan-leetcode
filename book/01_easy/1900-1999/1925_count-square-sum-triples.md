# 1925 — Count Square Sum Triples

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountSquareSumTriples(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1925: Count Square Sum Triples
// https://leetcode.com/problems/count-square-sum-triples/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSquareSumTriples(5))  // 2
	fmt.Println(CountSquareSumTriples(10)) // 4
}

// Time: O(n^2), Space: O(1)
func CountSquareSumTriples(n int) int {
	count := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			c2 := a*a + b*b
			c := 1
			for c*c < c2 {
				c++
			}
			if c*c == c2 && c <= n {
				count++
			}
		}
	}
	return count
}
```
