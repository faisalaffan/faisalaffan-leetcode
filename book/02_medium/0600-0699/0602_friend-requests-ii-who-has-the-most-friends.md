# 0602 — Friend Requests Ii Who Has The Most Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MostFriends(requests [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #602: Friend Requests II: Who Has the Most Friends
// https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Friend requests: {requester_id, accepter_id}
	requests := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	fmt.Println(MostFriends(requests))
}

func MostFriends(requests [][]int) int {
  // Membuat map (HashMap) — pencarian O(1)
	friendCount := make(map[int]int)
	for _, req := range requests {
		friendCount[req[0]]++
		friendCount[req[1]]++
	}

	maxCount := 0
	maxID := 0
	for id, count := range friendCount {
		if count > maxCount {
			maxCount = count
			maxID = id
		}
	}

	return maxID
}
```
