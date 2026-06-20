# 2239 — Find Closest Number To Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindClosestNumberToZero(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2239: Find Closest Number to Zero
// https://leetcode.com/problems/find-closest-number-to-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindClosestNumberToZero([]int{-4, -2, 1, 4, 8})) // 1
	fmt.Println(FindClosestNumberToZero([]int{2, -1, 1}))         // 1
}

// Time: O(n), Space: O(1)
func FindClosestNumberToZero(nums []int) int {
	closest := nums[0]
	for _, v := range nums[1:] {
		absV := v
		if absV < 0 {
			absV = -absV
		}
		absClosest := closest
		if absClosest < 0 {
			absClosest = -absClosest
		}
		if absV < absClosest || (absV == absClosest && v > closest) {
			closest = v
		}
	}
	return closest
}
```
