# 2671 — Frequency Tracker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() FrequencyTracker
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2671: Frequency Tracker
// https://leetcode.com/problems/frequency-tracker/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type FrequencyTracker struct {
	freq   map[int]int
	freqOf map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		freq:   make(map[int]int),
		freqOf: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	oldFreq := this.freq[number]
	if this.freqOf[oldFreq] > 0 {
		this.freqOf[oldFreq]--
	}
	this.freq[number]++
	this.freqOf[oldFreq+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.freq[number] == 0 {
		return
	}
	oldFreq := this.freq[number]
	this.freqOf[oldFreq]--
	this.freq[number]--
	if this.freq[number] > 0 {
		this.freqOf[oldFreq-1]++
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freqOf[frequency] > 0
}

func main() {
	ft := Constructor()

	// Test case 1
	ft.Add(3)
	ft.Add(3)
	fmt.Println("Test 1:", ft.HasFrequency(2))
	// Expected: true

	// Test case 2
	ft.DeleteOne(3)
	fmt.Println("Test 2:", ft.HasFrequency(1))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", ft.HasFrequency(2))
	// Expected: false
}
```
