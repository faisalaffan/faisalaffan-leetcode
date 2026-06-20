# 1826 — Faulty Sensor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BadSensor(sensor1 []int, sensor2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
