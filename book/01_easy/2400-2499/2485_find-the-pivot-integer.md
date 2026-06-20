# 2485 — Find The Pivot Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindThePivotInteger(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2485: Find the Pivot Integer
// https://leetcode.com/problems/find-the-pivot-integer/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindThePivotInteger(8)) // 6
	fmt.Println(FindThePivotInteger(1)) // 1
	fmt.Println(FindThePivotInteger(4)) // -1
}

func FindThePivotInteger(n int) int {
	total := n * (n + 1) / 2
	sum := 0
	for x := 1; x <= n; x++ {
		sum += x
		if sum == total-sum+x {
			return x
		}
	}
	return -1
}
```
