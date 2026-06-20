# 1894 — Find The Student That Will Replace The Chalk

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func ChalkReplacer(chalk []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1894: Find the Student that Will Replace the Chalk
// https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ChalkReplacer([]int{5, 1, 5}, 22))
	fmt.Println(ChalkReplacer([]int{3, 4, 1, 2}, 25))
	fmt.Println(ChalkReplacer([]int{5, 2, 3}, 9))
}

// Time: O(n), Space: O(1)
func ChalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, c := range chalk {
		sum += c
	}
	k %= sum

	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return 0
}
```
