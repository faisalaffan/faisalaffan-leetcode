# 2073 — Time Needed To Buy Tickets

## Deskripsi

**Soal:** [2073. Time Needed To Buy Tickets](https://leetcode.com/problems/time-needed-to-buy-tickets/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2073: Time Needed to Buy Tickets
// https://leetcode.com/problems/time-needed-to-buy-tickets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TimeNeededToBuyTickets([]int{2, 3, 2}, 2)) // 6
	fmt.Println(TimeNeededToBuyTickets([]int{5, 1, 1, 1}, 0)) // 8
}

// Time: O(n), Space: O(1)
func TimeNeededToBuyTickets(tickets []int, k int) int {
	time := 0
	for i, t := range tickets {
		if i <= k {
			if t <= tickets[k] {
				time += t
			} else {
				time += tickets[k]
			}
		} else {
			if t < tickets[k] {
				time += t
			} else {
				time += tickets[k] - 1
			}
		}
	}
	return time
}
```
