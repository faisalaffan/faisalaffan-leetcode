# 0500 — Keyboard Row

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func KeyboardRow(words []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n*k), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #500: Keyboard Row
// https://leetcode.com/problems/keyboard-row/
// Difficulty: Easy

import "fmt"

// Time: O(n*k), Space: O(n)
func KeyboardRow(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
  // HashMap: O(1) lookup
	rowMap := make(map[byte]int)
	for i, row := range rows {
		for j := 0; j < len(row); j++ {
			rowMap[row[j]] = i
		}
	}
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		targetRow := rowMap[toLower(word[0])]
		sameRow := true
		for i := 1; i < len(word); i++ {
			if rowMap[toLower(word[i])] != targetRow {
				sameRow = false
				break
			}
		}
		if sameRow {
			result = append(result, word)
		}
	}
	return result
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(KeyboardRow([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(KeyboardRow([]string{"omk"}))
	fmt.Println(KeyboardRow([]string{"adsdf", "sfd"}))
}
```
