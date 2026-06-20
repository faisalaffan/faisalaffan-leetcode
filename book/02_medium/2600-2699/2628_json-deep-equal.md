# 2628 — Json Deep Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func jsonDeepEqual(o1, o2 any) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2628: JSON Deep Equal
// https://leetcode.com/problems/json-deep-equal/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func jsonDeepEqual(o1, o2 any) bool {
	return reflect.DeepEqual(o1, o2)
}

func deepEqualJSON(a, b string) bool {
	var v1, v2 any
	if err := json.Unmarshal([]byte(a), &v1); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &v2); err != nil {
		return false
	}
	return jsonDeepEqual(v1, v2)
}

func main() {
	// Test case 1: equal objects
	fmt.Println("Test 1:", deepEqualJSON(`{"a":1,"b":2}`, `{"b":2,"a":1}`))
	// Expected: true

	// Test case 2: different types
	fmt.Println("Test 2:", deepEqualJSON(`{"a":1}`, `{"a":"1"}`))
	// Expected: false

	// Test case 3: arrays
	fmt.Println("Test 3:", deepEqualJSON(`[1,2,3]`, `[1,2,3]`))
	// Expected: true
}
```
