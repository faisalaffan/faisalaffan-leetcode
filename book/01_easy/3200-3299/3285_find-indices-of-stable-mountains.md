# 3285 — Find Indices Of Stable Mountains

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindIndicesOfStableMountains(height []int, threshold int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3285: Find Indices of Stable Mountains
// https://leetcode.com/problems/find-indices-of-stable-mountains/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIndicesOfStableMountains([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(FindIndicesOfStableMountains([]int{10, 1, 10, 1, 10}, 3))
}

// FindIndicesOfStableMountains returns indices of stable mountains (where the previous mountain's height > threshold).
// Time: O(n). Space: O(n).
func FindIndicesOfStableMountains(height []int, threshold int) []int {
	result := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			result = append(result, i)
		}
	}
	return result
}
```
