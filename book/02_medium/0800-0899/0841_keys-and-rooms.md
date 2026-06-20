# 0841 — Keys And Rooms

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func KeysAndRooms(rooms [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n + k) where k = total keys  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #841: Keys and Rooms
// https://leetcode.com/problems/keys-and-rooms/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KeysAndRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(KeysAndRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(KeysAndRooms([][]int{{2}, {}, {1}}))
}

// Time: O(n + k) where k = total keys | Space: O(n)
func KeysAndRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	stack := []int{0}
	visited[0] = true
	count := 1

	for len(stack) > 0 {
		room := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, key := range rooms[room] {
			if !visited[key] {
				visited[key] = true
				count++
				stack = append(stack, key)
			}
		}
	}

	return count == n
}
```
