# 2682 — Find The Losers Of The Circular Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLosersOfTheCircularGame(n int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2682: Find the Losers of the Circular Game
// https://leetcode.com/problems/find-the-losers-of-the-circular-game/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheLosersOfTheCircularGame(5, 2))
	fmt.Println(FindTheLosersOfTheCircularGame(4, 4))
}

func FindTheLosersOfTheCircularGame(n int, k int) []int {
	visited := make([]bool, n)
	i := 0
	step := k
	for !visited[i] {
		visited[i] = true
		i = (i + step) % n
		step += k
	}

	result := []int{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			result = append(result, i+1) // 1-indexed
		}
	}
	return result
}
```
