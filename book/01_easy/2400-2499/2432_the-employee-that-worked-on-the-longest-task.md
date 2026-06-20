# 2432 — The Employee That Worked On The Longest Task

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TheEmployeeThatWorkedOnTheLongestTask(n int, logs [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2432: The Employee That Worked on the Longest Task
// https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(TheEmployeeThatWorkedOnTheLongestTask(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))   // 1
	fmt.Println(TheEmployeeThatWorkedOnTheLongestTask(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}})) // 3
}

func TheEmployeeThatWorkedOnTheLongestTask(n int, logs [][]int) int {
	bestID := logs[0][0]
	bestTime := logs[0][1]
	prevEnd := logs[0][1]

	for i := 1; i < len(logs); i++ {
		id := logs[i][0]
		start := prevEnd
		end := logs[i][1]
		duration := end - start
		if duration > bestTime || (duration == bestTime && id < bestID) {
			bestID = id
			bestTime = duration
		}
		prevEnd = end
	}
	return bestID
}
```
