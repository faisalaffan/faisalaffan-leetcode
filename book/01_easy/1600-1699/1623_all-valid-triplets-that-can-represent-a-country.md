# 1623 — All Valid Triplets That Can Represent A Country

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func AllValidTripletsThatCanRepresentACountry() any`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #1623: All Valid Triplets That Can Represent a Country
// https://leetcode.com/problems/all-valid-triplets-that-can-represent-a-country/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(AllValidTripletsThatCanRepresentACountry())
}

func AllValidTripletsThatCanRepresentACountry() any {
	// TODO: implement
	return nil
}
```
