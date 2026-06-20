# 1450 — Number Of Students Doing Homework At A Given Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func busyStudent(startTime []int, endTime []int, queryTime int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1450: Number of Students Doing Homework at a Given Time
// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
// Difficulty: Easy
//
// LeetCode submission: func busyStudent(startTime []int, endTime []int, queryTime int) int

import "fmt"

func main() {
	fmt.Println(NumberOfStudentsDoingHomeworkAtAGivenTime([]int{1, 2, 3}, []int{3, 2, 7}, 4)) // 1
	fmt.Println(NumberOfStudentsDoingHomeworkAtAGivenTime([]int{4}, []int{4}, 4))             // 1
}

// Time: O(n), Space: O(1)
func NumberOfStudentsDoingHomeworkAtAGivenTime(startTime []int, endTime []int, queryTime int) int {
	count := 0
  // Range loop
	for i := range startTime {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
```
