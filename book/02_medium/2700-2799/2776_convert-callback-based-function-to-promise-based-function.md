# 2776 — Convert Callback Based Function To Promise Based Function

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConvertCallbackBasedFunctionToPromiseBasedFunction(fn Callback) PromiseFunc
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2776: Convert Callback Based Function to Promise Based Function
// https://leetcode.com/problems/convert-callback-based-function-to-promise-based-function/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

type Callback func(args ...interface{}) (interface{}, error)
type PromiseFunc func(args ...interface{}) (<-chan interface{}, <-chan error)

func ConvertCallbackBasedFunctionToPromiseBasedFunction(fn Callback) PromiseFunc {
	return func(args ...interface{}) (<-chan interface{}, <-chan error) {
		resCh := make(chan interface{}, 1)
		errCh := make(chan error, 1)
		go func() {
			result, err := fn(args...)
			if err != nil {
				errCh <- err
			} else {
				resCh <- result
			}
			close(resCh)
			close(errCh)
		}()
		return resCh, errCh
	}
}

func main() {
	add := func(args ...interface{}) (interface{}, error) {
		return args[0].(int) + args[1].(int), nil
	}
	promiseFn := ConvertCallbackBasedFunctionToPromiseBasedFunction(add)
	resCh, errCh := promiseFn(3, 4)
	select {
	case res := <-resCh:
		fmt.Println(res)
	case err := <-errCh:
		fmt.Println(err)
	}

	// Test with single arg
	promiseFn2 := ConvertCallbackBasedFunctionToPromiseBasedFunction(func(args ...interface{}) (interface{}, error) {
		return args[0].(int) * 2, nil
	})
	resCh2, _ := promiseFn2(10)
	select {
	case res := <-resCh2:
		fmt.Println(res)
	}
}
```
