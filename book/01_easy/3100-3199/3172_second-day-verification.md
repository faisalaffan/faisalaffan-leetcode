# 3172 — Second Day Verification

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func SecondDayVerification(actions [][]string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	submitted := make(map[int]int)
  // HashMap: O(1) lookup
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
  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```
