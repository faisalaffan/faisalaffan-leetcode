# 0735 — Asteroid Collision

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func asteroidCollision(asteroids []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

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
  // Alokasi slice
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
