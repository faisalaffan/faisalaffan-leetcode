# 2147 — Number Of Ways To Divide A Long Corridor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfWays(corridor string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2147: Number of Ways to Divide a Long Corridor
// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/
// Difficulty: Hard
//
// Collect indices of seats. Each pair of seats (consecutive in the even-odd sense)
// defines a mandatory section boundary. Multiply gaps between pairs.

import "fmt"

func main() {
	fmt.Println(numberOfWays("SSPPSPS")) // 3
	fmt.Println(numberOfWays("PPSPSP"))  // 1
	fmt.Println(numberOfWays("S"))       // 0
	fmt.Println(numberOfWays("SS"))      // 1
	fmt.Println(numberOfWays("SPPSS"))   // 2
}

const mod2147 = 1_000_000_007

func numberOfWays(corridor string) int {
	var seats []int
	for i, ch := range corridor {
		if ch == 'S' {
			seats = append(seats, i)
		}
	}
	if len(seats) == 0 || len(seats)%2 != 0 {
		return 0
	}

	ans := 1
	for i := 1; i < len(seats)-1; i += 2 {
		gap := seats[i+1] - seats[i]
		ans = (ans * gap) % mod2147
	}
	return ans
}

func NumberOfWaysToDivideALongCorridor() any {
	return numberOfWays("SSPPSPS")
}
```
