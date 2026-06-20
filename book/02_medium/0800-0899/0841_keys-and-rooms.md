# 0841 — Keys And Rooms

## Deskripsi

**Soal:** [0841. Keys And Rooms](https://leetcode.com/problems/keys-and-rooms/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + k) where k = total keys  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
