# 2728 — Count Houses In A Circular Street

## Deskripsi

**Soal:** [2728. Count Houses In A Circular Street](https://leetcode.com/problems/count-houses-in-a-circular-street/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2728: Count Houses in a Circular Street
// https://leetcode.com/problems/count-houses-in-a-circular-street/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: Interactive problem adapted to Go. Uses street interface.

import "fmt"

func main() {
	// Simulate: street with 3 houses, open first door
	street := &StreetSim{houses: []bool{true, false, false}, idx: 0}
	fmt.Println(CountHousesInACircularStreet(street))
}

type Street interface {
	OpenDoor()
	CloseDoor()
	IsDoorOpen() bool
	MoveRight()
	MoveLeft()
}

type StreetSim struct {
	houses []bool
	idx    int
}

func (s *StreetSim) OpenDoor()    { s.houses[s.idx] = true }
func (s *StreetSim) CloseDoor()   { s.houses[s.idx] = false }
func (s *StreetSim) IsDoorOpen() bool { return s.houses[s.idx] }
func (s *StreetSim) MoveRight()   { s.idx = (s.idx + 1) % len(s.houses) }
func (s *StreetSim) MoveLeft()    { s.idx = (s.idx - 1 + len(s.houses)) % len(s.houses) }

func CountHousesInACircularStreet(street Street) int {
	// Open current door as marker
	street.OpenDoor()

	count := 0
	for {
		street.MoveRight()
		count++
		if street.IsDoorOpen() {
			break
		}
	}

	street.CloseDoor()
	return count
}
```
