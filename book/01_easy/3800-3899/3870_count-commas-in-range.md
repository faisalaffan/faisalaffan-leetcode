# 3870 — Count Commas In Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountCommasInRange(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3870: Count Commas in Range
// https://leetcode.com/problems/count-commas-in-range/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountCommasInRange(1002))
	fmt.Println(CountCommasInRange(998))
	fmt.Println(CountCommasInRange(1500000))
}

// Time: O(log n)
// Space: O(1)
func CountCommasInRange(n int) int {
	total := 0
	threshold := 1000
	for n >= threshold {
		total += n - threshold + 1
		threshold *= 1000
	}
	return total
}
```
