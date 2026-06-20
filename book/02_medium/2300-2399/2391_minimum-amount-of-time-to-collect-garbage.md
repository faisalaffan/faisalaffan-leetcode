# 2391 — Minimum Amount Of Time To Collect Garbage

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func garbageCollection(garbage []string, travel []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(1) where k = types (3)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2391: Minimum Amount of Time to Collect Garbage
// https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
// Difficulty: Medium
// Time: O(n * k) | Space: O(1) where k = types (3)
// Sum collection (1 min per unit) + travel to last house containing each type.

import "fmt"

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3})) // 21
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))     // 37
}

func garbageCollection(garbage []string, travel []int) int {
  // Alokasi slice integer
	last := make([]int, 3) // 0=G, 1=P, 2=M
	total := 0
	for i, g := range garbage {
		total += len(g)
		for _, ch := range g {
			switch ch {
			case 'G':
				last[0] = i
			case 'P':
				last[1] = i
			case 'M':
				last[2] = i
			}
		}
	}

	// prefix sums for travel
  // Alokasi slice integer
	pref := make([]int, len(travel)+1)
	for i, t := range travel {
		pref[i+1] = pref[i] + t
	}

	for _, l := range last {
		total += pref[l]
	}
	return total
}
```
