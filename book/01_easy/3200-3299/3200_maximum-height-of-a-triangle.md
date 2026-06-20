# 3200 — Maximum Height Of A Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxHeight(red, blue int, firstRed bool) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Trie

**Waktu:** O(sqrt(n)). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Trie** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3200: Maximum Height of a Triangle
// https://leetcode.com/problems/maximum-height-of-a-triangle/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumHeightOfATriangle(2, 4))
	fmt.Println(MaximumHeightOfATriangle(2, 1))
	fmt.Println(MaximumHeightOfATriangle(10, 10))
}

// maxHeight tries building a triangle starting with the given first color.
func maxHeight(red, blue int, firstRed bool) int {
	h := 0
	need := 1
	for {
		if firstRed {
			if red < need {
				break
			}
			red -= need
		} else {
			if blue < need {
				break
			}
			blue -= need
		}
		h++
		need++
		firstRed = !firstRed
	}
	return h
}

// MaximumHeightOfATriangle returns the maximum height of a triangle using red and blue balls.
// Time: O(sqrt(n)). Space: O(1).
func MaximumHeightOfATriangle(red int, blue int) int {
	h1 := maxHeight(red, blue, true)
	h2 := maxHeight(red, blue, false)
	if h1 > h2 {
		return h1
	}
	return h2
}
```
