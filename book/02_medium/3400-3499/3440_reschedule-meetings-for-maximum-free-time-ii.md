# 3440 — Reschedule Meetings For Maximum Free Time Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxFreeTime3440(eventTime int, startTime []int, endTime []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3440: Reschedule Meetings for Maximum Free Time II
// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-ii/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func maxFreeTime3440(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
  // Alokasi slice integer
	v := make([]int, 0, n+1)
	last := 0
	for i := 0; i < n; i++ {
		v = append(v, startTime[i]-last)
		last = endTime[i]
	}
	if last <= eventTime {
		v = append(v, eventTime-last)
	}

	m := len(v)
  // Alokasi slice integer
	q := make([]int, m)
	q[m-1] = 0
	for i := m - 2; i >= 0; i-- {
		if v[i+1] > q[i+1] {
			q[i] = v[i+1]
		} else {
			q[i] = q[i+1]
		}
	}

	mx := 0
	ans := 0
	for i := 1; i < m; i++ {
		length := endTime[i-1] - startTime[i-1]
		t := v[i] + v[i-1]
		if length <= mx || length <= q[i] {
			if t+length > ans {
				ans = t + length
			}
		} else {
			if t > ans {
				ans = t
			}
		}
		if v[i-1] > mx {
			mx = v[i-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFreeTime3440(10, []int{0, 3, 7, 9}, []int{1, 4, 8, 10})) // 6
	fmt.Println(maxFreeTime3440(5, []int{1, 3}, []int{2, 4})) // 2
}
```
