# 0277 — Find The Celebrity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func knows(a, b int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #277: Find the Celebrity
// https://leetcode.com/problems/find-the-celebrity/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

var knowsMatrix [][]int

func knows(a, b int) bool {
	return knowsMatrix[a][b] == 1
}

func findCelebrity(n int) int {
	candidate := 0

	for i := 1; i < n; i++ {
		if knows(candidate, i) {
			candidate = i
		}
	}

	for i := 0; i < n; i++ {
		if i == candidate {
			continue
		}
		if knows(candidate, i) || !knows(i, candidate) {
			return -1
		}
	}

	return candidate
}

func main() {
	knowsMatrix = [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(findCelebrity(3))

	knowsMatrix = [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(findCelebrity(3))

	knowsMatrix = [][]int{{1, 1}, {0, 1}}
	fmt.Println(findCelebrity(2))
}
```
