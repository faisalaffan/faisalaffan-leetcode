# 3680 — Generate Schedule

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func generateSchedule(n int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)


## 💻 Solusi Go

```go
package main

// LeetCode #3680: Generate Schedule
// https://leetcode.com/problems/generate-schedule/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func generateSchedule(n int) [][]int {
	if n < 5 {
		return [][]int{}
	}

	totalMatches := n * (n - 1)
  // Matriks 2D
	schedule := make([][]int, 0, totalMatches)

	// Phase 1: offset 2 to n-2
	for offset := 2; offset < n-1; offset++ {
		for team := 0; team < n; team++ {
			schedule = append(schedule, []int{team, (team + offset) % n})
		}
	}

	// Phase 2: wrap-around pairs
	for team := 0; team < n; team++ {
		schedule = append(schedule, []int{team, (team + 1) % n})
		schedule = append(schedule, []int{(team + 4) % n, (team + 3) % n})
	}

	return schedule
}

func main() {
	fmt.Println(len(generateSchedule(5)))
	fmt.Println(len(generateSchedule(3)))
	fmt.Println(len(generateSchedule(6)))
}
```
