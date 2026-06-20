# 3179 — Find The N Th Value After K Seconds

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func valueAfterKSeconds(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3179: Find the N-th Value After K Seconds
// https://leetcode.com/problems/find-the-n-th-value-after-k-seconds/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func valueAfterKSeconds(n int, k int) int {
	const mod = 1_000_000_007
  // Alokasi slice integer
	arr := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range arr {
		arr[i] = 1
	}

	for s := 0; s < k; s++ {
		for i := 1; i < n; i++ {
			arr[i] = (arr[i] + arr[i-1]) % mod
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(valueAfterKSeconds(4, 5)) // Expected: 56
	fmt.Println(valueAfterKSeconds(5, 3)) // Expected: 35
}
```
