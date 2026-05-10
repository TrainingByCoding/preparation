# Day 35: Pipeline Pattern

## 🎯 Goal
Build **channel pipelines** — stages connected by channels where each stage transforms data. The Go way of stream processing.

## 🧠 Key Concepts
- Pipeline = chain of stages, each receives from previous stage, transforms, sends to next
- Each stage runs in its own goroutine → all stages run **concurrently**
- Data flows through channels between stages
- Cancellation: use `done` channel or `context.Context` to stop all stages
- Real-world use: ETL pipelines, image processing, log processing, API request chains

## 📖 Pattern

```
generate() ──► square() ──► filter() ──► print
```

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums { out <- n }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in { out <- n * n }
        close(out)
    }()
    return out
}

func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if n%2 == 0 { out <- n }
        }
        close(out)
    }()
    return out
}

// Chain them:
nums   := generate(1, 2, 3, 4, 5)
squared := square(nums)
evens  := filterEven(squared)
for v := range evens { fmt.Println(v) }
```

## ✅ Learning Checklist
- [ ] Implement a 3-stage pipeline (generate → transform → filter)
- [ ] Understand each stage runs concurrently as data flows
- [ ] Pipeline with context cancellation (stop mid-stream)
- [ ] Batch/buffer pipeline (collect N items before passing downstream)

## 🛠️ Practice Exercises
1. **Number pipeline** — generate 1-10 → square → filter evens → print
2. **String pipeline** — generate words → uppercase → filter len > 4 → print
3. **CSV pipeline** — simulate reading CSV rows → parse → validate → write to output
4. **Cancellable pipeline** — generate infinite numbers → filter primes → stop after 5 primes found
5. **BONUS** — Add a `done` channel so consumer can cancel producer mid-stream

## 🔍 Interview Questions
- "How do you build a streaming data processor in Go?"
- "How do pipelines differ from just calling functions in sequence?"
- "How would you add cancellation to a pipeline?"
