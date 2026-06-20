# 1776 — Car Fleet Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getCollisionTimes(cars [][]int) []float64
```

> **💡 Hint:** Monotonic stack (processing from right to left).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1776: Car Fleet II
// https://leetcode.com/problems/car-fleet-ii/
// Difficulty: Hard
//
// Approach: Monotonic stack (processing from right to left).
// Each car has (position, speed). Compute collision time with the next car
// ahead. Use a stack to track the collision chain.
//
// For car i, we check the car immediately ahead (j = i+1).
// If speed[i] <= speed[j], car i will never catch car j → answer[i] = -1.
// Otherwise, compute collision time t = (pos[j] - pos[i]) / (speed[i] - speed[j]).
// If the collision happens before or when car j hits its own blocker,
// then i collides with j at time t. Otherwise, i's collision time is
// determined by its collision with car j's blocker (recursive chain).

import (
	"fmt"
)

type Car struct {
	pos, speed float64
}

func getCollisionTimes(cars [][]int) []float64 {
	n := len(cars)
	ans := make([]float64, n)
	// Stack holds indices of cars that form collision chains (from right)
  // Alokasi slice integer
	stack := make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		pos := float64(cars[i][0])
		speed := float64(cars[i][1])

		ans[i] = -1.0

		// Remove cars that can never be caught
		for len(stack) > 0 {
			j := stack[len(stack)-1]
			// If current car is slower or equal speed, it can't catch car j
			if speed <= float64(cars[j][1]) {
				stack = stack[:len(stack)-1]
				continue
			}
			// Compute collision time with car j
			posJ := float64(cars[j][0])
			speedJ := float64(cars[j][1])
			t := (posJ - pos) / (speed - speedJ)

			// If car j has a collision and this happens after that,
			// car i will actually collide with whatever j collides with.
			if ans[j] > 0 && t >= ans[j] {
				stack = stack[:len(stack)-1]
				continue
			}

			ans[i] = t
			break
		}

		stack = append(stack, i)
	}

	return ans
}

func main() {
	// Example test case (LeetCode Example 1)
	cars := [][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}
	fmt.Println("cars=[[1,2],[2,1],[4,3],[7,2]] →", getCollisionTimes(cars))
	// Expected: [1, -1, 3, -1]

	// Additional tests
	cars2 := [][]int{{1, 1}, {2, 2}, {3, 4}}
	fmt.Println("cars=[[1,1],[2,2],[3,4]] →", getCollisionTimes(cars2))

	// Single car should have -1
	cars3 := [][]int{{1, 1}}
	fmt.Println("cars=[[1,1]] →", getCollisionTimes(cars3))

	// User-specified test (corrected interpretation)
	cars4 := [][]int{{1, 2}, {2, 1}, {3, 3}, {5, 4}}
	fmt.Println("cars=[[1,2],[2,1],[3,3],[5,4]] →", getCollisionTimes(cars4))
}
```
