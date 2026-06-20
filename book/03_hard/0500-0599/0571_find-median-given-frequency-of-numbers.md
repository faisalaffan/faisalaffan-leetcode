# 0571 — Find Median Given Frequency Of Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMedianGivenFrequencyOfNumbers(numbers []NumberFreq) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N log N) where N = number of distinct num values (sorting)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #571: Find Median Given Frequency of Numbers
// https://leetcode.com/problems/find-median-given-frequency-of-numbers/
// Difficulty: Hard [Paid]
//
// Given a Numbers table with (Num, Frequency), calculate the median of all numbers.
// Each Num appears Frequency times. The median is the middle value (or average of
// two middle values for even counts).

// NumberFreq represents a row in the Numbers table.
type NumberFreq struct {
	Num       int
	Frequency int
}

// findMedianGivenFrequencyOfNumbers computes the median from a frequency table.
// Time: O(N log N) where N = number of distinct num values (sorting)
// Space: O(N)
func findMedianGivenFrequencyOfNumbers(numbers []NumberFreq) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Sort by Num.
  // Custom sort dengan comparator
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Num < numbers[j].Num
	})

	// Compute total count.
	total := 0
	for _, nf := range numbers {
		total += nf.Frequency
	}

	// Find median position(s) - 1-indexed.
	// For total=odd, median is value at position (total+1)/2.
	// For total=even, median is average of values at positions total/2 and total/2+1.
	pos1 := (total + 1) / 2 // 1-indexed lower median position
	pos2 := total/2 + 1     // 1-indexed upper median position (same as pos1 for odd)

	// Scan cumulative frequencies to find values at pos1 and pos2.
	cumSum := 0
	val1, val2 := 0, 0
	found1, found2 := false, false

	for _, nf := range numbers {
		cumSum += nf.Frequency

		if !found1 && cumSum >= pos1 {
			val1 = nf.Num
			found1 = true
		}
		if !found2 && cumSum >= pos2 {
			val2 = nf.Num
			found2 = true
		}
		if found1 && found2 {
			break
		}
	}

	// For odd total, pos1 == pos2 so val1 == val2.
	// For even total, average of two middle values.
	return float64(val1+val2) / 2.0
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0571 Find Median Given Frequency of Numbers ===")

	// Example: Nums 0(7x), 1(1x), 2(3x), 3(1x).
	// Sorted: 0,0,0,0,0,0,0,1,2,2,2,3 (total=12)
	// Median at positions 6 and 7 -> 0 and 0 -> avg = 0
	numbers := []NumberFreq{
		{Num: 0, Frequency: 7},
		{Num: 1, Frequency: 1},
		{Num: 2, Frequency: 3},
		{Num: 3, Frequency: 1},
	}
	med := findMedianGivenFrequencyOfNumbers(numbers)
	fmt.Printf("Test 1 - Median = %.1f (expected 0.0)\n", med)

	// Odd total: 1(1x), 2(2x), 3(3x) -> 1,2,2,3,3,3 (total=6)
	// Even total: median positions 3 and 4 -> values 2 and 3 -> avg = 2.5
	numbers2 := []NumberFreq{
		{Num: 1, Frequency: 1},
		{Num: 2, Frequency: 2},
		{Num: 3, Frequency: 3},
	}
	med2 := findMedianGivenFrequencyOfNumbers(numbers2)
	fmt.Printf("Test 2 - Median = %.1f (expected 2.5)\n", med2)

	// Single element.
	numbers3 := []NumberFreq{
		{Num: 100, Frequency: 5},
	}
	med3 := findMedianGivenFrequencyOfNumbers(numbers3)
	fmt.Printf("Test 3 - Median = %.1f (expected 100.0)\n", med3)

	// Even count with two distinct middle values.
	numbers4 := []NumberFreq{
		{Num: 5, Frequency: 1},
		{Num: 10, Frequency: 1},
	}
	med4 := findMedianGivenFrequencyOfNumbers(numbers4)
	fmt.Printf("Test 4 - Median = %.1f (expected 7.5)\n", med4)

	// Large frequency.
	numbers5 := []NumberFreq{
		{Num: 1, Frequency: 100},
		{Num: 2, Frequency: 1},
		{Num: 3, Frequency: 100},
	}
	med5 := findMedianGivenFrequencyOfNumbers(numbers5)
	fmt.Printf("Test 5 - Median = %.1f (expected 2.0)\n", med5)

	// Empty.
	med6 := findMedianGivenFrequencyOfNumbers(nil)
	fmt.Printf("Test 6 - Empty = %.1f (expected 0.0)\n", med6)
}
```
