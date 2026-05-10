# Day 32: Factory Pattern

## 🎯 Goal
Use the **Factory Pattern** to create objects without exposing creation logic — essential for writing flexible, testable Go code.

## 🧠 Key Concepts
- Factory = a function/method that creates and returns objects
- In Go: no classes, so factory is just a constructor function `NewXxx()`
- **Simple Factory** — one function creates different types based on input
- **Factory Method** — each type implements a common interface, factory picks which one
- Decouples creation from usage → easy to swap implementations (e.g. in tests)

## 📖 Pattern

### Simple Factory
```go
type Notifier interface {
    Send(msg string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (e *EmailNotifier) Send(msg string) { fmt.Println("Email:", msg) }
func (s *SMSNotifier) Send(msg string)   { fmt.Println("SMS:", msg) }

// Factory function — caller doesn't know which struct is created
func NewNotifier(notifType string) Notifier {
    switch notifType {
    case "email":
        return &EmailNotifier{}
    case "sms":
        return &SMSNotifier{}
    default:
        return &EmailNotifier{} // safe default
    }
}
```

### Usage
```go
n := NewNotifier("sms")
n.Send("Your order shipped!") // → SMS: Your order shipped!
```

## ✅ Learning Checklist
- [ ] Implement a factory function that returns an interface
- [ ] Understand why returning interface > returning concrete type (flexibility)
- [ ] Use factory pattern to swap real DB with mock DB in tests
- [ ] Know the difference between Simple Factory and Abstract Factory

## 🛠️ Practice Exercises
1. **Notifier Factory** — Email, SMS, Push notifier — factory picks based on type string
2. **Worker Factory** — create different worker types (FastWorker, SlowWorker) with same `Process(job int) int` interface
3. **Storage Factory** — `NewStorage("memory")` or `NewStorage("file")` — same `Save(key, value string)` interface
4. **BONUS** — Implement Abstract Factory: create a UI factory that produces `Button` + `Checkbox` for different themes (dark/light)

## 🔍 Interview Questions
- "How do you implement dependency injection in Go?"
- "How would you make your DB layer testable?" (answer: factory returns interface, swap in tests)
- "What's the difference between factory and constructor?"
