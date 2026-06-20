# 3295 — Report Spam Message

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func reportSpam(message []string, bannedWords []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m) Space: O(m)  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3295: Report Spam Message
// https://leetcode.com/problems/report-spam-message/
// Difficulty: Medium
// Time: O(n + m) Space: O(m)

import "fmt"

func main() {
	fmt.Println(reportSpam([]string{"hello", "world", "leetcode"}, []string{"world", "hello"})) // true
	fmt.Println(reportSpam([]string{"hello", "programming", "fun"}, []string{"world", "hello"})) // false
	fmt.Println(reportSpam([]string{"a", "b", "c", "d"}, []string{"a", "b", "x"}))               // true
}

func reportSpam(message []string, bannedWords []string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	banned := make(map[string]struct{}, len(bannedWords))
	for _, w := range bannedWords {
		banned[w] = struct{}{}
	}
	count := 0
	for _, w := range message {
		if _, ok := banned[w]; ok {
			count++
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
```
