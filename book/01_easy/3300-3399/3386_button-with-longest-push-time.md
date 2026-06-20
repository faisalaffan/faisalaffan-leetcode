# 3386 — Button With Longest Push Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ButtonWithLongestPushTime(events [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3386: Button with Longest Push Time
// https://leetcode.com/problems/button-with-longest-push-time/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ButtonWithLongestPushTime([][]int{{1, 2}, {2, 5}, {3, 9}, {1, 15}}))
	fmt.Println(ButtonWithLongestPushTime([][]int{{10, 5}, {1, 7}}))
}

// ButtonWithLongestPushTime returns the button index with the longest duration between consecutive events.
// Each event is [button_index, timestamp].
// Time: O(n). Space: O(1).
func ButtonWithLongestPushTime(events [][]int) int {
	maxDuration := 0
	buttonIndex := events[0][0]
	prevTime := events[0][1]

	for i := 1; i < len(events); i++ {
		duration := events[i][1] - prevTime
		if duration > maxDuration || (duration == maxDuration && events[i][0] < buttonIndex) {
			maxDuration = duration
			buttonIndex = events[i][0]
		}
		prevTime = events[i][1]
	}
	return buttonIndex
}
```
