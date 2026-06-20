# 3172 — Second Day Verification

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func SecondDayVerification(actions [][]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	submitted := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
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
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```
