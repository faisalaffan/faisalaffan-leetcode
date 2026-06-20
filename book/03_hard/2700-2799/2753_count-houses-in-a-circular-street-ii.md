# 2753 — Count Houses In A Circular Street Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func newStreet(doors []bool) *Street
```

> **💡 Hint:** Use the Street interface operations.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
