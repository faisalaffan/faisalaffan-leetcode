# 2976 — Minimum Cost To Convert String I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCostConvert(source string, target string, original []byte, changed []byte, cost []int) (ans int64)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(26^3 + n)  
**Kompleksitas Ruang:** O(26^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat matriks/slice 2D untuk DP
	g := make([][]int, 26)
  // Range loop: iterasi dengan indeks + nilai
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
  // Loop linear O(n): iterasi setiap elemen
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
  // Loop linear O(n): iterasi setiap elemen
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
