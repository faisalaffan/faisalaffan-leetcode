# 1107 — New Users Daily Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func newUsersDailyCount(traffic [][3]int) map[int]int
```

> **💡 Hint:** Track first login date per user, count by date

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) where n = len(traffic)  
**Kompleksitas Ruang:** O(m) where m = unique users

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	firstLogin := make(map[int]int) // userID -> first login date
	for _, t := range traffic {
		userID, date, isLogin := t[0], t[1], t[2]
		if isLogin == 1 {
			if _, exists := firstLogin[userID]; !exists || date < firstLogin[userID] {
				firstLogin[userID] = date
			}
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[int]int)
	for _, date := range firstLogin {
		result[date]++
	}

	return result
}
```
