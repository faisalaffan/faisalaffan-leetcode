# 2703 — Return Length Of Arguments Passed

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ReturnLengthOfArgumentsPassed(args ...interface{}) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2703: Return Length of Arguments Passed
// https://leetcode.com/problems/return-length-of-arguments-passed/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns number of arguments.

import "fmt"

func main() {
	fmt.Println(ReturnLengthOfArgumentsPassed(1, 2, 3))
	fmt.Println(ReturnLengthOfArgumentsPassed("a", "b"))
}

func ReturnLengthOfArgumentsPassed(args ...interface{}) int {
	return len(args)
}
```
