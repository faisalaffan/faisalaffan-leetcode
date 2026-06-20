# 0489 — Robot Room Cleaner

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func cleanRoomClient(robot *Robot) 
```

> **💡 Hint:** DFS backtracking with simulated robot API.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, DFS, Backtracking

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #489: Robot Room Cleaner
// https://leetcode.com/problems/robot-room-cleaner/
// Difficulty: Hard [Paid]
// Approach: DFS backtracking with simulated robot API.
// The robot has 4 methods: move(), turnLeft(), turnRight(), clean().
// We explore the room using DFS, keeping track of visited cells.
// Directions: 0=up, 1=right, 2=down, 3=left

import (
	"fmt"
)

func main() {
	fmt.Println("489 - Robot Room Cleaner")

	// Test with a simulated room
	room := [][]int{
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	row, col := 1, 3

	robot := NewSimulatedRobot(room, row, col)
	cleanRoom(robot)

	cleanedCount := 0
	for _, r := range robot.sim.cleaned {
		for _, cell := range r {
			if cell {
				cleanedCount++
			}
		}
	}
	fmt.Println("Cleaning complete. Cleaned:", cleanedCount)
	fmt.Println("Total cleaned cells in room:")
	count := 0
	for _, r := range room {
		for _, cell := range r {
			if cell == 1 {
				count++
			}
		}
	}
	fmt.Printf("Expected to clean: %d cells\n", count)
}

// Robot interface as provided by LeetCode
type Robot struct {
	// This would be the LeetCode API
	// For our simulation, we expose the methods
	sim *simulatedRobot
}

func (r *Robot) Move() bool     { return r.sim.move() }
func (r *Robot) TurnLeft()      { r.sim.turnLeft() }
func (r *Robot) TurnRight()     { r.sim.turnRight() }
func (r *Robot) Clean()         { r.sim.clean() }

// Direction vectors: 0=up, 1=right, 2=down, 3=left
var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

// cleanRoomClient is the actual solution function
func cleanRoomClient(robot *Robot) {
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[[2]int]bool)
	backtrack(robot, 0, 0, 0, visited) // start at (0, 0), facing up
}

func backtrack(robot *Robot, r, c, dir int, visited map[[2]int]bool) {
	key := [2]int{r, c}
	if visited[key] {
		return
	}
	visited[key] = true
	robot.Clean()

	// Try all 4 directions
	for i := 0; i < 4; i++ {
		newDir := (dir + i) % 4
		nr := r + dirs[newDir][0]
		nc := c + dirs[newDir][1]

		if robot.Move() {
			backtrack(robot, nr, nc, newDir, visited)
			// Go back
			robot.TurnLeft()
			robot.TurnLeft()
			robot.Move()
			robot.TurnLeft()
			robot.TurnLeft()
		} else {
			// Can't move, turn right to try next direction
			robot.TurnRight()
		}
	}

	// Turn back to original direction before returning
	// Actually, since we return to original position, the direction
	// should be restored. Let's adjust: after exploring all 4 directions,
	// we should be back at original direction.
	// Each Move() returns us, and unsuccessful moves just TurnRight.
	// At the end, we've turned right 4 times = full circle = back to original.
}

// For LeetCode submission, the function is normally:
// func cleanRoom(robot *Robot) { ... }
// We'll use this as the main function
func cleanRoom(robot *Robot) {
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[[2]int]bool)
	dfsClean(robot, 0, 0, 0, visited)
}

func dfsClean(robot *Robot, r, c, dir int, visited map[[2]int]bool) {
	key := [2]int{r, c}
	if visited[key] {
		return
	}
	visited[key] = true
	robot.Clean()

	// Try all 4 directions in order
	for i := 0; i < 4; i++ {
		newDir := (dir + i) % 4
		nr := r + dirs[newDir][0]
		nc := c + dirs[newDir][1]

		if robot.Move() {
			dfsClean(robot, nr, nc, newDir, visited)
			// Backtrack: turn 180, move, turn 180
			robot.TurnLeft()
			robot.TurnLeft()
			robot.Move()
			// Restore direction: turn right and then left = 180 total, but
			// we need to go back to original dir. Let's see:
			// Before backtrack, robot is at new cell facing newDir.
			// After TurnLeft+TurnLeft, faces (newDir+2)%4.
			// After Move(), back at (r,c) facing (newDir+2)%4.
			// We need to restore to original dir.
			// Turn right twice to go from (newDir+2)%4 to dir.
			robot.TurnRight()
			robot.TurnRight()
		}
		// Turn right to try next direction
		robot.TurnRight()
	}

	// After loop: we have turned right 4 times = back to original dir
}

// ---- Simulation for testing ----

type simulatedRobot struct {
	room    [][]int
	r, c    int
	dir     int // 0=up, 1=right, 2=down, 3=left
	visited map[[2]int]bool
	cleaned [][]bool
}

func NewSimulatedRobot(room [][]int, r, c int) *Robot {
  // Membuat matriks/slice 2D untuk DP
	cleaned := make([][]bool, len(room))
  // Range loop: iterasi dengan indeks + nilai
	for i := range cleaned {
		cleaned[i] = make([]bool, len(room[i]))
	}
	return &Robot{
		sim: &simulatedRobot{
			room:    room,
			r:       r,
			c:       c,
			dir:     0, // start facing up
			visited: make(map[[2]int]bool),
			cleaned: cleaned,
		},
	}
}

func (s *simulatedRobot) move() bool {
	nr := s.r + dirs[s.dir][0]
	nc := s.c + dirs[s.dir][1]

	// Check bounds and walls
	if nr < 0 || nr >= len(s.room) || nc < 0 || nc >= len(s.room[0]) {
		return false
	}
	if s.room[nr][nc] == 0 {
		return false
	}

	s.r = nr
	s.c = nc
	s.visited[[2]int{nr, nc}] = true
	return true
}

func (s *simulatedRobot) turnLeft() {
	s.dir = (s.dir + 3) % 4
}

func (s *simulatedRobot) turnRight() {
	s.dir = (s.dir + 1) % 4
}

func (s *simulatedRobot) clean() {
	if s.r >= 0 && s.r < len(s.room) && s.c >= 0 && s.c < len(s.room[0]) {
		s.cleaned[s.r][s.c] = true
		s.visited[[2]int{s.r, s.c}] = true
	}
}
```
