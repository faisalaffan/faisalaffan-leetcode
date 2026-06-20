# 1603 — Design Parking System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewParkingSystem(big int, medium int, small int) ParkingSystem
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1603: Design Parking System
// https://leetcode.com/problems/design-parking-system/
// Difficulty: Easy

import "fmt"

type ParkingSystem struct {
	spots [3]int
}

func NewParkingSystem(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{spots: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.spots[carType-1] > 0 {
		this.spots[carType-1]--
		return true
	}
	return false
}

func main() {
	ps := NewParkingSystem(1, 1, 0)
	fmt.Println(ps.AddCar(1))
	fmt.Println(ps.AddCar(2))
	fmt.Println(ps.AddCar(3))
	fmt.Println(ps.AddCar(1))
}
```
