# 0756 — Pyramid Transition Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func pyramidTransition(bottom string, allowed []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** O(7^b) worst case where b is number of blocks  
**Kompleksitas Ruang:** O(7^b)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #756: Pyramid Transition Matrix
// https://leetcode.com/problems/pyramid-transition-matrix/
// Difficulty: Medium
// Time: O(7^b) worst case where b is number of blocks
// Space: O(7^b)

import "fmt"

func main() {
	fmt.Println(pyramidTransition("BCD", []string{"BCG", "CDE", "GEA", "FFF"}))
	fmt.Println(pyramidTransition("AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
}

func pyramidTransition(bottom string, allowed []string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	memo := make(map[string]bool)
  // Membuat map (HashMap) — pencarian O(1)
	patterns := make(map[string][]byte)

	for _, a := range allowed {
		key := a[:2]
		patterns[key] = append(patterns[key], a[2])
	}

	var dfs func(row string, next string, idx int) bool
	dfs = func(row string, next string, idx int) bool {
		if len(row) == 1 {
			return true
		}

		key := row + "#" + next
		if val, ok := memo[key]; ok {
			return val
		}

		if idx == len(row)-1 {
			if dfs(next, "", 0) {
				memo[key] = true
				return true
			}
			memo[key] = false
			return false
		}

		chars := patterns[row[idx:idx+2]]
		for _, c := range chars {
			if dfs(row, next+string(c), idx+1) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(bottom, "", 0)
}
```
