# Go Quick Reference Cheat Sheet 🚀

## 📌 Keep This Open While Coding

---

## 🔧 Basic Syntax

### Variables & Types
```go
// Declaration
var name string = "Go"
age := 25                    // Short declaration
const PI = 3.14

// Common types
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
bool, string, byte, rune
```

### Arrays vs Slices
```go
// Array (fixed size)
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice (dynamic)
slice := []int{1, 2, 3}
slice = append(slice, 4)
subSlice := slice[1:3]       // [2, 3]
```

### Maps
```go
m := make(map[string]int)
m["key"] = 42
value, exists := m["key"]
delete(m, "key")
```

---

## 🔄 Control Flow

### Loops
```go
// Only "for" exists in Go
for i := 0; i < 10; i++ {}   // Classic
for i < 10 {}                // While-style
for {}                       // Infinite
for i, v := range slice {}   // Range
```

### If/Else
```go
if x > 0 {
    // ...
} else if x < 0 {
    // ...
} else {
    // ...
}

// With initialization
if val := getValue(); val > 0 {
    // val only exists here
}
```

### Switch
```go
switch day {
case "Mon":
    // No break needed!
case "Tue", "Wed":
    // Multiple cases
default:
    // ...
}
```

---

## 📦 Functions

### Basic Function
```go
func add(a, b int) int {
    return a + b
}

// Multiple returns
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named returns
func split(sum int) (x, y int) {
    x = sum / 2
    y = sum - x
    return  // Naked return
}
```

### Variadic Functions
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
sum(1, 2, 3, 4)
```

---

## 🎯 Pointers

```go
var p *int              // Pointer declaration
x := 42
p = &x                  // & = address of x
fmt.Println(*p)         // * = dereference (get value)

func modify(p *int) {
    *p = 100            // Changes original value
}
```

**Remember**: 
- `&` = "address of"
- `*` = "value at address"

---

## ⚡ Goroutines & Channels

### Goroutines
```go
go myFunction()         // Run in background

// Anonymous function
go func() {
    fmt.Println("Running async!")
}()
```

### Channels
```go
// Unbuffered (blocks until received)
ch := make(chan int)

// Buffered (holds 5 items)
ch := make(chan int, 5)

ch <- 42                // Send
value := <-ch           // Receive
close(ch)               // Close channel

// Check if closed
value, ok := <-ch
if !ok {
    // Channel closed
}
```

### Channel Patterns
```go
// Range over channel (until closed)
for value := range ch {
    fmt.Println(value)
}

// Select (like switch for channels)
select {
case msg := <-ch1:
    fmt.Println(msg)
case msg := <-ch2:
    fmt.Println(msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no message")
}
```

---

## 🔒 Synchronization

### WaitGroup
```go
var wg sync.WaitGroup

wg.Add(1)              // Add 1 to counter
go func() {
    defer wg.Done()    // Decrement when done
    // Do work
}()
wg.Wait()              // Block until counter is 0
```

### Mutex
```go
var mu sync.Mutex

mu.Lock()
// Critical section
mu.Unlock()

// Or use defer
mu.Lock()
defer mu.Unlock()
// Safe to use
```

---

## 🏗️ Structs & Methods

### Struct
```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
p := Person{"Bob", 25}             // Positional

// Anonymous struct
data := struct {
    ID   int
    Name string
}{ID: 1, Name: "Test"}
```

### Methods
```go
// Value receiver (copy)
func (p Person) SayHello() {
    fmt.Println("Hello", p.Name)
}

// Pointer receiver (modify)
func (p *Person) Birthday() {
    p.Age++
}
```

### Embedding (Composition)
```go
type Employee struct {
    Person              // Embedded
    Role string
}

e := Employee{
    Person: Person{Name: "Eve", Age: 28},
    Role: "Engineer",
}
e.SayHello()           // Can call Person methods
```

---

## 🎭 Interfaces

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Empty interface (any type)
var anything interface{}
anything = 42
anything = "string"

// Type assertion
str, ok := anything.(string)

// Type switch
switch v := anything.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
}
```

---

## ⚠️ Error Handling

```go
// Create error
err := errors.New("something failed")
err := fmt.Errorf("failed with code: %d", code)

// Check error
if err != nil {
    return err
}

// Wrap errors (Go 1.13+)
err = fmt.Errorf("detailed context: %w", originalErr)

// Custom error
type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
```

---

## 🕐 Context

```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Check if cancelled
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // Continue
}
```

---

## 📊 Common Patterns

### Worker Pool
```go
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start workers
for w := 1; w <= 5; w++ {
    go worker(jobs, results)
}

// Send jobs
for j := 1; j <= 100; j++ {
    jobs <- j
}
close(jobs)

// Collect results
for r := 1; r <= 100; r++ {
    <-results
}
```

### Producer-Consumer
```go
// Producer
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}()

// Consumer
for item := range ch {
    process(item)
}
```

### Fan-Out, Fan-In
```go
// Fan-out: Multiple workers read from same channel
// Fan-in: Multiple channels merge into one
func fanIn(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for {
            select {
            case v := <-ch1:
                out <- v
            case v := <-ch2:
                out <- v
            }
        }
    }()
    return out
}
```

---

## 🧪 Testing

```go
// File: math_test.go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}

// Table-driven tests
func TestAddTable(t *testing.T) {
    tests := []struct{
        a, b, want int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
    }
    
    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", 
                tt.a, tt.b, got, tt.want)
        }
    }
}

// Run: go test
// Coverage: go test -cover
```

---

## 🚨 Common Mistakes & How to Avoid

| Mistake | Why It's Bad | Fix |
|---------|-------------|-----|
| Forgetting `defer close(ch)` | Goroutine leak | Always close channels when done |
| Not checking `err != nil` | Silent failures | Always check errors immediately |
| Passing large structs by value | Slow, copies data | Use pointers for large structs |
| `go func()` in loop without passing vars | Race condition | Pass loop var as parameter |
| Using mutex without `defer unlock` | Deadlock | Always use `defer mu.Unlock()` |
| Modifying slice during range | Undefined behavior | Use index-based loop instead |

---

## 🎯 Big-O Complexity Quick Reference

| Operation | Array/Slice | Map | Channel |
|-----------|-------------|-----|---------|
| Access by index | O(1) | - | - |
| Search | O(n) | - | - |
| Insert | O(n) | O(1) avg | O(1) |
| Delete | O(n) | O(1) avg | - |
| Iterate | O(n) | O(n) | O(n) |

---

## 🔍 Debugging Commands

```bash
# Run program
go run main.go

# Build binary
go build

# Format code
go fmt ./...

# Run tests
go test
go test -v              # Verbose
go test -cover          # Coverage
go test -bench=.        # Benchmarks

# Show dependencies
go mod tidy
go mod graph

# Profile
go build -o myapp
./myapp -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Race detection
go run -race main.go
```

---

## 💡 Interview Quick Tips

1. **Always initialize channels and maps**
   - Channels: `make(chan int)`
   - Maps: `make(map[string]int)`

2. **Nil vs Empty**
   - `var s []int` → `s == nil` (true)
   - `s := []int{}` → `s == nil` (false)

3. **Concurrency mantra**: "Share memory by communicating, don't communicate by sharing memory"

4. **When to use buffered channels**: 
   - You know exact capacity needed
   - Sender shouldn't block immediately
   - Prevents goroutine leaks

5. **Pointer receiver when**:
   - Method modifies the receiver
   - Receiver is large struct
   - Consistency (if one method uses pointer, all should)

---

## 📚 Must-Know Standard Packages

```go
import (
    "fmt"           // Formatting & printing
    "errors"        // Error creation
    "strings"       // String manipulation
    "strconv"       // String conversion
    "time"          // Time operations
    "sync"          // Synchronization primitives
    "context"       // Cancellation & deadlines
    "encoding/json" // JSON parsing
    "net/http"      // HTTP client/server
    "io/ioutil"     // Deprecated: use io & os instead
    "os"            // OS operations
)
```

---

## 🎓 Remember

- **KISS**: Keep It Simple, Stupid
- **DRY**: Don't Repeat Yourself  
- **Prefer composition over inheritance** (Go has no inheritance!)
- **Make the zero value useful** (design structs that work without initialization)
- **Errors are values** (not exceptions!)

---

**Print this out or keep it in a second monitor. Refer to it constantly!**
