# 2951 — Find The Peaks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindThePeaks(mountain []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2951: Find the Peaks
// https://leetcode.com/problems/find-the-peaks/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPeaks
	fmt.Println(FindThePeaks([]int{2, 4, 4}))    // []
	fmt.Println(FindThePeaks([]int{1, 4, 3, 8, 5})) // [1, 3]
}

// Time: O(n) | Space: O(1) excluding output
// LeetCode submission name: findPeaks
func FindThePeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}
```
