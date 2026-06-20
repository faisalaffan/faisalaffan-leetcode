# 1904 — The Number Of Full Rounds You Have Played

## Deskripsi

**Soal:** [1904. The Number Of Full Rounds You Have Played](https://leetcode.com/problems/the-number-of-full-rounds-you-have-played/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1904: The Number of Full Rounds You Have Played
// https://leetcode.com/problems/the-number-of-full-rounds-you-have-played/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfRounds("09:31", "10:14"))
	fmt.Println(NumberOfRounds("21:30", "03:00"))
}

// Time: O(1), Space: O(1)
func NumberOfRounds(loginTime string, logoutTime string) int {
	t1 := toMinutes(loginTime)
	t2 := toMinutes(logoutTime)

	if t1 > t2 {
		t2 += 24 * 60
	}

	// First full round starts at ceil(t1/15)*15
	start := ((t1 + 14) / 15) * 15
	// Last full round ends at floor(t2/15)*15
	end := (t2 / 15) * 15

	if end < start {
		return 0
	}
	return (end - start) / 15
}

func toMinutes(time string) int {
	hours := int(time[0]-'0')*10 + int(time[1]-'0')
	mins := int(time[3]-'0')*10 + int(time[4]-'0')
	return hours*60 + mins
}
```
