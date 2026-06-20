# 3005 — Count Elements With Maximum Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountElementsWithMaximumFrequency(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3005: Count Elements With Maximum Frequency
// https://leetcode.com/problems/count-elements-with-maximum-frequency/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maxFrequencyElements
	fmt.Println(CountElementsWithMaximumFrequency([]int{1, 2, 2, 3, 1, 4})) // 4
	fmt.Println(CountElementsWithMaximumFrequency([]int{1, 2, 3, 4, 5}))    // 5
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: maxFrequencyElements
func CountElementsWithMaximumFrequency(nums []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	maxFreq := 0
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}
	total := 0
	for _, f := range freq {
		if f == maxFreq {
			total += f
		}
	}
	return total
}
```
