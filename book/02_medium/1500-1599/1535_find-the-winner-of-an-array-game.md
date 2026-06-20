# 1535 — Find The Winner Of An Array Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func GetWinner(arr []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1535: Find the Winner of an Array Game
// https://leetcode.com/problems/find-the-winner-of-an-array-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
	fmt.Println(GetWinner([]int{3, 2, 1}, 10))
	fmt.Println(GetWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000))
}

func GetWinner(arr []int, k int) int {
	// Time: O(N), Space: O(1)
	if k == 0 {
		return 0
	}

	// If k >= n, the maximum element wins
	n := len(arr)
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	if k >= n {
		return maxVal
	}

	current := arr[0]
	wins := 0

	for i := 1; i < n; i++ {
		if current > arr[i] {
			wins++
		} else {
			current = arr[i]
			wins = 1
		}

		if wins == k {
			return current
		}
	}

	// If we've gone through the whole array, the max element wins
	return maxVal
}
```
