# 3580 — Find Consistently Improving Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindConsistentlyImprovingEmployees(scores []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3580: Find Consistently Improving Employees
// https://leetcode.com/problems/find-consistently-improving-employees/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	scores := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 1:", FindConsistentlyImprovingEmployees(scores))
	// Test case 2
	scores2 := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 2:", FindConsistentlyImprovingEmployees(scores2))
	// Test case 3
	scores3 := []int{1, 3, 2, 4, 5}
	fmt.Println("Test 3:", FindConsistentlyImprovingEmployees(scores3))
}

func FindConsistentlyImprovingEmployees(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	count := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			count++
		}
	}
	return count
}
```
