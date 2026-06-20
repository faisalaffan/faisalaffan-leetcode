# 1426 — Counting Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countElements(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1426: Counting Elements
// https://leetcode.com/problems/counting-elements/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func countElements(arr []int) int

import "fmt"

func main() {
	fmt.Println(CountingElements([]int{1, 2, 3}))       // 2
	fmt.Println(CountingElements([]int{1, 1, 3, 3, 5, 5, 7, 7})) // 0
	fmt.Println(CountingElements([]int{1, 1, 2, 2}))    // 2
}

// Time: O(n), Space: O(n)
func CountingElements(arr []int) int {
  // HashMap: O(1) lookup
	seen := make(map[int]bool, len(arr))
	for _, v := range arr {
		seen[v] = true
	}
	count := 0
	for _, v := range arr {
		if seen[v+1] {
			count++
		}
	}
	return count
}
```
