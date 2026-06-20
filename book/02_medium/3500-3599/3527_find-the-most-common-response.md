# 3527 — Find The Most Common Response

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheMostCommonResponse(responses []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3527: Find the Most Common Response
// https://leetcode.com/problems/find-the-most-common-response/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindTheMostCommonResponse([]string{"A", "B", "A", "C", "B", "A"}))
	// Test case 2
	fmt.Println("Test 2:", FindTheMostCommonResponse([]string{"X", "Y", "Z"}))
	// Test case 3
	fmt.Println("Test 3:", FindTheMostCommonResponse([]string{"M", "M", "M"}))
}

func FindTheMostCommonResponse(responses []string) string {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	maxFreq := 0
	mostCommon := ""
	for _, r := range responses {
		freq[r]++
		if freq[r] > maxFreq || (freq[r] == maxFreq && (mostCommon == "" || r < mostCommon)) {
			maxFreq = freq[r]
			mostCommon = r
		}
	}
	return mostCommon
}
```
