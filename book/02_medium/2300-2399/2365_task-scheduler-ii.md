# 2365 — Task Scheduler Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func taskSchedulerII(tasks []int, space int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2365: Task Scheduler II
// https://leetcode.com/problems/task-scheduler-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Track last completion day for each task type. If within cooldown, advance day.

import "fmt"

func main() {
	fmt.Println(taskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3)) // 9
	fmt.Println(taskSchedulerII([]int{5, 8, 8, 5}, 2))       // 6
}

func taskSchedulerII(tasks []int, space int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	last := make(map[int]int64)
	var day int64 = 0
	for _, t := range tasks {
		day++
		if prev, ok := last[t]; ok && day-prev <= int64(space) {
			day = prev + int64(space) + 1
		}
		last[t] = day
	}
	return day
}
```
