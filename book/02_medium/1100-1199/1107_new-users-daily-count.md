# 1107 — New Users Daily Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func newUsersDailyCount(traffic [][3]int) map[int]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = len(traffic)  |  **Ruang:** O(m) where m = unique users

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1107: New Users Daily Count
// https://leetcode.com/problems/new-users-daily-count/
// Difficulty: Medium
//
// Approach: Track first login date per user, count by date
// Time: O(n) where n = len(traffic)
// Space: O(m) where m = unique users

import "fmt"

func main() {
	// traffic: (user_id, activity date, is_login)
	traffic := [][3]int{
		{1, 1, 1},  // user 1 login on day 1
		{2, 1, 1},  // user 2 login on day 1
		{3, 2, 1},  // user 3 login on day 2
		{1, 3, 0},  // user 1 logout on day 3
		{2, 3, 0},  // user 2 logout on day 3
		{4, 3, 1},  // user 4 login on day 3
	}
	fmt.Println(newUsersDailyCount(traffic))
}

func newUsersDailyCount(traffic [][3]int) map[int]int {
  // HashMap: O(1) lookup
	firstLogin := make(map[int]int) // userID -> first login date
	for _, t := range traffic {
		userID, date, isLogin := t[0], t[1], t[2]
		if isLogin == 1 {
			if _, exists := firstLogin[userID]; !exists || date < firstLogin[userID] {
				firstLogin[userID] = date
			}
		}
	}

  // HashMap: O(1) lookup
	result := make(map[int]int)
	for _, date := range firstLogin {
		result[date]++
	}

	return result
}
```
