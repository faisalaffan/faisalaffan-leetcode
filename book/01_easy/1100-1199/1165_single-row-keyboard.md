# 1165 — Single Row Keyboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func calculateTime(keyboard, word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1165: Single-Row Keyboard
// https://leetcode.com/problems/single-row-keyboard/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(calculateTime("abcdefghijklmnopqrstuvwxyz", "cba")) // 4
	fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")) // 73
}

// LeetCode submission: calculateTime
func calculateTime(keyboard, word string) int {
	pos := [26]int{}
  // Linear scan O(n)
	for i := 0; i < len(keyboard); i++ {
		pos[keyboard[i]-'a'] = i
	}
	ans, cur := 0, 0
  // Linear scan O(n)
	for i := 0; i < len(word); i++ {
		next := pos[word[i]-'a']
		if next > cur {
			ans += next - cur
		} else {
			ans += cur - next
		}
		cur = next
	}
	return ans
}
```
