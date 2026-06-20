# 0849 — Maximize Distance To Closest Person

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximizeDistanceToClosestPerson(seats []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #849: Maximize Distance to Closest Person
// https://leetcode.com/problems/maximize-distance-to-closest-person/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximizeDistanceToClosestPerson([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println(MaximizeDistanceToClosestPerson([]int{1, 0, 0, 0}))
	fmt.Println(MaximizeDistanceToClosestPerson([]int{0, 1}))
}

// Time: O(n) | Space: O(1)
func MaximizeDistanceToClosestPerson(seats []int) int {
	n := len(seats)
	ans := 0
	lastPerson := -1

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if lastPerson == -1 {
				ans = i
			} else {
				dist := (i - lastPerson) / 2
				if dist > ans {
					ans = dist
				}
			}
			lastPerson = i
		}
	}

	// Check distance from last person to the end
	if seats[n-1] == 0 {
		dist := n - 1 - lastPerson
		if dist > ans {
			ans = dist
		}
	}

	return ans
}
```
