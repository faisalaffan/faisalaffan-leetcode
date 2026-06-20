# 0596 — Classes With At Least 5 Students

## Deskripsi

**Soal:** [0596. Classes With At Least 5 Students](https://leetcode.com/problems/classes-with-at-least-5-students/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func ClassesWithAtLeastFiveStudents() string`

## Solusi Go

```go
package main

// LeetCode #596: Classes With at Least 5 Students
// https://leetcode.com/problems/classes-with-at-least-5-students/
// Difficulty: Easy

import "fmt"

func ClassesWithAtLeastFiveStudents() string {
	return "SELECT class FROM Courses GROUP BY class HAVING COUNT(student) >= 5"
}

func main() {
	fmt.Println(ClassesWithAtLeastFiveStudents())
}
```
