# 0535 — Encode And Decode Tinyurl

## Deskripsi

**Soal:** [0535. Encode And Decode Tinyurl](https://leetcode.com/problems/encode-and-decode-tinyurl/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) for both encode and decode  
**Kompleksitas Ruang:** O(n) where n = number of encoded URLs

**Algoritma:** —

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	key := make([]byte, keyLen)
  // Iterasi seluruh elemen
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
