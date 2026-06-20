# 2879 — Display The First Three Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func DisplayTheFirstThreeRows(df [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = min(3, rows)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2879: Display the First Three Rows
// https://leetcode.com/problems/display-the-first-three-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we return the first 3 rows of the dataframe.

import "fmt"

func main() {
	// LeetCode name: selectFirstRows
	fmt.Println(DisplayTheFirstThreeRows([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}}))
	// [[1 15] [2 11] [3 11]]

	fmt.Println(DisplayTheFirstThreeRows([][]int{{1, 15}}))
	// [[1 15]]
}

// Time: O(n) where n = min(3, rows) | Space: O(1)
// LeetCode submission name: selectFirstRows
func DisplayTheFirstThreeRows(df [][]int) [][]int {
	if len(df) > 3 {
		return df[:3]
	}
	return df
}
```
