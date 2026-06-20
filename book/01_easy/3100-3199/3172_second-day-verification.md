# 3172 — Second Day Verification

## Deskripsi

**Soal:** [3172. Second Day Verification](https://leetcode.com/problems/second-day-verification/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3172: Second Day Verification
// https://leetcode.com/problems/second-day-verification/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: determine which user IDs had their verification
// completed on the second day.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: secondDayVerify
	// Input: [user_id, day, action] where action is 'submitted' or 'verified'
	actions := [][]string{
		{"1", "1", "submitted"},
		{"1", "2", "verified"},
		{"2", "1", "submitted"},
		{"3", "1", "submitted"},
		{"3", "3", "verified"},
	}
	fmt.Println(SecondDayVerification(actions))
	// [1]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: secondDayVerify
func SecondDayVerification(actions [][]string) []int {
  // Membuat map untuk pencarian O(1): key → value
	submitted := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	verified := make(map[int]int)
	for _, row := range actions {
		id := parseInt(row[0])
		day := parseInt(row[1])
		action := row[2]
		if action == "submitted" {
			submitted[id] = day
		} else if action == "verified" {
			verified[id] = day
		}
	}
	result := []int{}
	for id, submitDay := range submitted {
		if verifyDay, ok := verified[id]; ok && verifyDay == submitDay+1 {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```
