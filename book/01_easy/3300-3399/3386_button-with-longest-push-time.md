# 3386 — Button With Longest Push Time

## Deskripsi

**Soal:** [3386. Button With Longest Push Time](https://leetcode.com/problems/button-with-longest-push-time/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3386: Button with Longest Push Time
// https://leetcode.com/problems/button-with-longest-push-time/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ButtonWithLongestPushTime([][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}}))
	fmt.Println(ButtonWithLongestPushTime([][]int{{10, 5}, {1, 7}}))
}

// ButtonWithLongestPushTime returns the button index with the longest duration between consecutive events.
// Each event is [button_index, timestamp].
// Time: O(n). Space: O(1).
func ButtonWithLongestPushTime(events [][]int) int {
	maxDuration := 0
	buttonIndex := events[0][0]
	prevTime := events[0][1]

	for i := 1; i < len(events); i++ {
		duration := events[i][1] - prevTime
		if duration > maxDuration || (duration == maxDuration && events[i][0] < buttonIndex) {
			maxDuration = duration
			buttonIndex = events[i][0]
		}
		prevTime = events[i][1]
	}
	return buttonIndex
}
```
