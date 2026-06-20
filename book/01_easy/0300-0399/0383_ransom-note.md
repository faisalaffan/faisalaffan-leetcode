# 0383 — Ransom Note

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RansomNote(ransomNote, magazine string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #383: Ransom Note
// https://leetcode.com/problems/ransom-note/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func RansomNote(ransomNote, magazine string) bool {
	count := [26]int{}
	for _, c := range magazine {
		count[c-'a']++
	}
	for _, c := range ransomNote {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(RansomNote("a", "b"))
	fmt.Println(RansomNote("aa", "ab"))
	fmt.Println(RansomNote("aa", "aab"))
}
```
