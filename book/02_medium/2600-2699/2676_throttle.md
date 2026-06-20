# 2676 — Throttle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func newThrottle(fn func(int) , interval time.Duration) *ThrottledFn
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(1) per call  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2676: Throttle
// https://leetcode.com/problems/throttle/
// Difficulty: Medium [Paid]
// Time: O(1) per call | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

type ThrottledFn struct {
	fn       func(int)
	interval time.Duration
	mu       sync.Mutex
	timer    *time.Timer
	lastCall time.Time
	pending  bool
	lastArg  int
}

func (t *ThrottledFn) Call(arg int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.lastArg = arg
	if !t.pending {
		t.pending = true
		go func() {
			time.Sleep(t.interval)
			t.mu.Lock()
			t.fn(t.lastArg)
			t.pending = false
			t.mu.Unlock()
		}()
	}
}

func newThrottle(fn func(int), interval time.Duration) *ThrottledFn {
	return &ThrottledFn{
		fn:       fn,
		interval: interval,
	}
}

func main() {
	results := []int{}
	mu := sync.Mutex{}
	throttled := newThrottle(func(x int) {
		mu.Lock()
		results = append(results, x)
		mu.Unlock()
	}, 50*time.Millisecond)

	throttled.Call(1)
	throttled.Call(2)
	throttled.Call(3)

	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	fmt.Println("Test 1:", results)
	mu.Unlock()
	// Expected: [3] (only last call within throttle window)

	throttled.Call(4)
	time.Sleep(100 * time.Millisecond)
	mu.Lock()
	fmt.Println("Test 2:", results)
	mu.Unlock()
	// Expected: [3, 4]
}
```
