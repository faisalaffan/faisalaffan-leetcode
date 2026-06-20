# 2753 — Count Houses In A Circular Street Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func newStreet(doors []bool) *Street`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2753: Count Houses in a Circular Street II
// https://leetcode.com/problems/count-houses-in-a-circular-street-ii/
// Difficulty: Hard [Paid]
//
// Approach: Use the Street interface operations.
// 1. Open door at starting house.
// 2. Move right, closing every door we pass.
// 3. Count steps until we find an open door (the start).
// 4. Return count + 1 (for the start house).

import "fmt"

// Street simulates the problem's Street interface.
type Street struct {
	doors []bool
	pos   int
}

func newStreet(doors []bool) *Street {
	return &Street{doors: doors, pos: 0}
}

func (s *Street) openDoor()  { s.doors[s.pos] = true }
func (s *Street) closeDoor() { s.doors[s.pos] = false }
func (s *Street) isDoorOpen() bool { return s.doors[s.pos] }
func (s *Street) moveRight() { s.pos = (s.pos + 1) % len(s.doors) }

// houseCount counts houses on the circular street using only the Street API.
func houseCount(street *Street) int {
	// Open door at start position
	street.openDoor()
	street.moveRight()
	steps := 0

	// Walk until we find the open start door
	// Close every door we pass to avoid counting them later
	for !street.isDoorOpen() {
		street.closeDoor()
		street.moveRight()
		steps++
	}

	// Close the starting door (cleanup)
	street.closeDoor()
	// +1 accounts for the starting house itself
	return steps + 1
}

func main() {
	// 5 houses, all doors initially closed
	fmt.Println(houseCount(newStreet([]bool{false, false, false, false, false})))
	// Single house
	fmt.Println(houseCount(newStreet([]bool{false})))
	// Two houses
	fmt.Println(houseCount(newStreet([]bool{false, false})))
	// Three houses
	fmt.Println(houseCount(newStreet([]bool{false, false, false})))
	// Ten houses
	fmt.Println(houseCount(newStreet([]bool{false, false, false, false, false, false, false, false, false, false})))
	// Some doors already open (should not affect since we close as we go)
	fmt.Println(houseCount(newStreet([]bool{true, false, false})))
}
```
