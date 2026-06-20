# 2432 — The Employee That Worked On The Longest Task

## Deskripsi

**Soal:** [2432. The Employee That Worked On The Longest Task](https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2432: The Employee That Worked on the Longest Task
// https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(TheEmployeeThatWorkedOnTheLongestTask(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))   // 1
	fmt.Println(TheEmployeeThatWorkedOnTheLongestTask(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}})) // 3
}

func TheEmployeeThatWorkedOnTheLongestTask(n int, logs [][]int) int {
	bestID := logs[0][0]
	bestTime := logs[0][1]
	prevEnd := logs[0][1]

	for i := 1; i < len(logs); i++ {
		id := logs[i][0]
		start := prevEnd
		end := logs[i][1]
		duration := end - start
		if duration > bestTime || (duration == bestTime && id < bestID) {
			bestID = id
			bestTime = duration
		}
		prevEnd = end
	}
	return bestID
}
```
