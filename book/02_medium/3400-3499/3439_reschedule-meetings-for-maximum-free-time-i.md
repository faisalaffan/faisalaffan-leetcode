# 3439 — Reschedule Meetings For Maximum Free Time I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3439: Reschedule Meetings for Maximum Free Time I
// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
  // Alokasi slice integer
	gaps := make([]int, 0, n+1)
	gaps = append(gaps, startTime[0])
	for i := 1; i < n; i++ {
		gaps = append(gaps, startTime[i]-endTime[i-1])
	}
	gaps = append(gaps, eventTime-endTime[n-1])

	window := 0
	for i := 0; i < k+1 && i < len(gaps); i++ {
		window += gaps[i]
	}
	ans := window
	for i := k + 1; i < len(gaps); i++ {
		window += gaps[i] - gaps[i-(k+1)]
		if window > ans {
			ans = window
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFreeTime(10, 1, []int{0, 3, 7, 9}, []int{1, 4, 8, 10})) // 3
	fmt.Println(maxFreeTime(5, 2, []int{1, 3}, []int{2, 4})) // 2
}
```
