# 2704 — To Be Or Not To Be

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ToBeOrNotToBe(val interface{}) *Expect
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2704: To Be Or Not To Be
// https://leetcode.com/problems/to-be-or-not-to-be/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Expect-type assertion.

import (
	"fmt"
	"reflect"
)

func main() {
	{
		expect := ToBeOrNotToBe(5)
		result, err := expect.toBe(5)
		fmt.Println(result, err)
	}
	{
		expect := ToBeOrNotToBe(5)
		result, err := expect.notToBe(5)
		fmt.Println(result, err)
	}
}

type Expect struct {
	val interface{}
}

func (e *Expect) toBe(expected interface{}) (bool, string) {
	if reflect.DeepEqual(e.val, expected) {
		return true, ""
	}
	return false, "Not Equal"
}

func (e *Expect) notToBe(expected interface{}) (bool, string) {
	if !reflect.DeepEqual(e.val, expected) {
		return true, ""
	}
	return false, "Equal"
}

func ToBeOrNotToBe(val interface{}) *Expect {
	return &Expect{val: val}
}
```
