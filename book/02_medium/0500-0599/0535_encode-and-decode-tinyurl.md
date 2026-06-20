# 0535 — Encode And Decode Tinyurl

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Encode(longUrl string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) for both encode and decode  
**Kompleksitas Ruang:** O(n) where n = number of encoded URLs

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #535: Encode and Decode TinyURL
// https://leetcode.com/problems/encode-and-decode-tinyurl/
// Difficulty: Medium
// Time: O(1) for both encode and decode
// Space: O(n) where n = number of encoded URLs

import (
	"fmt"
	"math/rand"
)

func main() {
	url := "https://leetcode.com/problems/design-tinyurl"
	encoded := Encode(url)
	fmt.Println("Encoded:", encoded)
	decoded := Decode(encoded)
	fmt.Println("Decoded:", decoded)
}

var urlMap = make(map[string]string)
var keyLen = 6
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(longUrl string) string {
	key := make([]byte, keyLen)
  // Range loop: iterasi dengan indeks + nilai
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}
	shortKey := string(key)
	urlMap[shortKey] = longUrl
	return shortKey
}

func Decode(shortUrl string) string {
	return urlMap[shortUrl]
}
```
