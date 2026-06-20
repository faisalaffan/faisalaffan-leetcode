# 2644 — Find The Maximum Divisibility Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheMaximumDivisibilityScore(nums []int, divisors []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2644: Find the Maximum Divisibility Score
// https://leetcode.com/problems/find-the-maximum-divisibility-score/
// Difficulty: Easy
// Time: O(|nums| * |divisors|) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumDivisibilityScore([]int{2, 3, 4, 5, 6}, []int{2, 3, 4}))
	fmt.Println(FindTheMaximumDivisibilityScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}))
}

func FindTheMaximumDivisibilityScore(nums []int, divisors []int) int {
	ans := divisors[0]
	maxScore := 0
	for _, d := range divisors {
		score := 0
		for _, n := range nums {
			if n%d == 0 {
				score++
			}
		}
		if score > maxScore || (score == maxScore && d < ans) {
			maxScore = score
			ans = d
		}
	}
	return ans
}
```
