# 2802 — Find The K Th Lucky Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheKThLuckyNumber(k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log k)  
**Kompleksitas Ruang:** O(log k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2802: Find The K-th Lucky Number
// https://leetcode.com/problems/find-the-k-th-lucky-number/
// Difficulty: Medium [Paid]
// Time: O(log k) | Space: O(log k)

import "fmt"

func FindTheKThLuckyNumber(k int) string {
	// K-th lucky number: binary representation of k+1, then replace 0->4, 1->7
	// n = k + 1
	n := k + 1
	binary := fmt.Sprintf("%b", n)
	// Remove first '1' (it's the leading bit from the offset)
	result := make([]byte, len(binary)-1)
	for i := 1; i < len(binary); i++ {
		if binary[i] == '0' {
			result[i-1] = '4'
		} else {
			result[i-1] = '7'
		}
	}
	return string(result)
}

func main() {
	fmt.Println(FindTheKThLuckyNumber(1))
	fmt.Println(FindTheKThLuckyNumber(3))
	fmt.Println(FindTheKThLuckyNumber(5))
}
```
