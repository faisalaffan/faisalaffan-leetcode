# 3001 — Minimum Moves To Capture The Queen

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3001: Minimum Moves to Capture The Queen
// https://leetcode.com/problems/minimum-moves-to-capture-the-queen/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minMovesToCaptureTheQueen(1, 1, 8, 8, 2, 3))
	fmt.Println(minMovesToCaptureTheQueen(5, 3, 3, 4, 5, 2))
}

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	// Rook and queen on same row, bishop not in between
	if a == e && !(a == c && (d-b)*(d-f) < 0) {
		return 1
	}
	// Rook and queen on same column, bishop not in between
	if b == f && !(b == d && (c-a)*(c-e) < 0) {
		return 1
	}
	// Bishop and queen on same diagonal, rook not in between
	if c+d == e+f && !(a+b == e+f && (a-c)*(a-e) < 0) {
		return 1
	}
	if c-d == e-f && !(a-b == e-f && (a-c)*(a-e) < 0) {
		return 1
	}
	return 2
}
```
