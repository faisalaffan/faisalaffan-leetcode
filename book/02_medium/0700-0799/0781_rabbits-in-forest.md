# 0781 — Rabbits In Forest

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numRabbits(answers []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #781: Rabbits in Forest
// https://leetcode.com/problems/rabbits-in-forest/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
}

func numRabbits(answers []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	for _, a := range answers {
		count[a]++
	}

	result := 0
	for k, v := range count {
		groupSize := k + 1
		groups := (v + groupSize - 1) / groupSize
		result += groups * groupSize
	}

	return result
}
```
