# 1947 — Maximum Compatibility Score Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaxCompatibilitySum(students [][]int, mentors [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(m! * n * k) where m = len(students) <= 8, Space: O(m)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1947: Maximum Compatibility Score Sum
// https://leetcode.com/problems/maximum-compatibility-score-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxCompatibilitySum([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}))
	fmt.Println(MaxCompatibilitySum([][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}))
}

// Time: O(m! * n * k) where m = len(students) <= 8, Space: O(m)
func MaxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
  // Matriks 2D
	score := make([][]int, m)
	for i := 0; i < m; i++ {
		score[i] = make([]int, m)
		for j := 0; j < m; j++ {
			s := 0
			for k := 0; k < len(students[i]); k++ {
				if students[i][k] == mentors[j][k] {
					s++
				}
			}
			score[i][j] = s
		}
	}

	used := make([]bool, m)
	return backtrackMax(0, m, score, used)
}

func backtrackMax(student int, m int, score [][]int, used []bool) int {
	if student == m {
		return 0
	}
	maxScore := 0
	for mentor := 0; mentor < m; mentor++ {
		if !used[mentor] {
			used[mentor] = true
			cur := score[student][mentor] + backtrackMax(student+1, m, score, used)
			if cur > maxScore {
				maxScore = cur
			}
			used[mentor] = false
		}
	}
	return maxScore
}
```
