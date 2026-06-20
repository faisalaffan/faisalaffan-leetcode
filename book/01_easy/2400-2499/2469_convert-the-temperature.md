# 2469 — Convert The Temperature

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConvertTheTemperature(celsius float64) []float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2469: Convert the Temperature
// https://leetcode.com/problems/convert-the-temperature/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(ConvertTheTemperature(36.50)) // [309.65, 97.7]
	fmt.Println(ConvertTheTemperature(122.11)) // [395.26, 251.798]
}

func ConvertTheTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32.0}
}
```
