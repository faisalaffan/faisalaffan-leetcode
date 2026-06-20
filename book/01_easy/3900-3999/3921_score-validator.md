# 3921 — Score Validator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ScoreValidator(events []string) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3921: Score Validator
// https://leetcode.com/problems/score-validator/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ScoreValidator([]string{"1", "4", "W", "6", "WD"}))
	fmt.Println(ScoreValidator([]string{"WD", "NB", "0", "4", "4"}))
	fmt.Println(ScoreValidator([]string{"W", "W", "W", "W", "W", "W", "W", "W", "W", "W", "W"}))
}

// Time: O(n)
// Space: O(1)
func ScoreValidator(events []string) []int {
	score, counter := 0, 0
	for _, e := range events {
		if counter == 10 {
			break
		}
		if e == "W" {
			counter++
		} else if e == "WD" || e == "NB" {
			score++
		} else {
			v, _ := strconv.Atoi(e)
			score += v
		}
	}
	return []int{score, counter}
}
```
