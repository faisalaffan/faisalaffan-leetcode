# 2618 — Check If Object Instance Of Class

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func checkIfInstanceOf(obj any, classType reflect.Type) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(d) where d is depth of type hierarchy  |  **Ruang:** O(1)


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
