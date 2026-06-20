# 1103 — Distribute Candies To People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func distributeCandies(candies int, numPeople int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(candies))  
**Kompleksitas Ruang:** O(numPeople)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1103: Distribute Candies to People
// https://leetcode.com/problems/distribute-candies-to-people/
// Difficulty: Easy
// Time: O(sqrt(candies)) | Space: O(numPeople)

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))  // [1,2,3,1]
	fmt.Println(distributeCandies(10, 3)) // [5,2,3]
}

// LeetCode submission: distributeCandies
func distributeCandies(candies int, numPeople int) []int {
  // Alokasi slice integer
	ans := make([]int, numPeople)
	give := 1
	for candies > 0 {
		for i := 0; i < numPeople && candies > 0; i++ {
			if give <= candies {
				ans[i] += give
				candies -= give
			} else {
				ans[i] += candies
				candies = 0
			}
			give++
		}
	}
	return ans
}
```
