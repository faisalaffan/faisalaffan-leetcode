# 1185 — Day Of The Week

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func dayOfTheWeek(day, month, year int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1185: Day of the Week
// https://leetcode.com/problems/day-of-the-week/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019)) // "Saturday"
	fmt.Println(dayOfTheWeek(18, 7, 1999)) // "Sunday"
	fmt.Println(dayOfTheWeek(15, 8, 1993)) // "Sunday"
}

// LeetCode submission: dayOfTheWeek
func dayOfTheWeek(day, month, year int) string {
	// Tomohiko Sakamoto's algorithm
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if month < 3 {
		year--
	}
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	idx := (year + year/4 - year/100 + year/400 + t[month-1] + day) % 7
	return days[idx]
}
```
