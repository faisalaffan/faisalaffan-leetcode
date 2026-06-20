# 3508 — Implement Router

## Deskripsi

**Soal:** [3508. Implement Router](https://leetcode.com/problems/implement-router/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func NewRouter() *Router`

## Solusi Go

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
