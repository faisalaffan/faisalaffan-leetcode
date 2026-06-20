# 3586 — Find Covid Recovery Patients

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindCovidRecoveryPatients(records []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3586: Find COVID Recovery Patients
// https://leetcode.com/problems/find-covid-recovery-patients/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	records := []int{1, 0, 1, 0, 0, 1}
	fmt.Println("Test 1:", FindCovidRecoveryPatients(records))
	// Test case 2
	records2 := []int{1, 1, 1}
	fmt.Println("Test 2:", FindCovidRecoveryPatients(records2))
	// Test case 3
	records3 := []int{0, 0, 0}
	fmt.Println("Test 3:", FindCovidRecoveryPatients(records3))
}

func FindCovidRecoveryPatients(records []int) int {
	// Count patients who have recovered (positive followed by negative)
	recovered := 0
	for i := 1; i < len(records); i++ {
		if records[i-1] == 1 && records[i] == 0 {
			recovered++
		}
	}
	return recovered
}
```
