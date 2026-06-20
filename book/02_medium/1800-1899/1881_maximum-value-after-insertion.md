# 1881 — Maximum Value After Insertion

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxValue(n string, x int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1881: Maximum Value After Insertion
// https://leetcode.com/problems/maximum-value-after-insertion/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxValue("99", 9))
	fmt.Println(MaxValue("-13", 2))
	fmt.Println(MaxValue("73", 6))
}

// Time: O(n), Space: O(n)
func MaxValue(n string, x int) string {
	result := make([]byte, 0, len(n)+1)
	negative := n[0] == '-'

	if negative {
		result = append(result, '-')
		inserted := false
		for i := 1; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x < digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	} else {
		inserted := false
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(n); i++ {
			digit := int(n[i] - '0')
			if !inserted && x > digit {
				result = append(result, byte(x+'0'))
				inserted = true
			}
			result = append(result, n[i])
		}
		if !inserted {
			result = append(result, byte(x+'0'))
		}
	}
	return string(result)
}
```
