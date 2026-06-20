# 2526 — Find Consecutive Integers From A Data Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(value int, k int) DataStream
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per call  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2526: Find Consecutive Integers from a Data Stream
// https://leetcode.com/problems/find-consecutive-integers-from-a-data-stream/
// Difficulty: Medium
// Time: O(1) per call | Space: O(1)
// Track count of consecutive `value` seen.

import "fmt"

type DataStream struct {
	value, k, count int
}

func main() {
	ds := Constructor(4, 3)
	fmt.Println(ds.Consec(4)) // false
	fmt.Println(ds.Consec(4)) // false
	fmt.Println(ds.Consec(4)) // true
	fmt.Println(ds.Consec(3)) // false
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k}
}

func (ds *DataStream) Consec(num int) bool {
	if num == ds.value {
		ds.count++
	} else {
		ds.count = 0
	}
	return ds.count >= ds.k
}
```
