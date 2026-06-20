# 1773 — Count Items Matching A Rule

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CountMatches(items [][]string, ruleKey string, ruleValue string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1773: Count Items Matching a Rule
// https://leetcode.com/problems/count-items-matching-a-rule/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	idx := 0
	switch ruleKey {
	case "color":
		idx = 1
	case "name":
		idx = 2
	}
	count := 0
	for _, item := range items {
		if item[idx] == ruleValue {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver"))
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}
```
