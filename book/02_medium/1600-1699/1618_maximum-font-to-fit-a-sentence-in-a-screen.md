# 1618 — Maximum Font To Fit A Sentence In A Screen

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaxFont(text string, w int, h int, fonts []int, fontInfo interface{ getFontWidth(int) int; getFontHeight(int) int }) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** O(log N * L), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1618: Maximum Font to Fit a Sentence in a Screen
// https://leetcode.com/problems/maximum-font-to-fit-a-sentence-in-a-screen/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Problem: Given a screen width and height, find the maximum font size
	// that can display all characters of a string within the screen.
	// We have a FontInfo API that gives width and height for a font size.

	sentence := "Hello World"
	width := 80
	height := 30

	// Available font sizes
	fonts := []int{6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 36, 48, 72}

	// Mock font info
	fontInfo := &mockFontInfo{}

	maxFont := MaxFont(sentence, width, height, fonts, fontInfo)
	fmt.Println("Maximum font size:", maxFont)
}

type FontInfo interface {
	getFontWidth(fontSize int) int
	getFontHeight(fontSize int) int
}

type mockFontInfo struct{}

func (m *mockFontInfo) getFontWidth(fontSize int) int {
	// Approximate: width ~ font_size * 0.6 per character
	return fontSize * 6 / 10
}

func (m *mockFontInfo) getFontHeight(fontSize int) int {
	return fontSize
}

func MaxFont(text string, w int, h int, fonts []int, fontInfo interface{ getFontWidth(int) int; getFontHeight(int) int }) int {
	// Time: O(log N * L), Space: O(1)
	// Binary search on font sizes
	left, right := 0, len(fonts)-1
	result := -1

  // Binary search loop
	for left <= right {
		mid := left + (right-left)/2
		fontSize := fonts[mid]

		charWidth := fontInfo.getFontWidth(fontSize)
		charHeight := fontInfo.getFontHeight(fontSize)

		// Check if text fits
		charsPerLine := w / charWidth
		if charsPerLine == 0 {
			right = mid - 1
			continue
		}

		lines := (len(text) + charsPerLine - 1) / charsPerLine
		if lines*charHeight <= h {
			result = fontSize
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
```
