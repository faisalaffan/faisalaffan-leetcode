# 1560 — Most Visited Sector In A Circular Track

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func mostVisited(n int, rounds []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1560: Most Visited Sector in a Circular Track
// https://leetcode.com/problems/most-visited-sector-in-a-circular-track/
// Difficulty: Easy
//
// LeetCode submission: func mostVisited(n int, rounds []int) []int

import "fmt"

func main() {
	fmt.Println(MostVisitedSectorInACircularTrack(4, []int{1, 3, 1, 2})) // [1 2]
	fmt.Println(MostVisitedSectorInACircularTrack(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2})) // [2]
	fmt.Println(MostVisitedSectorInACircularTrack(7, []int{1, 3, 5, 7})) // [1 2 3 4 5 6 7]
}

// Time: O(n), Space: O(n)
func MostVisitedSectorInACircularTrack(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]
	if start <= end {
  // Alokasi slice
		res := make([]int, end-start+1)
  // Range loop
		for i := range res {
			res[i] = start + i
		}
		return res
	}
  // Alokasi slice
	res := make([]int, 0, n-end+start)
	for i := 1; i <= end; i++ {
		res = append(res, i)
	}
	for i := start; i <= n; i++ {
		res = append(res, i)
	}
	return res
}
```
