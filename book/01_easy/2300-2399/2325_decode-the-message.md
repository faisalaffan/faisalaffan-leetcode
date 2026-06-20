# 2325 — Decode The Message

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DecodeTheMessage(key string, message string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2325: Decode the Message
// https://leetcode.com/problems/decode-the-message/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DecodeTheMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")) // "this is a secret"
	fmt.Println(DecodeTheMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb")) // "the five boxing wizards jump quickly"
}

func DecodeTheMessage(key string, message string) string {
	mapping := make([]byte, 26)
	idx := byte(0)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && mapping[key[i]-'a'] == 0 {
			mapping[key[i]-'a'] = 'a' + idx
			idx++
		}
	}

	res := make([]byte, len(message))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			res[i] = ' '
		} else {
			res[i] = mapping[message[i]-'a']
		}
	}
	return string(res)
}
```
