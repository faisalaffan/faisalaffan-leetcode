# 3439 — Reschedule Meetings For Maximum Free Time I

## Deskripsi

**Soal:** [3439. Reschedule Meetings For Maximum Free Time I](https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int`

## Solusi Go

```go
package main

// LeetCode #3439: Reschedule Meetings for Maximum Free Time I
// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
  // Membuat slice untuk menyimpan hasil
	gaps := make([]int, 0, n+1)
	gaps = append(gaps, startTime[0])
	for i := 1; i < n; i++ {
		gaps = append(gaps, startTime[i]-endTime[i-1])
	}
	gaps = append(gaps, eventTime-endTime[n-1])

	window := 0
	for i := 0; i < k+1 && i < len(gaps); i++ {
		window += gaps[i]
	}
	ans := window
	for i := k + 1; i < len(gaps); i++ {
		window += gaps[i] - gaps[i-(k+1)]
		if window > ans {
			ans = window
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFreeTime(10, 1, []int{0, 3, 7, 9}, []int{1, 4, 8, 10})) // 3
	fmt.Println(maxFreeTime(5, 2, []int{1, 3}, []int{2, 4})) // 2
}
```
