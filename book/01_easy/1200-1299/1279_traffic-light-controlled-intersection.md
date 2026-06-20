# 1279 — Traffic Light Controlled Intersection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NewTrafficLight() *TrafficLight`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


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
