# 0735 — Asteroid Collision

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func asteroidCollision(asteroids []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #735: Asteroid Collision
// https://leetcode.com/problems/asteroid-collision/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
}

func asteroidCollision(asteroids []int) []int {
  // Alokasi slice integer
	stack := make([]int, 0)

	for _, a := range asteroids {
		for len(stack) > 0 && a < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top+ a < 0 {
				stack = stack[:len(stack)-1]
			} else if top+ a == 0 {
				stack = stack[:len(stack)-1]
				a = 0
				break
			} else {
				a = 0
				break
			}
		}
		if a != 0 {
			stack = append(stack, a)
		}
	}

	return stack
}
```
