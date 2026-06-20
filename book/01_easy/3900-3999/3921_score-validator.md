# 3921 — Score Validator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ScoreValidator(events []string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3921: Score Validator
// https://leetcode.com/problems/score-validator/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ScoreValidator([]string{"1", "4", "W", "6", "WD"}))
	fmt.Println(ScoreValidator([]string{"WD", "NB", "0", "4", "4"}))
	fmt.Println(ScoreValidator([]string{"W", "W", "W", "W", "W", "W", "W", "W", "W", "W", "W"}))
}

// Time: O(n)
// Space: O(1)
func ScoreValidator(events []string) []int {
	score, counter := 0, 0
	for _, e := range events {
		if counter == 10 {
			break
		}
		if e == "W" {
			counter++
		} else if e == "WD" || e == "NB" {
			score++
		} else {
			v, _ := strconv.Atoi(e)
			score += v
		}
	}
	return []int{score, counter}
}
```
