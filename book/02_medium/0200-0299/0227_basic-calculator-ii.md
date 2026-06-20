# 0227 — Basic Calculator Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func calculate(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #227: Basic Calculator II
// https://leetcode.com/problems/basic-calculator-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func calculate(s string) int {
	stack := []int{}
	num := 0
	op := '+'

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}

		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			if s[i] == ' ' && i != len(s)-1 {
				continue
			}
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			op = ch
			num = 0
		}
	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}
```
