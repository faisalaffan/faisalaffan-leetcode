# 3582 — Generate Tag For Video Caption

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GenerateTagForVideoCaption(caption string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3582: Generate Tag for Video Caption
// https://leetcode.com/problems/generate-tag-for-video-caption/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(GenerateTagForVideoCaption("Leetcode daily streak achieved"))
	fmt.Println(GenerateTagForVideoCaption("can I Go There"))
	fmt.Println(GenerateTagForVideoCaption("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
}

// Time: O(n)
// Space: O(n)
func GenerateTagForVideoCaption(caption string) string {
	words := strings.Fields(caption)
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
		} else {
			runes := []rune(w)
			for j, r := range runes {
				if j == 0 {
					runes[j] = unicode.ToUpper(r)
				} else {
					runes[j] = unicode.ToLower(r)
				}
			}
			words[i] = string(runes)
		}
	}

	result := "#" + strings.Join(words, "")
	if len(result) > 100 {
		result = result[:100]
	}
	return result
}
```
