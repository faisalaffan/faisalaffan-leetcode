# 1868 — Product Of Two Run Length Encoded Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n1 + n2), Space: O(n1 + n2) for result  |  **Ruang:** O(n1 + n2) for result


## 💻 Solusi Go

```go
package main

// LeetCode #1868: Product of Two Run-Length Encoded Arrays
// https://leetcode.com/problems/product-of-two-run-length-encoded-arrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(FindRLEArray([][]int{{1, 3}, {2, 3}}, [][]int{{6, 3}, {3, 3}}))
	fmt.Println(FindRLEArray([][]int{{1, 2}, {2, 2}}, [][]int{{5, 1}, {2, 1}}))
}

// Time: O(n1 + n2), Space: O(n1 + n2) for result
func FindRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
  // Matriks 2D
	result := make([][]int, 0)
	i, j := 0, 0

	for i < len(encoded1) && j < len(encoded2) {
		val := encoded1[i][0] * encoded2[j][0]
		minFreq := min(encoded1[i][1], encoded2[j][1])

		// Merge with previous if same value
		if len(result) > 0 && result[len(result)-1][0] == val {
			result[len(result)-1][1] += minFreq
		} else {
			result = append(result, []int{val, minFreq})
		}

		encoded1[i][1] -= minFreq
		encoded2[j][1] -= minFreq

		if encoded1[i][1] == 0 {
			i++
		}
		if encoded2[j][1] == 0 {
			j++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
