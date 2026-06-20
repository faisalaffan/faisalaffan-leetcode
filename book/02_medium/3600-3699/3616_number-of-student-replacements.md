# 3616 — Number Of Student Replacements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfStudentReplacements(grades []int, threshold int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3616: Number of Student Replacements
// https://leetcode.com/problems/number-of-student-replacements/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grades := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 1:", NumberOfStudentReplacements(grades, 3))
	// Test case 2
	grades2 := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 2:", NumberOfStudentReplacements(grades2, 3))
	// Test case 3
	grades3 := []int{1, 1, 1}
	fmt.Println("Test 3:", NumberOfStudentReplacements(grades3, 2))
}

func NumberOfStudentReplacements(grades []int, threshold int) int {
	replacements := 0
	for _, g := range grades {
		if g < threshold {
			replacements++
		}
	}
	return replacements
}
```
