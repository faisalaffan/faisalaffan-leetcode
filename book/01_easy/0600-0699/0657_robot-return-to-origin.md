# 0657 — Robot Return To Origin

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func judgeCircle(moves string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #657: Robot Return to Origin
// https://leetcode.com/problems/robot-return-to-origin/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))   // true
	fmt.Println(judgeCircle("LL"))   // false
	fmt.Println(judgeCircle(""))     // true
}

// judgeCircle returns true if the robot returns to origin after executing all moves.
// Time: O(n). Space: O(1).
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
```
