# 0849 — Maximize Distance To Closest Person

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeDistanceToClosestPerson(seats []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #849: Maximize Distance to Closest Person
// https://leetcode.com/problems/maximize-distance-to-closest-person/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximizeDistanceToClosestPerson([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println(MaximizeDistanceToClosestPerson([]int{1, 0, 0, 0}))
	fmt.Println(MaximizeDistanceToClosestPerson([]int{0, 1}))
}

// Time: O(n) | Space: O(1)
func MaximizeDistanceToClosestPerson(seats []int) int {
	n := len(seats)
	ans := 0
	lastPerson := -1

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if lastPerson == -1 {
				ans = i
			} else {
				dist := (i - lastPerson) / 2
				if dist > ans {
					ans = dist
				}
			}
			lastPerson = i
		}
	}

	// Check distance from last person to the end
	if seats[n-1] == 0 {
		dist := n - 1 - lastPerson
		if dist > ans {
			ans = dist
		}
	}

	return ans
}
```
