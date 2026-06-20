# 1165 — Single Row Keyboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateTime(keyboard, word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1165: Single-Row Keyboard
// https://leetcode.com/problems/single-row-keyboard/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(calculateTime("abcdefghijklmnopqrstuvwxyz", "cba")) // 4
	fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")) // 73
}

// LeetCode submission: calculateTime
func calculateTime(keyboard, word string) int {
	pos := [26]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(keyboard); i++ {
		pos[keyboard[i]-'a'] = i
	}
	ans, cur := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		next := pos[word[i]-'a']
		if next > cur {
			ans += next - cur
		} else {
			ans += cur - next
		}
		cur = next
	}
	return ans
}
```
