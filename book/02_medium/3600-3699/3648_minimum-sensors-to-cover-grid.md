# 3648 — Minimum Sensors To Cover Grid

## Deskripsi

**Soal:** [3648. Minimum Sensors To Cover Grid](https://leetcode.com/problems/minimum-sensors-to-cover-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minimumSensorsToCoverGrid(n int, m int, k int) int`

## Solusi Go

```go
package main

// LeetCode #3648: Minimum Sensors to Cover Grid
// https://leetcode.com/problems/minimum-sensors-to-cover-grid/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func minimumSensorsToCoverGrid(n int, m int, k int) int {
	s := 2*k + 1
	rows := (n + s - 1) / s // ceil division
	cols := (m + s - 1) / s
	return rows * cols
}

func main() {
	fmt.Println(minimumSensorsToCoverGrid(5, 5, 1))
	fmt.Println(minimumSensorsToCoverGrid(2, 2, 2))
	fmt.Println(minimumSensorsToCoverGrid(3, 4, 0))
}
```
