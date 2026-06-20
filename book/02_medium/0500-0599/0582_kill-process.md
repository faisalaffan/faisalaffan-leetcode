# 0582 — Kill Process

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KillProcess(pid []int, ppid []int, kill int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #582: Kill Process
// https://leetcode.com/problems/kill-process/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	fmt.Println(KillProcess(pid, ppid, kill))
}

func KillProcess(pid []int, ppid []int, kill int) []int {
	// Build adjacency list: parent -> children
  // Membuat map (HashMap) — pencarian O(1)
	children := make(map[int][]int)
	for i, p := range ppid {
		children[p] = append(children[p], pid[i])
	}

	// BFS/DFS to find all processes to kill
	result := []int{}
	queue := []int{kill}
	for len(queue) > 0 {
		process := queue[0]
		queue = queue[1:]
		result = append(result, process)
		queue = append(queue, children[process]...)
	}

	return result
}
```
