# 2698 — Find The Punishment Number Of An Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func punishmentNumber(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * 2^len(num))  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2698: Find the Punishment Number of an Integer
// https://leetcode.com/problems/find-the-punishment-number-of-an-integer/
// Difficulty: Medium
// Time: O(n * 2^len(num)) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func punishmentNumber(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		sq := i * i
		sqStr := strconv.Itoa(sq)
		if canPartition(sqStr, i) {
			total += sq
		}
	}
	return total
}

func canPartition(s string, target int) bool {
	if len(s) == 0 {
		return target == 0
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		prefix, _ := strconv.Atoi(s[:i+1])
		if prefix > target {
			break
		}
		if canPartition(s[i+1:], target-prefix) {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", punishmentNumber(10))
	// Expected: 182 (1+81+100)

	// Test case 2
	fmt.Println("Test 2:", punishmentNumber(37))
	// Expected: 1478

	// Test case 3
	fmt.Println("Test 3:", punishmentNumber(1))
	// Expected: 1
}
```
