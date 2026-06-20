# 0621 — Task Scheduler

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LeastInterval(tasks []byte, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (26 letters)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #621: Task Scheduler
// https://leetcode.com/problems/task-scheduler/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (26 letters)

import "fmt"

func main() {
	fmt.Println(LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(LeastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}

func LeastInterval(tasks []byte, n int) int {
  // Alokasi slice integer
	counts := make([]int, 26)
	maxFreq := 0
	for _, t := range tasks {
		counts[t-'A']++
		if counts[t-'A'] > maxFreq {
			maxFreq = counts[t-'A']
		}
	}

	maxCount := 0
	for _, c := range counts {
		if c == maxFreq {
			maxCount++
		}
	}

	partLen := maxFreq - 1
	emptySlots := partLen * (n - (maxCount - 1))
	availableTasks := len(tasks) - maxFreq*maxCount
	idles := 0
	if emptySlots > availableTasks {
		idles = emptySlots - availableTasks
	}

	return len(tasks) + idles
}
```
