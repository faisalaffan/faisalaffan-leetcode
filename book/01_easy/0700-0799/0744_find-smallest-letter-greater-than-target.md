# 0744 — Find Smallest Letter Greater Than Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func nextGreatestLetter(letters []byte, target byte) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #744: Find Smallest Letter Greater Than Target
// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))) // 'c'
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))) // 'f'
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))) // 'x'
}

// nextGreatestLetter finds the smallest letter in letters that is greater than target.
// Time: O(log n). Space: O(1).
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l%len(letters)]
}
```
