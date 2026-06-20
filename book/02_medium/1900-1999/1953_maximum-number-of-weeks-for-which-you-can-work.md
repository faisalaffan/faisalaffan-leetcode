# 1953 — Maximum Number Of Weeks For Which You Can Work

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfWeeks(milestones []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1953: Maximum Number of Weeks for Which You Can Work
// https://leetcode.com/problems/maximum-number-of-weeks-for-which-you-can-work/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfWeeks([]int{1, 2, 3}))
	fmt.Println(NumberOfWeeks([]int{5, 2, 1}))
	fmt.Println(NumberOfWeeks([]int{9, 3, 6, 8, 2, 1}))
}

// Time: O(n), Space: O(1)
func NumberOfWeeks(milestones []int) int64 {
	var sum int64 = 0
	maxVal := 0
	for _, m := range milestones {
		sum += int64(m)
		if m > maxVal {
			maxVal = m
		}
	}

	rest := sum - int64(maxVal)
	if int64(maxVal) > rest+1 {
		return 2*rest + 1
	}
	return sum
}
```
