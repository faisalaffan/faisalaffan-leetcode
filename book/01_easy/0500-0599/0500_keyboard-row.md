# 0500 — Keyboard Row

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KeyboardRow(words []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*k), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #500: Keyboard Row
// https://leetcode.com/problems/keyboard-row/
// Difficulty: Easy

import "fmt"

// Time: O(n*k), Space: O(n)
func KeyboardRow(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
  // Membuat map (HashMap) — pencarian O(1)
	rowMap := make(map[byte]int)
	for i, row := range rows {
		for j := 0; j < len(row); j++ {
			rowMap[row[j]] = i
		}
	}
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		targetRow := rowMap[toLower(word[0])]
		sameRow := true
		for i := 1; i < len(word); i++ {
			if rowMap[toLower(word[i])] != targetRow {
				sameRow = false
				break
			}
		}
		if sameRow {
			result = append(result, word)
		}
	}
	return result
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(KeyboardRow([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(KeyboardRow([]string{"omk"}))
	fmt.Println(KeyboardRow([]string{"adsdf", "sfd"}))
}
```
