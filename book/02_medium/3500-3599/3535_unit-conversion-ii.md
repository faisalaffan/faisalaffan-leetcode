# 3535 — Unit Conversion Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func UnitConversionIi(value float64, fromUnit, toUnit string) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3535: Unit Conversion II
// https://leetcode.com/problems/unit-conversion-ii/
// Difficulty: Medium [Paid]
// Complexity: O(n + q) time, O(n) space

import "fmt"

func main() {
	// Test case 1: temperature
	fmt.Println("Test 1:", UnitConversionIi(32, "F", "C"))
	// Test case 2
	fmt.Println("Test 2:", UnitConversionIi(0, "C", "F"))
	// Test case 3: mass
	fmt.Println("Test 3:", UnitConversionIi(1, "kg", "g"))
}

func UnitConversionIi(value float64, fromUnit, toUnit string) float64 {
	// Temperature conversions
	if fromUnit == "C" && toUnit == "F" {
		return value*9/5 + 32
	}
	if fromUnit == "F" && toUnit == "C" {
		return (value - 32) * 5 / 9
	}
	if fromUnit == "C" && toUnit == "K" {
		return value + 273.15
	}
	if fromUnit == "K" && toUnit == "C" {
		return value - 273.15
	}
	if fromUnit == "F" && toUnit == "K" {
		return (value-32)*5/9 + 273.15
	}
	if fromUnit == "K" && toUnit == "F" {
		return (value-273.15)*9/5 + 32
	}

	// Mass conversions (base: g)
	toGram := map[string]float64{
		"mg": 0.001,
		"g":  1.0,
		"kg": 1000.0,
		"lb": 453.592,
		"oz": 28.3495,
	}
	if f1, ok := toGram[fromUnit]; ok {
		if f2, ok2 := toGram[toUnit]; ok2 {
			return value * f1 / f2
		}
	}

	return value
}
```
