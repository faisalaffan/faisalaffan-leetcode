# 2618 — Check If Object Instance Of Class

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func checkIfInstanceOf(obj any, classType reflect.Type) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(d) where d is depth of type hierarchy  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2618: Check if Object Instance of Class
// https://leetcode.com/problems/check-if-object-instance-of-class/
// Difficulty: Medium
// Time: O(d) where d is depth of type hierarchy | Space: O(1)

import (
	"fmt"
	"reflect"
)

func checkIfInstanceOf(obj any, classType reflect.Type) bool {
	if obj == nil || classType == nil {
		return false
	}

	t := reflect.TypeOf(obj)
	for t != nil {
		if t == classType {
			return true
		}
		if classType.Kind() == reflect.Interface && t.Implements(classType) {
			return true
		}
		// Move to pointer type and check
		ptrT := reflect.PointerTo(t)
		if ptrT == classType {
			return true
		}
		// Move to underlying type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}
	return false
}

type MyClass struct{}

func main() {
	// Test case 1
	obj := &MyClass{}
	fmt.Println("Test 1:", checkIfInstanceOf(obj, reflect.TypeOf(MyClass{})))
	// Expected: true

	// Test case 2: int is instance of int
	fmt.Println("Test 2:", checkIfInstanceOf(5, reflect.TypeOf(0)))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", checkIfInstanceOf("hello", reflect.TypeOf(0)))
	// Expected: false
}
```
