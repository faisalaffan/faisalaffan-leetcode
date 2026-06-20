# 1600 — Throne Inheritance

## Deskripsi

**Soal:** [1600. Throne Inheritance](https://leetcode.com/problems/throne-inheritance/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1600: Throne Inheritance
// https://leetcode.com/problems/throne-inheritance/
// Difficulty: Medium

import "fmt"

func main() {
	t := ConstructorThrone("king")
	t.Birth("king", "andy")
	t.Birth("king", "bob")
	t.Birth("king", "catherine")
	t.Birth("andy", "matthew")
	t.Birth("bob", "alex")
	t.Birth("bob", "asha")

	inheritance := t.GetInheritanceOrder()
	fmt.Println("Inheritance order:", inheritance)

	t.Death("bob")
	inheritance2 := t.GetInheritanceOrder()
	fmt.Println("After Bob's death:", inheritance2)
}

type ThroneInheritance struct {
	king     string
	children map[string][]string
	dead     map[string]bool
}

func ConstructorThrone(kingName string) ThroneInheritance {
	return ThroneInheritance{
		king:     kingName,
		children: make(map[string][]string),
		dead:     make(map[string]bool),
	}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.children[parentName] = append(t.children[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	// Preorder traversal of the family tree
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	t.dfs(t.king, &result)
	return result
}

func (t *ThroneInheritance) dfs(name string, result *[]string) {
	if !t.dead[name] {
		*result = append(*result, name)
	}
	for _, child := range t.children[name] {
		t.dfs(child, result)
	}
}
```
