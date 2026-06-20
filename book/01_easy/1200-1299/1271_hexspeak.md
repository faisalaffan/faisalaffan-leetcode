# 1271 — Hexspeak

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func toHexspeak(num string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1271: Hexspeak
// https://leetcode.com/problems/hexspeak/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(toHexspeak("257"))  // "IOI"
	fmt.Println(toHexspeak("3"))    // "ERROR"
	fmt.Println(toHexspeak("619"))  // "ERROR" (619=26B, B not allowed)
}

// LeetCode submission: toHexspeak
func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	hex := strconv.FormatInt(int64(n), 16)
	replacer := map[byte]byte{
		'0': 'O',
		'1': 'I',
	}
	ans := make([]byte, len(hex))
  // Range loop: iterasi dengan indeks + nilai
	for i := range hex {
		if r, ok := replacer[hex[i]]; ok {
			ans[i] = r
		} else if hex[i] >= 'a' && hex[i] <= 'f' {
			ans[i] = hex[i] - 'a' + 'A'
		} else {
			return "ERROR"
		}
	}
	return string(ans)
}
```
