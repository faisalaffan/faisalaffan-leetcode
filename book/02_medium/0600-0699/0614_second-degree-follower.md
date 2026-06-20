# 0614 — Second Degree Follower

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SecondDegreeFollowers(follows [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #614: Second Degree Follower
// https://leetcode.com/problems/second-degree-follower/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Follow relationships: {follower, followee}
	follows := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 1},
		{3, 2},
		{4, 5},
	}
	fmt.Println(SecondDegreeFollowers(follows))
}

func SecondDegreeFollowers(follows [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	followers := make(map[int]map[int]bool)

	for _, f := range follows {
		follower, followee := f[0], f[1]
		if followers[followee] == nil {
			followers[followee] = make(map[int]bool)
		}
		followers[followee][follower] = true
	}

	result := []int{}
	for userID, fMap := range followers {
		if len(fMap) >= 2 {
			result = append(result, userID)
		}
	}

	return result
}
```
