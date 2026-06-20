# 2976 — Minimum Cost To Convert String I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumCostConvert(source string, target string, original []byte, changed []byte, cost []int) (ans int64)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(26^3 + n)  |  **Ruang:** O(26^2)


## 💻 Solusi Go

```go
package main

// LeetCode #2976: Minimum Cost to Convert String I
// https://leetcode.com/problems/minimum-cost-to-convert-string-i/
// Difficulty: Medium
// Time: O(26^3 + n) | Space: O(26^2)

import "fmt"

func main() {
	fmt.Println(minimumCostConvert("abcd", "acbe", []byte("ab"), []byte("ce"), []int{5, 3}))
	fmt.Println(minimumCostConvert("aaaa", "bbbb", []byte("a"), []byte("b"), []int{2}))
}

func minimumCostConvert(source string, target string, original []byte, changed []byte, cost []int) (ans int64) {
	const inf = 1 << 29
  // Matriks 2D
	g := make([][]int, 26)
  // Range loop
	for i := range g {
		g[i] = make([]int, 26)
		for j := range g[i] {
			if i == j {
				g[i][j] = 0
			} else {
				g[i][j] = inf
			}
		}
	}
  // Linear scan O(n)
	for i := 0; i < len(original); i++ {
		x := int(original[i] - 'a')
		y := int(changed[i] - 'a')
		z := cost[i]
		if z < g[x][y] {
			g[x][y] = z
		}
	}
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if g[i][k]+g[k][j] < g[i][j] {
					g[i][j] = g[i][k] + g[k][j]
				}
			}
		}
	}
  // Linear scan O(n)
	for i := 0; i < len(source); i++ {
		x := int(source[i] - 'a')
		y := int(target[i] - 'a')
		if x != y {
			if g[x][y] >= inf {
				return -1
			}
			ans += int64(g[x][y])
		}
	}
	return
}
```
