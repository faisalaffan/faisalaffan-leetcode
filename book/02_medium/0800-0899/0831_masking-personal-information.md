# 0831 — Masking Personal Information

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaskingPersonalInformation(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #831: Masking Personal Information
// https://leetcode.com/problems/masking-personal-information/
// Difficulty: Medium

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(MaskingPersonalInformation("LeetCode@LeetCode.com"))
	fmt.Println(MaskingPersonalInformation("AB@qq.com"))
	fmt.Println(MaskingPersonalInformation("1(234)567-890"))
	fmt.Println(MaskingPersonalInformation("86-(10)12345678"))
}

// Time: O(n) | Space: O(n)
func MaskingPersonalInformation(s string) string {
	at := strings.IndexByte(s, '@')
	if at != -1 {
		s = strings.ToLower(s)
		return string(s[0]) + "*****" + string(s[at-1]) + s[at:]
	}

	var digits strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits.WriteRune(c)
		}
	}
	d := digits.String()
	local := d[len(d)-4:]
	masked := "***-***-" + local

	if len(d) == 10 {
		return masked
	}
	return "+" + strings.Repeat("*", len(d)-10) + "-" + masked
}
```
