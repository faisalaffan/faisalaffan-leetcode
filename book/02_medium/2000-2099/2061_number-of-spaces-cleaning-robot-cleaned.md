# 2061 — Number Of Spaces Cleaning Robot Cleaned

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfCleanRooms(room [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2061: Number of Spaces Cleaning Robot Cleaned
// https://leetcode.com/problems/number-of-spaces-cleaning-robot-cleaned/
// Difficulty: Medium [Paid]
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func numberOfCleanRooms(room [][]int) int {
	m, n := len(room), len(room[0])
  // Membuat matriks/slice 2D untuk DP
	visited := make([][][4]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([][4]bool, n)
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, down, left, up
  // Membuat map (HashMap) — pencarian O(1)
	cleaned := make(map[[2]int]bool)
	dir := 0
	r, c := 0, 0
	cleaned[[2]int{0, 0}] = true

	for {
		if visited[r][c][dir] {
			break
		}
		visited[r][c][dir] = true

		// Try to move in current direction
		nextR, nextC := r+dirs[dir][0], c+dirs[dir][1]

		if nextR >= 0 && nextR < m && nextC >= 0 && nextC < n && room[nextR][nextC] == 0 {
			r, c = nextR, nextC
			cleaned[[2]int{r, c}] = true
		} else {
			dir = (dir + 1) % 4
		}
	}

	return len(cleaned)
}

func main() {
	// Test case 1
	room1 := [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}}
	fmt.Println("Test 1:", numberOfCleanRooms(room1))
	// Expected: 7

	// Test case 2
	room2 := [][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}
	fmt.Println("Test 2:", numberOfCleanRooms(room2))
	// Expected: 1

	// Test case 3
	room3 := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println("Test 3:", numberOfCleanRooms(room3))
	// Expected: 9
}
```
