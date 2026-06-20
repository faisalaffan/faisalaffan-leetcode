# 1344 — Angle Between Hands Of A Clock

## Deskripsi

**Soal:** [1344. Angle Between Hands Of A Clock](https://leetcode.com/problems/angle-between-hands-of-a-clock/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1344: Angle Between Hands of a Clock
// https://leetcode.com/problems/angle-between-hands-of-a-clock/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(angleClock(12, 30)) // 165

	// Test case 2
	fmt.Println(angleClock(3, 30)) // 75

	// Test case 3
	fmt.Println(angleClock(3, 15)) // 7.5

	// Test case 4
	fmt.Println(angleClock(4, 50)) // 155
}

// Time: O(1)
// Space: O(1)
func angleClock(hour int, minutes int) float64 {
	// Minute hand: 360 degrees in 60 minutes = 6 degrees per minute
	minAngle := float64(minutes) * 6.0

	// Hour hand: 360 degrees in 12 hours = 30 degrees per hour
	// Plus 0.5 degrees per minute (30 degrees / 60 minutes)
	hourAngle := float64(hour%12)*30.0 + float64(minutes)*0.5

	diff := math.Abs(hourAngle - minAngle)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}
```
