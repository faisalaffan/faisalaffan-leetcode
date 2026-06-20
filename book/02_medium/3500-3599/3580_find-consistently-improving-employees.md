# 3580 — Find Consistently Improving Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindConsistentlyImprovingEmployees(scores []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
