# 1947 — Maximum Compatibility Score Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxCompatibilitySum(students [][]int, mentors [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(m! * n * k) where m = len(students) <= 8, Space: O(m)  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat matriks/slice 2D untuk DP
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
