# 2383 — Minimum Hours Of Training To Win A Competition

## Deskripsi

**Soal:** [2383. Minimum Hours Of Training To Win A Competition](https://leetcode.com/problems/minimum-hours-of-training-to-win-a-competition/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2383: Minimum Hours of Training to Win a Competition
// https://leetcode.com/problems/minimum-hours-of-training-to-win-a-competition/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumHoursOfTrainingToWinACompetition(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1})) // 8
	fmt.Println(MinimumHoursOfTrainingToWinACompetition(2, 4, []int{1}, []int{3}))                   // 0
}

func MinimumHoursOfTrainingToWinACompetition(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	totalHours := 0
	curEnergy := initialEnergy
	curExp := initialExperience

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(energy); i++ {
		// Energy
		if curEnergy <= energy[i] {
			needed := energy[i] - curEnergy + 1
			totalHours += needed
			curEnergy += needed
		}
		curEnergy -= energy[i]

		// Experience
		if curExp <= experience[i] {
			needed := experience[i] - curExp + 1
			totalHours += needed
			curExp += needed
		}
		curExp += experience[i]
	}
	return totalHours
}
```
