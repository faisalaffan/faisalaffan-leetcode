# 0595 — Big Countries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func BigCountries() string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Trie

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Trie** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #595: Big Countries
// https://leetcode.com/problems/big-countries/
// Difficulty: Easy

import "fmt"

func BigCountries() string {
	return "SELECT name, population, area FROM World WHERE area >= 3000000 OR population >= 25000000"
}

func main() {
	fmt.Println(BigCountries())
}
```
