# 2620 — Counter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func counter(n int) func() int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2620: Counter
// https://leetcode.com/problems/counter/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript closure problem, adapted to Go. Returns a counter function.

import "fmt"

func main() {
	counter := counter(10)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter(n int) func() int {
	return func() int {
		n++
		return n - 1
	}
}
```
