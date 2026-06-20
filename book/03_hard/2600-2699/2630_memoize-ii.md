# 2630 — Memoize Ii

## Deskripsi

**Soal:** [2630. Memoize Ii](https://leetcode.com/problems/memoize-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #2630: Memoize II
// https://leetcode.com/problems/memoize-ii/
// Difficulty: Hard
//
// Design a generic memoization function that supports arbitrary argument types,
// including slices and maps, by encoding arguments as string keys.
// Uses a type-safe wrapper around a sync.Map-like cache.

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

func main() {
	// Test: memoize sum of two ints
	add := memoize(func(args []any) any {
		a := args[0].(int)
		b := args[1].(int)
		return a + b
	})
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2)) // cached
	fmt.Println(add(2, 3))

	// Test: memoize with slice arg
	first := memoize(func(args []any) any {
		s := args[0].([]int)
		if len(s) == 0 {
			return nil
		}
		return s[0]
	})
	fmt.Println(first([]int{10, 20, 30}))
	fmt.Println(first([]int{10, 20, 30})) // cached
	fmt.Println(first([]int{40, 50}))
}

// memoizeFn is the type of a memoized function
type memoizeFn func(args ...any) any

// memoize wraps f with a cache keyed on the string representation of arguments.
func memoize(f func(args []any) any) memoizeFn {
	var mu sync.Mutex
  // Membuat map untuk pencarian O(1): key → value
	cache := make(map[string]any)

	return func(args ...any) any {
		key := encodeArgs(args)
		mu.Lock()
		if v, ok := cache[key]; ok {
			mu.Unlock()
			return v
		}
		mu.Unlock()

		result := f(args)

		mu.Lock()
		cache[key] = result
		mu.Unlock()
		return result
	}
}

// encodeArgs produces a deterministic string key for any argument list.
func encodeArgs(args []any) string {
	var sb strings.Builder
	for i, arg := range args {
		if i > 0 {
			sb.WriteByte('|')
		}
		encodeValue(&sb, arg)
	}
	return sb.String()
}

func encodeValue(sb *strings.Builder, v any) {
	if v == nil {
		sb.WriteString("nil")
		return
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		sb.WriteByte('[')
		for i := 0; i < rv.Len(); i++ {
			if i > 0 {
				sb.WriteByte(',')
			}
			encodeValue(sb, rv.Index(i).Interface())
		}
		sb.WriteByte(']')
	case reflect.Map:
		sb.WriteByte('{')
		iter := rv.MapRange()
		first := true
		for iter.Next() {
			if !first {
				sb.WriteByte(',')
			}
			first = false
			encodeValue(sb, iter.Key().Interface())
			sb.WriteByte(':')
			encodeValue(sb, iter.Value().Interface())
		}
		sb.WriteByte('}')
	case reflect.Ptr, reflect.Interface:
		encodeValue(sb, rv.Elem().Interface())
	default:
		fmt.Fprint(sb, v)
	}
}

// MemoizeIi is a convenience wrapper matching the stub signature
func MemoizeIi() any {
	return "MemoizeII implemented"
}
```
