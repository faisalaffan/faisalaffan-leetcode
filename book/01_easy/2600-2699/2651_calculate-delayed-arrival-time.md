# 2651 — Calculate Delayed Arrival Time

## Deskripsi

**Soal:** [2651. Calculate Delayed Arrival Time](https://leetcode.com/problems/calculate-delayed-arrival-time/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2651: Calculate Delayed Arrival Time
// https://leetcode.com/problems/calculate-delayed-arrival-time/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CalculateDelayedArrivalTime(15, 5))
	fmt.Println(CalculateDelayedArrivalTime(23, 10))
}

func CalculateDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
```
