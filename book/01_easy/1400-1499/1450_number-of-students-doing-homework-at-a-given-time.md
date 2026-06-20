# 1450 — Number Of Students Doing Homework At A Given Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func busyStudent(startTime []int, endTime []int, queryTime int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Range loop: iterasi dengan indeks + nilai
	for i := range startTime {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
```
