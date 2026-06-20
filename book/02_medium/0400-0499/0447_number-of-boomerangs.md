# 0447 — Number Of Boomerangs

## Deskripsi

**Soal:** [0447. Number Of Boomerangs](https://leetcode.com/problems/number-of-boomerangs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfBoomerangs(points [][]int) int`

## Solusi Go

```go
package main

// LeetCode #447: Number of Boomerangs
// https://leetcode.com/problems/number-of-boomerangs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	total := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(points); i++ {
  // Membuat map untuk pencarian O(1): key → value
		distCount := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := dx*dx + dy*dy
			distCount[dist]++
		}
		for _, count := range distCount {
			total += count * (count - 1)
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numberOfBoomerangs([][]int{{0, 0}}))
	// Expected: 0
}
```
