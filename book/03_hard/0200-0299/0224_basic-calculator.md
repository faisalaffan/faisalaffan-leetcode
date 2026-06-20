# 0224 — Basic Calculator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func calculate(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #224: Basic Calculator
// https://leetcode.com/problems/basic-calculator/
// Difficulty: Hard

import "fmt"

func calculate(s string) int {
  // Alokasi slice integer
	stack := make([]int, 0)
	result := 0
	sign := 1
	num := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
			i--
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			stack = append(stack, result, sign)
			result = 0
			sign = 1
		} else if ch == ')' {
			result = stack[len(stack)-2] + stack[len(stack)-1]*result
			stack = stack[:len(stack)-2]
		}
	}

	return result
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
```
