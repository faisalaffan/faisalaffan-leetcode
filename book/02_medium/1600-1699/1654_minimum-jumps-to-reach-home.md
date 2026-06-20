# 1654 — Minimum Jumps To Reach Home

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumJumps(forbidden []int, a int, b int, x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, BFS

**Kompleksitas Waktu:** O(limit), Space: O(limit)  
**Kompleksitas Ruang:** O(limit)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1654: Minimum Jumps to Reach Home
// https://leetcode.com/problems/minimum-jumps-to-reach-home/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	fmt.Println(MinimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11))
	fmt.Println(MinimumJumps([]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7))
}

func MinimumJumps(forbidden []int, a int, b int, x int) int {
	// Time: O(limit), Space: O(limit)
	// BFS with state (position, direction) where direction 0=right, 1=left
  // Membuat map (HashMap) — pencarian O(1)
	forbiddenSet := make(map[int]bool)
	for _, f := range forbidden {
		forbiddenSet[f] = true
	}

	// Upper bound: we shouldn't go beyond max(x, max(forbidden)) + a + b
	limit := x
	for _, f := range forbidden {
		if f > limit {
			limit = f
		}
	}
	limit += a + b

	type state struct {
		pos int
		dir int // 0 = came from left (can go either way), 1 = came from right (can only go right)
	}

  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[struct{ pos, dir int }]bool)
	queue := []state{{pos: 0, dir: 0}}
	visited[struct{ pos, dir int }{0, 0}] = true
	jumps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.pos == x {
				return jumps
			}

			// Move forward (right)
			nextPos := curr.pos + a
			if nextPos <= limit && !forbiddenSet[nextPos] {
				key := struct{ pos, dir int }{nextPos, 0}
				if !visited[key] {
					visited[key] = true
					queue = append(queue, state{pos: nextPos, dir: 0})
				}
			}

			// Move backward (left) only if we didn't come from a backward move
			if curr.dir != 1 {
				nextPos = curr.pos - b
				if nextPos >= 0 && !forbiddenSet[nextPos] {
					key := struct{ pos, dir int }{nextPos, 1}
					if !visited[key] {
						visited[key] = true
						queue = append(queue, state{pos: nextPos, dir: 1})
					}
				}
			}
		}
		jumps++
	}

	return -1
}
```
