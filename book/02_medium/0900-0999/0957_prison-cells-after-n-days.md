# 0957 — Prison Cells After N Days

## Deskripsi

**Soal:** [0957. Prison Cells After N Days](https://leetcode.com/problems/prison-cells-after-n-days/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func prisonAfterNDays(cells []int, n int) []int`

## Solusi Go

```go
package main

// LeetCode #957: Prison Cells After N Days
// https://leetcode.com/problems/prison-cells-after-n-days/
// Difficulty: Medium

import "fmt"

// Time: O(1) | Space: O(1)
func prisonAfterNDays(cells []int, n int) []int {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[[8]int]int)
	cycle := false

	for n > 0 {
		key := toArray(cells)
		if day, ok := seen[key]; ok && !cycle {
			n %= day - n
			cycle = true
		}
		seen[key] = n

		if n > 0 {
			n--
			cells = nextDay(cells)
		}
	}

	return cells
}

func toArray(cells []int) [8]int {
	return [8]int{cells[0], cells[1], cells[2], cells[3], cells[4], cells[5], cells[6], cells[7]}
}

func nextDay(cells []int) []int {
  // Membuat slice untuk menyimpan hasil
	next := make([]int, 8)
	for i := 1; i < 7; i++ {
		if cells[i-1] == cells[i+1] {
			next[i] = 1
		}
	}
	return next
}

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}
```
