# Day 31: Singleton Pattern

## 🎯 Goal
Implement a **thread-safe Singleton** in Go — one of the most common design pattern interview questions.

## 🧠 Key Concepts
- Singleton = only ONE instance exists for the entire program lifetime
- Naive implementation has a **race condition** when multiple goroutines call it simultaneously
- `sync.Once` guarantees the init function runs **exactly once**, even under concurrent access
- Common real-world uses: DB connection pool, logger, config loader, metric registry

## 📖 Pattern

### ❌ Naive (Race Condition!)
```go
var instance *Database

func GetDB() *Database {
    if instance == nil {       // goroutine A and B both see nil at same time
        instance = &Database{} // both create a new instance → race!
    }
    return instance
}
```

### ✅ Thread-Safe with sync.Once
```go
type Database struct{ connection string }

var (
    dbInstance *Database
    dbOnce     sync.Once
)

func GetDB() *Database {
    dbOnce.Do(func() {
        dbInstance = &Database{connection: "postgres://localhost/mydb"}
        fmt.Println("DB initialized (only printed once!)")
    })
    return dbInstance
}
```

`sync.Once` internally uses a mutex — the init func runs once, then it's a no-op forever.

## ✅ Learning Checklist
- [ ] Can explain WHY naive singleton is not thread-safe
- [ ] Implement singleton with `sync.Once`
- [ ] Verify same pointer returned from 10 concurrent goroutines
- [ ] Know real use cases: DB pool, logger, config

## 🛠️ Practice Exercises
1. **DB Singleton** — implement `GetDB()`, verify `db1 == db2` is true
2. **Logger Singleton** — singleton logger with a prefix `[APP]`
3. **Concurrent proof** — 10 goroutines call `GetDB()`, all print same pointer address
4. **Config Singleton** — read `APP_ENV` env var once, return same config struct
5. **BONUS** — Why is singleton sometimes an anti-pattern? When should you use dependency injection instead?

## 🔍 Interview Questions
- "How do you implement a thread-safe singleton in Go?"
- "What is sync.Once and when would you use it?"
- "Can you reset a sync.Once?" (No — use pointer swap pattern for testing)
