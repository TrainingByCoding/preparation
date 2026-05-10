package main

import (
	"fmt"
	"os"
	"sync"
)

// ========================================
// Exercise 1: Thread-safe DB Singleton
// ========================================

type Database struct {
	connection string
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

// TODO: Implement GetDB() — returns same *Database every time
// Use sync.Once to make it thread-safe
// Print "DB connection created" inside the init func
func GetDB() *Database {
	// TODO: implement
	return nil
}

// ========================================
// Exercise 2: Logger Singleton
// ========================================

type Logger struct {
	prefix string
}

var (
	logInstance *Logger
	logOnce     sync.Once
)

// TODO: Implement GetLogger() — prefix should be "[APP]"
func GetLogger() *Logger {
	// TODO: implement
	return nil
}

func (l *Logger) Log(msg string) {
	fmt.Printf("%s %s\n", l.prefix, msg)
}

// ========================================
// Exercise 3: Concurrent proof
// ========================================

func exercise3() {
	var wg sync.WaitGroup
	// TODO: Launch 10 goroutines, each calls GetDB()
	// Each should print its goroutine ID and the pointer address
	// All addresses should be identical
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// TODO: get instance and print fmt.Printf("goroutine %d: %p\n", id, db)
		}(i)
	}
	wg.Wait()
}

// ========================================
// Exercise 4: Config Singleton
// ========================================

type Config struct {
	AppEnv string
	Debug  bool
}

var (
	configInstance *Config
	configOnce     sync.Once
)

// TODO: Implement GetConfig()
// Read APP_ENV from os.Getenv("APP_ENV"), default to "development"
// Set Debug = true if AppEnv == "development"
func GetConfig() *Config {
	// TODO: implement
	_ = os.Getenv // hint
	return nil
}

// ========================================
// Main
// ========================================

func main() {
	fmt.Println("=== Day 31: Singleton Pattern ===\n")

	fmt.Println("Exercise 1: DB Singleton")
	db1 := GetDB()
	db2 := GetDB()
	fmt.Printf("db1 addr: %p\n", db1)
	fmt.Printf("db2 addr: %p\n", db2)
	fmt.Printf("Same instance: %v\n", db1 == db2)

	fmt.Println("\nExercise 2: Logger Singleton")
	logger := GetLogger()
	if logger != nil {
		logger.Log("Application started")
		logger.Log("Another message")
	}

	fmt.Println("\nExercise 3: Concurrent goroutines")
	exercise3()

	fmt.Println("\nExercise 4: Config Singleton")
	cfg := GetConfig()
	if cfg != nil {
		fmt.Printf("Env: %s, Debug: %v\n", cfg.AppEnv, cfg.Debug)
	}
}

/*
SOLUTIONS:

// Exercise 1:
func GetDB() *Database {
	dbOnce.Do(func() {
		dbInstance = &Database{connection: "postgres://localhost/mydb"}
		fmt.Println("DB connection created")
	})
	return dbInstance
}

// Exercise 2:
func GetLogger() *Logger {
	logOnce.Do(func() {
		logInstance = &Logger{prefix: "[APP]"}
	})
	return logInstance
}

// Exercise 3:
go func(id int) {
	defer wg.Done()
	db := GetDB()
	fmt.Printf("goroutine %d: %p\n", id, db)
}(i)

// Exercise 4:
func GetConfig() *Config {
	configOnce.Do(func() {
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}
		configInstance = &Config{
			AppEnv: env,
			Debug:  env == "development",
		}
	})
	return configInstance
}
*/
