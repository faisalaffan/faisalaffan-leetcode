# 0500 — Keyboard Row

## Deskripsi

**Soal:** [0500. Keyboard Row](https://leetcode.com/problems/keyboard-row/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n*k), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func KeyboardRow(words []string) []string`

## Solusi Go

```go
package main

// LeetCode #500: Keyboard Row
// https://leetcode.com/problems/keyboard-row/
// Difficulty: Easy

import "fmt"

// Time: O(n*k), Space: O(n)
func KeyboardRow(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
  // Membuat map untuk pencarian O(1): key → value
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
