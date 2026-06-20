# 2728 — Count Houses In A Circular Street

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountHousesInACircularStreet(street Street) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

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
