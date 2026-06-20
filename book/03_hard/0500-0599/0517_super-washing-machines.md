# 0517 — Super Washing Machines

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findMinMoves(machines []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #517: Super Washing Machines
// https://leetcode.com/problems/super-washing-machines/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(findMinMoves([]int{1, 0, 5})) // Expected: 3
}

func findMinMoves(machines []int) int {
	n := len(machines)
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%n != 0 {
		return -1
	}
	target := sum / n

	ans := 0
	balance := 0
	for _, v := range machines {
		balance += v - target
		if balance > ans {
			ans = balance
		}
		if balance < -ans {
			ans = -balance
		}
		// A machine may need to receive from both sides simultaneously
		if v-target > ans {
			ans = v - target
		}
	}
	return ans
}
```
