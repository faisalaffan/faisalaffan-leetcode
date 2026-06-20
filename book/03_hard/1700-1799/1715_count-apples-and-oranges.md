# 1715 — Count Apples And Oranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countApplesAndOranges(boxes []Box, chests []Chest, boxChests []BoxChest) (int, int)`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1715: Count Apples and Oranges
// https://leetcode.com/problems/count-apples-and-oranges/
// Difficulty: Hard [Premium]

import "fmt"

type Box struct {
	BoxId       int
	AppleCount  int
	OrangeCount int
}

type Chest struct {
	ChestId     int
	AppleCount  int
	OrangeCount int
}

type BoxChest struct {
	BoxId   int
	ChestId int
}

func main() {
	boxes := []Box{
		{2, 2, 3},
		{3, 1, 4},
	}
	chests := []Chest{
		{1, 5, 6},
		{2, 7, 8},
	}
	boxChests := []BoxChest{
		{2, 1},
		{3, 2},
	}
	apples, oranges := countApplesAndOranges(boxes, chests, boxChests)
	fmt.Printf("Test 1 - Apples: %d, Oranges: %d\n", apples, oranges)
	fmt.Printf("Expected: apples=15, oranges=21\n\n")

	boxes2 := []Box{
		{1, 3, 5},
	}
	chests2 := []Chest{
		{3, 2, 4},
	}
	boxChests2 := []BoxChest{
		{1, 3},
	}
	apples2, oranges2 := countApplesAndOranges(boxes2, chests2, boxChests2)
	fmt.Printf("Test 2 - Apples: %d, Oranges: %d\n", apples2, oranges2)
	fmt.Printf("Expected: apples=5, oranges=9\n")
}

func countApplesAndOranges(boxes []Box, chests []Chest, boxChests []BoxChest) (int, int) {
  // HashMap: O(1) lookup
	chestMap := make(map[int]Chest)
	for _, c := range chests {
		chestMap[c.ChestId] = c
	}

  // HashMap: O(1) lookup
	boxChestMap := make(map[int]int)
	for _, bc := range boxChests {
		boxChestMap[bc.BoxId] = bc.ChestId
	}

	totalApples := 0
	totalOranges := 0

	for _, b := range boxes {
		totalApples += b.AppleCount
		totalOranges += b.OrangeCount
		if chestId, ok := boxChestMap[b.BoxId]; ok {
			if chest, ok := chestMap[chestId]; ok {
				totalApples += chest.AppleCount
				totalOranges += chest.OrangeCount
			}
		}
	}

	return totalApples, totalOranges
}
```
