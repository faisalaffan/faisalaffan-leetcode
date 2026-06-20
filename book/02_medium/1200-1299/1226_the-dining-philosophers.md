# 1226 — The Dining Philosophers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewDiningPhilosophers() *DiningPhilosophers
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** O(1) per eat call  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1226: The Dining Philosophers
// https://leetcode.com/problems/the-dining-philosophers/
// Difficulty: Medium

// 5 philosophers, 5 forks. Each needs 2 forks to eat.
// Deadlock avoidance: odd philosophers pick left then right,
// even philosophers pick right then left.

// Time: O(1) per eat call
// Space: O(1)

type DiningPhilosophers struct {
	mu     sync.Mutex
	forks  [5]sync.Mutex
}

func NewDiningPhilosophers() *DiningPhilosophers {
	return &DiningPhilosophers{}
}

func (dp *DiningPhilosophers) WantsToEat(philosopher int,
	eat func(),
	pickLeftFork func(),
	pickRightFork func(),
	putLeftFork func(),
	putRightFork func()) {

	left := philosopher
	right := (philosopher + 1) % 5

	// To avoid deadlock: always pick lower-numbered fork first
	if left < right {
		dp.forks[left].Lock()
		dp.forks[right].Lock()
	} else {
		dp.forks[right].Lock()
		dp.forks[left].Lock()
	}

	pickLeftFork()
	pickRightFork()
	eat()
	putLeftFork()
	putRightFork()

	dp.forks[left].Unlock()
	dp.forks[right].Unlock()
}

func main() {
	dp := NewDiningPhilosophers()
	var wg sync.WaitGroup

	eatCount := 0
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				dp.WantsToEat(p,
					func() {
						mu.Lock()
						eatCount++
						mu.Unlock()
					},
					func() { fmt.Printf("P%d picks left\n", p) },
					func() { fmt.Printf("P%d picks right\n", p) },
					func() { fmt.Printf("P%d puts left\n", p) },
					func() { fmt.Printf("P%d puts right\n", p) },
				)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Total eats: %d (expected: 10)\n", eatCount)
}
```
