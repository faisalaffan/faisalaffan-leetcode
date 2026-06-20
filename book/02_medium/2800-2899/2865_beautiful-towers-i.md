# 2865 — Beautiful Towers I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BeautifulTowersI(maxHeights []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2865: Beautiful Towers I
// https://leetcode.com/problems/beautiful-towers-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func BeautifulTowersI(maxHeights []int) int64 {
	n := len(maxHeights)
	var best int64

	for peak := 0; peak < n; peak++ {
		var total int64 = int64(maxHeights[peak])
		prev := maxHeights[peak]

		// Left side
		for i := peak - 1; i >= 0; i-- {
			h := maxHeights[i]
			if h > prev {
				h = prev
			}
			total += int64(h)
			prev = h
		}

		prev = maxHeights[peak]
		// Right side
		for i := peak + 1; i < n; i++ {
			h := maxHeights[i]
			if h > prev {
				h = prev
			}
			total += int64(h)
			prev = h
		}

		if total > best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(BeautifulTowersI([]int{5, 3, 4, 1, 1}))
	fmt.Println(BeautifulTowersI([]int{6, 5, 3, 9, 2, 7}))
}
```
