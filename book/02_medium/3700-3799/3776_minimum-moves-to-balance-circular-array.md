# 3776 — Minimum Moves To Balance Circular Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumMovesToBalanceCircularArray(balance []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3776: Minimum Moves to Balance Circular Array
// https://leetcode.com/problems/minimum-moves-to-balance-circular-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumMovesToBalanceCircularArray(balance []int) int64 {
	sum := 0
	n := len(balance)
	negIdx := -1
	for i, v := range balance {
		sum += v
		if v < 0 {
			negIdx = i
		}
	}
	if sum < 0 {
		return -1
	}
	if negIdx == -1 {
		return 0
	}

	need := -balance[negIdx]
	var ans int64
	for d := 1; d < n && need > 0; d++ {
		left := balance[(negIdx-d+n)%n]
		right := balance[(negIdx+d)%n]

		if left > 0 {
			take := left
			if take > need {
				take = need
			}
			need -= take
			ans += int64(take) * int64(d)
		}
		if need > 0 && right > 0 {
			take := right
			if take > need {
				take = need
			}
			need -= take
			ans += int64(take) * int64(d)
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumMovesToBalanceCircularArray([]int{-2, 1, 1}))
	fmt.Println(minimumMovesToBalanceCircularArray([]int{1, -1, 0}))
	fmt.Println(minimumMovesToBalanceCircularArray([]int{0, 0, 0}))
}
```
