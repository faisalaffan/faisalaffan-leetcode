# 3508 — Implement Router

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewRouter() *Router
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3508: Implement Router
// https://leetcode.com/problems/implement-router/
// Difficulty: Medium
// Complexity: O(1) per operation

import "fmt"

type Router struct {
	handlers map[string]func()
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]func())}
}

func (r *Router) AddRoute(path string, handler func()) {
	r.handlers[path] = handler
}

func (r *Router) Handle(path string) {
	if handler, ok := r.handlers[path]; ok {
		handler()
	}
}

func (r *Router) PrintRoutes() {
	for path := range r.handlers {
		fmt.Println("Route:", path)
	}
}

func main() {
	// Test case 1
	router := NewRouter()
	router.AddRoute("/home", func() { fmt.Println("Home") })
	router.AddRoute("/about", func() { fmt.Println("About") })
	router.Handle("/home")
	router.Handle("/about")
	router.PrintRoutes()

	// Test case 2
	router2 := NewRouter()
	router2.AddRoute("/api", func() { fmt.Println("API") })
	router2.Handle("/api")

	fmt.Println("ImplementRouter() done:", ImplementRouter())
}

func ImplementRouter() any {
	return "Router implemented"
}
```
