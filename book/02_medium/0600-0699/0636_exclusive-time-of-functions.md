# 0636 — Exclusive Time Of Functions

## Deskripsi

**Soal:** [0636. Exclusive Time Of Functions](https://leetcode.com/problems/exclusive-time-of-functions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #636: Exclusive Time of Functions
// https://leetcode.com/problems/exclusive-time-of-functions/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}))
}

func exclusiveTime(n int, logs []string) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		typ := parts[1]
		timestamp, _ := strconv.Atoi(parts[2])

		if typ == "start" {
			if len(stack) > 0 {
				result[stack[len(stack)-1]] += timestamp - prevTime
			}
			stack = append(stack, id)
			prevTime = timestamp
		} else {
			result[id] += timestamp - prevTime + 1
			stack = stack[:len(stack)-1]
			prevTime = timestamp + 1
		}
	}

	return result
}
```
