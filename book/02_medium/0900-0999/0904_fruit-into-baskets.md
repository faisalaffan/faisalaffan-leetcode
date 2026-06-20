# 0904 — Fruit Into Baskets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FruitIntoBaskets(fruits []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #904: Fruit Into Baskets
// https://leetcode.com/problems/fruit-into-baskets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FruitIntoBaskets([]int{1, 2, 1}))
	fmt.Println(FruitIntoBaskets([]int{0, 1, 2, 2}))
	fmt.Println(FruitIntoBaskets([]int{1, 2, 3, 2, 2}))
}

// Time: O(n) | Space: O(1)
func FruitIntoBaskets(fruits []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[int]int)
	left, ans := 0, 0

	for right, fruit := range fruits {
		cnt[fruit]++
		for len(cnt) > 2 {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				delete(cnt, fruits[left])
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}
```
