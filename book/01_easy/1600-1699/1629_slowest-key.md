# 1629 — Slowest Key

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SlowestKey(releaseTimes []int, keysPressed string) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1629: Slowest Key
// https://leetcode.com/problems/slowest-key/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SlowestKey(releaseTimes []int, keysPressed string) byte {
	maxDuration := releaseTimes[0]
	result := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > maxDuration || (duration == maxDuration && keysPressed[i] > result) {
			maxDuration = duration
			result = keysPressed[i]
		}
	}
	return result
}

func main() {
	fmt.Printf("%c\n", SlowestKey([]int{9, 29, 49, 50}, "cbcd"))
	fmt.Printf("%c\n", SlowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
}
```
