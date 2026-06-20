# 2731 — Movement Of Robots

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MovementOfRobots(nums []int, s string, d int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2731: Movement of Robots
// https://leetcode.com/problems/movement-of-robots/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MovementOfRobots(nums []int, s string, d int) int {
	n := len(nums)
  // Alokasi slice
	pos := make([]int, n)
	for i, num := range nums {
		pos[i] = num
		if s[i] == 'R' {
			pos[i] += d
		} else {
			pos[i] -= d
		}
	}

  // Sort O(n log n)
	sort.Ints(pos)

	const mod = 1_000_000_007
	var sum int64
	var prefix int64
	for i, p := range pos {
		sum = (sum + int64(p)*int64(i) - prefix) % mod
		prefix += int64(p)
	}

	return int(sum)
}

func main() {
	fmt.Println(MovementOfRobots([]int{1, 0}, "RL", 2))
	fmt.Println(MovementOfRobots([]int{-2, 0, 2}, "RLL", 3))
}
```
