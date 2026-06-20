# 2028 — Find Missing Observations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func missingRolls(rolls []int, mean int, n int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2028: Find Missing Observations
// https://leetcode.com/problems/find-missing-observations/
// Difficulty: Medium
// Time: O(n + m) | Space: O(m)

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	totalSum := mean * (m + n)
	knownSum := 0
	for _, v := range rolls {
		knownSum += v
	}
	missingSum := totalSum - knownSum

	if missingSum < n || missingSum > 6*n {
		return []int{}
	}

  // Alokasi slice integer
	result := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		val := missingSum / (n - i)
		if val > 6 {
			val = 6
		}
		result[i] = val
		missingSum -= val
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", missingRolls([]int{3, 2, 4, 3}, 4, 2))
	// Expected: [6,6]

	// Test case 2
	fmt.Println("Test 2:", missingRolls([]int{1, 5, 6}, 3, 4))
	// Expected: [2,3,2,2]

	// Test case 3
	fmt.Println("Test 3:", missingRolls([]int{1, 2, 3, 4}, 6, 4))
	// Expected: [] (impossible)
}
```
