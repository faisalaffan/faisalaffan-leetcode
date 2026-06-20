# 0484 — Find Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindPermutation(s string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #484: Find Permutation
// https://leetcode.com/problems/find-permutation/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindPermutation("I"))
	fmt.Println(FindPermutation("DI"))
}

func FindPermutation(s string) []int {
	n := len(s) + 1
  // Alokasi slice integer
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}

	// Reverse contiguous segments for each 'D'
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			j := i
			for j < len(s) && s[j] == 'D' {
				j++
			}
			// Reverse segment from i to j
			left, right := i, j
  // Two-pointer: gerakkan kiri atau kanan
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			i = j
		}
	}

	return result
}
```
