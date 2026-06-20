# 0535 — Encode And Decode Tinyurl

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func Encode(longUrl string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1) for both encode and decode  |  **Ruang:** O(n) where n = number of encoded URLs

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // Range loop
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
