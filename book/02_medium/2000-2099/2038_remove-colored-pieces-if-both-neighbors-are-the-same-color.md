# 2038 — Remove Colored Pieces If Both Neighbors Are The Same Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func winnerOfGame(colors string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2038: Remove Colored Pieces if Both Neighbors are the Same Color
// https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func winnerOfGame(colors string) bool {
	aMoves := 0
	bMoves := 0
	count := 1

	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[i-1] {
			count++
		} else {
			if colors[i-1] == 'A' && count >= 3 {
				aMoves += count - 2
			} else if colors[i-1] == 'B' && count >= 3 {
				bMoves += count - 2
			}
			count = 1
		}
	}
	if colors[len(colors)-1] == 'A' && count >= 3 {
		aMoves += count - 2
	} else if colors[len(colors)-1] == 'B' && count >= 3 {
		bMoves += count - 2
	}

	return aMoves > bMoves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", winnerOfGame("AAABABB"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", winnerOfGame("AA"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", winnerOfGame("ABBBBBBBAAA"))
	// Expected: false
}
```
