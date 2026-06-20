# 2754 — Bind Function To Context

## Deskripsi

**Soal:** [2754. Bind Function To Context](https://leetcode.com/problems/bind-function-to-context/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func BindFunctionToContext(fn func(Context, ...interface{}) interface`

## Solusi Go

```go
package main

// LeetCode #2754: Bind Function to Context
// https://leetcode.com/problems/bind-function-to-context/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

type Context map[string]interface{}

type BoundFunc struct {
	fn      func(ctx Context, args ...interface{}) interface{}
	ctx     Context
}

func BindFunctionToContext(fn func(Context, ...interface{}) interface{}, ctx Context) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		return fn(ctx, args...)
	}
}

func main() {
	ctx := Context{"multiplier": 3}
	bound := BindFunctionToContext(func(ctx Context, args ...interface{}) interface{} {
		m := ctx["multiplier"].(int)
		return args[0].(int) * m
	}, ctx)
	fmt.Println(bound(5))
	fmt.Println(bound(10))
}
```
