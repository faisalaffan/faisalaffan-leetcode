# 2700 — Differences Between Two Objects

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func objectDiff(obj1, obj2 any) map[string]any
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2700: Differences Between Two Objects
// https://leetcode.com/problems/differences-between-two-objects/
// Difficulty: Medium [Paid] (JS problem)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"reflect"
)

func objectDiff(obj1, obj2 any) map[string]any {
  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[string]any)
	diff(obj1, obj2, "", result)
	return result
}

func diff(a, b any, path string, result map[string]any) {
	if reflect.DeepEqual(a, b) {
		return
	}

	if a == nil || b == nil {
		result[path] = []any{a, b}
		return
	}

	map1, ok1 := a.(map[string]any)
	map2, ok2 := b.(map[string]any)
	if ok1 && ok2 {
  // Membuat map (HashMap) — pencarian O(1)
		allKeys := make(map[string]bool)
		for k := range map1 {
			allKeys[k] = true
		}
		for k := range map2 {
			allKeys[k] = true
		}
		for k := range allKeys {
			newPath := path
			if newPath == "" {
				newPath = k
			} else {
				newPath = path + "." + k
			}
			diff(map1[k], map2[k], newPath, result)
		}
		return
	}

	result[path] = []any{a, b}
}

func main() {
	// Test case 1: simple diff
	o1 := map[string]any{"a": 1, "b": 2}
	o2 := map[string]any{"a": 1, "b": 3}
	fmt.Println("Test 1:", objectDiff(o1, o2))
	// Expected: map[b:[2 3]]

	// Test case 2: missing key
	o3 := map[string]any{"a": 1, "c": 3}
	fmt.Println("Test 2:", objectDiff(o1, o3))
	// Expected: map[b:[2 <nil>] c:[<nil> 3]]

	// Test case 3: same objects
	fmt.Println("Test 3:", objectDiff(o1, o1))
	// Expected: map[]
}
```
