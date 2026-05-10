# Day 3: Basic Algorithms

## 🎯 Today's Goal
Understand common algorithms and bitwise operations.

## 📚 Files to Study
1. `../../fibbonacci.go` - Fibonacci with channels
2. `../../xor.go` - XOR operations

## 🧠 Key Concepts
- Fibonacci sequence (iterative vs recursive)
- Using goroutines with channels
- Bitwise operations (XOR, AND, OR)
- Channel patterns for generators

## ✅ Learning Checklist
- [ ] Can generate Fibonacci numbers
- [ ] Understand how channels are used for sequences
- [ ] Know when to close a channel
- [ ] Understand XOR properties
- [ ] Can use range with channels

## 🛠️ Practice Exercises

### Exercise 1: Fibonacci Without Goroutines
```go
// Write iterative Fibonacci (no channels)
func fib(n int) []int {
    // Return first n Fibonacci numbers
}
```

### Exercise 2: Understanding the Channel Pattern
```go
// Modify the fibonacci channel to generate even Fibonacci numbers only
```

### Exercise 3: XOR Properties
```go
// Use XOR to:
// 1. Swap two numbers without a temp variable
// 2. Find the single number in an array where all others appear twice
```

## 📝 Study Steps

### Step 1: Study fibonacci.go (7 min)
- Understand how the channel is created
- See how the goroutine generates values
- Notice when the channel is closed
- Understand the range loop consuming values

### Step 2: Study xor.go (3 min)
- Learn XOR properties
- See practical applications

### Step 3: Practice (10 min)
Complete the exercises

## 🎓 Key Patterns

**Generator Pattern with Channels**:
```go
func generator() chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch) // Important!
    }()
    return ch
}

// Usage
for val := range generator() {
    fmt.Println(val)
}
```

**XOR Magic**:
- `a ^ a = 0`
- `a ^ 0 = a`
- `a ^ b ^ b = a`

## ✅ Completion
- [ ] Studied both files
- [ ] Completed exercises
- [ ] Understand generator pattern
- [ ] Updated PROGRESS_TRACKER.md
- [ ] Confidence: ___/10

## 🔄 Spaced Repetition
- Review Day 1 (Pointers) - 5 minutes

## 🔜 Tomorrow: Day 4 - JSON & Data Handling
