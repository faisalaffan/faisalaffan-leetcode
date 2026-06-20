# 2515 — Shortest Distance To Target String In A Circular Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ShortestDistanceToTargetStringInACircularArray(words []string, target string, startIndex int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2515: Shortest Distance to Target String in a Circular Array
// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(ShortestDistanceToTargetStringInACircularArray([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1)) // 1
	fmt.Println(ShortestDistanceToTargetStringInACircularArray([]string{"a", "b", "leetcode"}, "leetcode", 0))                // 1
}

func ShortestDistanceToTargetStringInACircularArray(words []string, target string, startIndex int) int {
	n := len(words)
	minDist := n

	for i, w := range words {
		if w == target {
			dist := abs(i - startIndex)
			if dist > n-dist {
				dist = n - dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
