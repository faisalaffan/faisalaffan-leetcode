# 1826 — Faulty Sensor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BadSensor(sensor1 []int, sensor2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1826: Faulty Sensor
// https://leetcode.com/problems/faulty-sensor/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func BadSensor(sensor1 []int, sensor2 []int) int {
	n := len(sensor1)
	i := 0
	for i < n-1 && sensor1[i] == sensor2[i] {
		i++
	}
	if i == n-1 {
		return -1
	}
	// Try sensor1 as faulty (sensor1 drops value at i, rest shifted)
	ok1 := true
	for j := i; j < n-1; j++ {
		if sensor1[j] != sensor2[j+1] {
			ok1 = false
			break
		}
	}
	// Try sensor2 as faulty
	ok2 := true
	for j := i; j < n-1; j++ {
		if sensor2[j] != sensor1[j+1] {
			ok2 = false
			break
		}
	}
	if ok1 && !ok2 {
		return 1
	}
	if !ok1 && ok2 {
		return 2
	}
	return -1
}

func main() {
	fmt.Println(BadSensor([]int{2, 3, 4, 5}, []int{2, 1, 3, 4}))
	fmt.Println(BadSensor([]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 5}))
	fmt.Println(BadSensor([]int{2, 3, 2, 2, 3, 2}, []int{2, 3, 2, 3, 2, 7}))
}
```
