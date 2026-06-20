# 2590 — Design A Todo List

## Deskripsi

**Soal:** [2590. Design A Todo List](https://leetcode.com/problems/design-a-todo-list/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) per getUserTasks  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func Constructor() TodoList`

## Solusi Go

```go
package main

// LeetCode #2590: Design a Todo List
// https://leetcode.com/problems/design-a-todo-list/
// Difficulty: Medium [Paid]
// Time: O(n log n) per getUserTasks | Space: O(n)

import (
	"fmt"
	"sort"
)

type Task struct {
	ID        int
	Desc      string
	DueDate   int
	Tags      []string
	Completed bool
	UserID    int
}

type TodoList struct {
	tasks   []*Task
	nextID  int
}

func Constructor() TodoList {
	return TodoList{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func (this *TodoList) AddTask(userId int, taskDescription string, taskDueDate int, taskTags []string) int {
	id := this.nextID
	this.nextID++
	task := &Task{
		ID:        id,
		Desc:      taskDescription,
		DueDate:   taskDueDate,
		Tags:      taskTags,
		Completed: false,
		UserID:    userId,
	}
	this.tasks = append(this.tasks, task)
	return id
}

func (this *TodoList) GetAllTasks(userId int) []string {
	userTasks := []*Task{}
	for _, t := range this.tasks {
		if t.UserID == userId && !t.Completed {
			userTasks = append(userTasks, t)
		}
	}
	sort.Slice(userTasks, func(i, j int) bool {
		if userTasks[i].DueDate != userTasks[j].DueDate {
			return userTasks[i].DueDate < userTasks[j].DueDate
		}
		return userTasks[i].ID < userTasks[j].ID
	})
  // Membuat slice untuk menyimpan hasil
	res := make([]string, len(userTasks))
	for i, t := range userTasks {
		res[i] = t.Desc
	}
	return res
}

func (this *TodoList) GetTasksForTag(userId int, tag string) []string {
	userTasks := []*Task{}
	for _, t := range this.tasks {
		if t.UserID == userId && !t.Completed {
			for _, tg := range t.Tags {
				if tg == tag {
					userTasks = append(userTasks, t)
					break
				}
			}
		}
	}
	sort.Slice(userTasks, func(i, j int) bool {
		if userTasks[i].DueDate != userTasks[j].DueDate {
			return userTasks[i].DueDate < userTasks[j].DueDate
		}
		return userTasks[i].ID < userTasks[j].ID
	})
  // Membuat slice untuk menyimpan hasil
	res := make([]string, len(userTasks))
	for i, t := range userTasks {
		res[i] = t.Desc
	}
	return res
}

func (this *TodoList) CompleteTask(userId int, taskId int) {
	for _, t := range this.tasks {
		if t.ID == taskId && t.UserID == userId {
			t.Completed = true
			break
		}
	}
}

func main() {
	// Test case
	todo := Constructor()
	id1 := todo.AddTask(1, "Task1", 50, []string{})
	id2 := todo.AddTask(1, "Task2", 30, []string{"tag1"})
	todo.AddTask(2, "Task3", 10, []string{})
	fmt.Println("Test 1:", todo.GetAllTasks(1))
	// Expected: ["Task2", "Task1"]

	todo.CompleteTask(1, id2)
	fmt.Println("Test 2:", todo.GetAllTasks(1))
	// Expected: ["Task1"]

	todo.CompleteTask(1, id1)
	fmt.Println("Test 3:", todo.GetAllTasks(1))
	// Expected: []
}
```
