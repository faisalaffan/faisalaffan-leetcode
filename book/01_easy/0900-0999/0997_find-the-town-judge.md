# 0997 — Find The Town Judge

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findJudge(n int, trust [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(E). Space: O(V).  |  **Ruang:** O(V).


## 💻 Solusi Go

```go
package main

// LeetCode #997: Find the Town Judge
// https://leetcode.com/problems/find-the-town-judge/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))             // 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))    // 3
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // -1
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}))    // -1
}

// findJudge finds the town judge (trusted by everyone, trusts no one).
// Time: O(E). Space: O(V).
func findJudge(n int, trust [][]int) int {
  // Alokasi slice
	inDeg := make([]int, n+1)
  // Alokasi slice
	outDeg := make([]int, n+1)
	for _, t := range trust {
		outDeg[t[0]]++
		inDeg[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if inDeg[i] == n-1 && outDeg[i] == 0 {
			return i
		}
	}
	return -1
}
```
