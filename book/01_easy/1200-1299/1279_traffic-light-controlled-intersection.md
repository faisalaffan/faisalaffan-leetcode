# 1279 — Traffic Light Controlled Intersection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewTrafficLight() *TrafficLight
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1279: Traffic Light Controlled Intersection
// https://leetcode.com/problems/traffic-light-controlled-intersection/
// Difficulty: Easy [Paid] (Concurrency)
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
)

type TrafficLight struct {
	mu      sync.Mutex
	greenOn int // 1 = road A, 2 = road B
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{greenOn: 1}
}

func (t *TrafficLight) CarArrived(carId, roadId, direction int, turnGreen, crossCar func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.greenOn != roadId {
		t.greenOn = roadId
		turnGreen()
	}
	crossCar()
}

func main() {
	light := NewTrafficLight()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		light.CarArrived(1, 1, 1, func() {
			fmt.Println("Turn green for Road A")
		}, func() {
			fmt.Println("Car 1 crossing on Road A")
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		light.CarArrived(2, 1, 2, func() {
			fmt.Println("Turn green for Road A")
		}, func() {
			fmt.Println("Car 2 crossing on Road A")
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		light.CarArrived(3, 2, 4, func() {
			fmt.Println("Turn green for Road B")
		}, func() {
			fmt.Println("Car 3 crossing on Road B")
		})
	}()

	wg.Wait()
}
```
