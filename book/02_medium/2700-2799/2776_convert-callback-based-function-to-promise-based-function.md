# 2776 — Convert Callback Based Function To Promise Based Function

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConvertCallbackBasedFunctionToPromiseBasedFunction(fn Callback) PromiseFunc`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


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
