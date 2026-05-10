package main

import (
	"fmt"
	"strings"
)

// ========================================
// Exercise 1: Notifier Factory
// ========================================

type Notifier interface {
	Send(msg string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}
type PushNotifier struct{}

// TODO: implement Send for each notifier
// Email: fmt.Println("📧 Email:", msg)
// SMS:   fmt.Println("📱 SMS:", msg)
// Push:  fmt.Println("🔔 Push:", msg)

// TODO: Implement NewNotifier(notifType string) Notifier
// Handle "email", "sms", "push" — default to email
func NewNotifier(notifType string) Notifier {
	// TODO: implement
	return nil
}

func exercise1() {
	types := []string{"email", "sms", "push", "unknown"}
	for _, t := range types {
		n := NewNotifier(t)
		if n != nil {
			n.Send("Hello from factory!")
		}
	}
}

// ========================================
// Exercise 2: Worker Factory
// ========================================

type Worker interface {
	Process(job int) int
}

type FastWorker struct{}
type SlowWorker struct{}

// FastWorker: multiplies job by 10
// SlowWorker: adds 1 to job (simulates slow processing)
// TODO: implement Process for both

// TODO: Implement NewWorker(workerType string) Worker
func NewWorker(workerType string) Worker {
	// TODO: implement
	return nil
}

func exercise2() {
	fast := NewWorker("fast")
	slow := NewWorker("slow")
	jobs := []int{1, 2, 3}
	for _, j := range jobs {
		fmt.Printf("Job %d → fast: %d, slow: %d\n", j, fast.Process(j), slow.Process(j))
	}
}

// ========================================
// Exercise 3: Storage Factory
// ========================================

type Storage interface {
	Save(key, value string)
	Get(key string) string
}

type MemoryStorage struct {
	data map[string]string
}

type FileStorage struct {
	// simulated — just print what would be written
	prefix string
}

// TODO: implement Save and Get for MemoryStorage (use the map)
// TODO: implement Save and Get for FileStorage (just print "FileStorage: Save key=..., value=...")

// TODO: Implement NewStorage(storageType string) Storage
// "memory" → MemoryStorage, "file" → FileStorage
func NewStorage(storageType string) Storage {
	// TODO: implement
	_ = strings.ToLower // hint
	return nil
}

func exercise3() {
	mem := NewStorage("memory")
	mem.Save("name", "Golang")
	fmt.Println(mem.Get("name"))

	file := NewStorage("file")
	file.Save("config", "production")
	fmt.Println(file.Get("config"))
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 32: Factory Pattern ===\n")

	fmt.Println("Exercise 1: Notifier Factory")
	exercise1()

	fmt.Println("\nExercise 2: Worker Factory")
	exercise2()

	fmt.Println("\nExercise 3: Storage Factory")
	exercise3()
}

/*
SOLUTIONS:

// Exercise 1:
func (e *EmailNotifier) Send(msg string) { fmt.Println("📧 Email:", msg) }
func (s *SMSNotifier) Send(msg string)   { fmt.Println("📱 SMS:", msg) }
func (p *PushNotifier) Send(msg string)  { fmt.Println("🔔 Push:", msg) }

func NewNotifier(notifType string) Notifier {
	switch notifType {
	case "sms":   return &SMSNotifier{}
	case "push":  return &PushNotifier{}
	default:      return &EmailNotifier{}
	}
}

// Exercise 2:
func (f *FastWorker) Process(job int) int { return job * 10 }
func (s *SlowWorker) Process(job int) int { return job + 1 }

func NewWorker(workerType string) Worker {
	if workerType == "fast" { return &FastWorker{} }
	return &SlowWorker{}
}

// Exercise 3:
func (m *MemoryStorage) Save(key, value string) { m.data[key] = value }
func (m *MemoryStorage) Get(key string) string  { return m.data[key] }
func (f *FileStorage) Save(key, value string)   { fmt.Printf("FileStorage: Save %s=%s\n", key, value) }
func (f *FileStorage) Get(key string) string    { fmt.Printf("FileStorage: Get %s\n", key); return "" }

func NewStorage(storageType string) Storage {
	switch strings.ToLower(storageType) {
	case "file":   return &FileStorage{prefix: "app"}
	default:       return &MemoryStorage{data: make(map[string]string)}
	}
}
*/
