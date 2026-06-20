# 2739 — Total Distance Traveled

## Deskripsi

**Soal:** [2739. Total Distance Traveled](https://leetcode.com/problems/total-distance-traveled/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2739: Total Distance Traveled
// https://leetcode.com/problems/total-distance-traveled/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalDistanceTraveled(5, 10))
	fmt.Println(TotalDistanceTraveled(1, 2))
}

func TotalDistanceTraveled(mainTank int, additionalTank int) int {
	total := 0
	for mainTank > 0 {
		if mainTank >= 5 {
			mainTank -= 5
			total += 50
			if additionalTank > 0 {
				additionalTank--
				mainTank++
			}
		} else {
			total += mainTank * 10
			mainTank = 0
		}
	}
	return total
}
```
