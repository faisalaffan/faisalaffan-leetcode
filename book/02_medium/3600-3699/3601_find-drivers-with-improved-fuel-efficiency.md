# 3601 — Find Drivers With Improved Fuel Efficiency

## Deskripsi

**Soal:** [3601. Find Drivers With Improved Fuel Efficiency](https://leetcode.com/problems/find-drivers-with-improved-fuel-efficiency/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3601: Find Drivers with Improved Fuel Efficiency
// https://leetcode.com/problems/find-drivers-with-improved-fuel-efficiency/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	before := []float64{10, 15, 20}
	after := []float64{12, 14, 22}
	fmt.Println("Test 1:", FindDriversWithImprovedFuelEfficiency(before, after))
	// Test case 2
	before2 := []float64{10, 10}
	after2 := []float64{9, 11}
	fmt.Println("Test 2:", FindDriversWithImprovedFuelEfficiency(before2, after2))
	// Test case 3
	before3 := []float64{5}
	after3 := []float64{6}
	fmt.Println("Test 3:", FindDriversWithImprovedFuelEfficiency(before3, after3))
}

func FindDriversWithImprovedFuelEfficiency(before, after []float64) int {
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(before) && i < len(after); i++ {
		if after[i] > before[i] {
			count++
		}
	}
	return count
}
```
