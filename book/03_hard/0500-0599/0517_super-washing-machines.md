# 0517 — Super Washing Machines

## Deskripsi

**Soal:** [0517. Super Washing Machines](https://leetcode.com/problems/super-washing-machines/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #517: Super Washing Machines
// https://leetcode.com/problems/super-washing-machines/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println(findMinMoves([]int{1, 0, 5})) // Expected: 3
}

func findMinMoves(machines []int) int {
	n := len(machines)
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%n != 0 {
		return -1
	}
	target := sum / n

	ans := 0
	balance := 0
	for _, v := range machines {
		balance += v - target
		if balance > ans {
			ans = balance
		}
		if balance < -ans {
			ans = -balance
		}
		// A machine may need to receive from both sides simultaneously
		if v-target > ans {
			ans = v - target
		}
	}
	return ans
}
```
