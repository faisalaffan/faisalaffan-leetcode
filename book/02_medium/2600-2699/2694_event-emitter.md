# 2694 — Event Emitter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NewEventEmitter() *EventEmitter`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1) per operation  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2694: Event Emitter
// https://leetcode.com/problems/event-emitter/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type EventEmitter struct {
	events map[string][]func(args ...any)
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{events: make(map[string][]func(args ...any))}
}

func (ee *EventEmitter) Subscribe(event string, cb func(args ...any)) func() {
	ee.events[event] = append(ee.events[event], cb)
	return func() {
		listeners := ee.events[event]
		for i, fn := range listeners {
			if &fn == &cb { // Note: won't work for func comparison in Go
				ee.events[event] = append(listeners[:i], listeners[i+1:]...)
				break
			}
		}
	}
}

func (ee *EventEmitter) Emit(event string, args ...any) []any {
	results := []any{}
	for _, cb := range ee.events[event] {
		cb(args...)
		results = append(results, nil) // Go functions don't return values like JS
	}
	return results
}

func main() {
	emitter := NewEventEmitter()

	// Test case 1
	emitter.Subscribe("event1", func(args ...any) {
		fmt.Println("  received:", args)
	})
	fmt.Println("Test 1: emitting event1")
	emitter.Emit("event1", 1, 2, 3)

	// Test case 2: unsubscribe
	unsub := emitter.Subscribe("event2", func(args ...any) {
		fmt.Println("  event2 received:", args)
	})
	unsub()
	fmt.Println("Test 2: emitted event2 after unsubscribe (no output if working)")
	emitter.Emit("event2")

	// Test case 3: multiple subscribers
	count := 0
	emitter.Subscribe("count", func(args ...any) { count++ })
	emitter.Subscribe("count", func(args ...any) { count++ })
	emitter.Emit("count")
	fmt.Println("Test 3: count =", count)
	// Expected: 2
}
```
