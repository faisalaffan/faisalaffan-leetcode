# 3158 — Find The Xor Of Numbers Which Appear Twice

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheXorOfNumbersWhichAppearTwice(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3158: Find the XOR of Numbers Which Appear Twice
// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: duplicateNumbersXOR
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 2, 1})) // 3
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 3}))    // 0
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: duplicateNumbersXOR
func FindTheXorOfNumbersWhichAppearTwice(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	xor := 0
	for v, f := range freq {
		if f == 2 {
			xor ^= v
		}
	}
	return xor
}
```
