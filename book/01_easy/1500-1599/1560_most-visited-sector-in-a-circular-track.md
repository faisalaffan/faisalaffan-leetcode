# 1560 — Most Visited Sector In A Circular Track

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostVisited(n int, rounds []int) []int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1560: Most Visited Sector in a Circular Track
// https://leetcode.com/problems/most-visited-sector-in-a-circular-track/
// Difficulty: Easy
//
// LeetCode submission: func mostVisited(n int, rounds []int) []int

import "fmt"

func main() {
	fmt.Println(MostVisitedSectorInACircularTrack(4, []int{1, 3, 1, 2})) // [1 2]
	fmt.Println(MostVisitedSectorInACircularTrack(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2})) // [2]
	fmt.Println(MostVisitedSectorInACircularTrack(7, []int{1, 3, 5, 7})) // [1 2 3 4 5 6 7]
}

// Time: O(n), Space: O(n)
func MostVisitedSectorInACircularTrack(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]
	if start <= end {
  // Alokasi slice integer
		res := make([]int, end-start+1)
  // Range loop: iterasi dengan indeks + nilai
		for i := range res {
			res[i] = start + i
		}
		return res
	}
  // Alokasi slice integer
	res := make([]int, 0, n-end+start)
	for i := 1; i <= end; i++ {
		res = append(res, i)
	}
	for i := start; i <= n; i++ {
		res = append(res, i)
	}
	return res
}
```
