# 0299 — Bulls And Cows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getHint(secret string, guess string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #299: Bulls and Cows
// https://leetcode.com/problems/bulls-and-cows/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretCount := [10]int{}
	guessCount := [10]int{}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretCount[secret[i]-'0']++
			guessCount[guess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		if secretCount[i] < guessCount[i] {
			cows += secretCount[i]
		} else {
			cows += guessCount[i]
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("1", "0"))
}
```
