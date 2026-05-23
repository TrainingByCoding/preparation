# Memory Leak example - pprof 

**Go does NOT have classic memory leaks** like C/C++.

What people _actually_ mean is **unbounded memory growth** due to:

-   keeping references alive
    
-   global caches/maps
    
-   goroutines holding memory
    
-   slices growing and never released
___

## 1️⃣ Leaky Go Program (pprof enabled)

# `main.go`
``` .go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// GLOBAL slice → this is the "leak"
var leakyGlobal [][]byte

func main() {
	// Start pprof server
	go func() {
		fmt.Println("pprof running on :6060")
		http.ListenAndServe(":6060", nil)
	}()

	// Keep allocating memory forever
	for {
		allocate()
		time.Sleep(200 * time.Millisecond)
	}
}

func allocate() {
	// Allocate 1MB
	b := make([]byte, 1024*1024)

	// Store reference globally → GC can NEVER free it
	leakyGlobal = append(leakyGlobal, b)

	fmt.Println("Allocated MBs:", len(leakyGlobal))
}

___

## 2️⃣ Run the Program
go run .\pprof.go
You’ll see: Memory will keep climbing.

Allocated MBs: 1
Allocated MBs: 2
Allocated MBs: 3
.
.
.
___

## 3️⃣ Capture Memory Profile

In another terminal: go tool pprof http://localhost:6060/debug/pprof/heap
___

## 4️⃣ Commands to Clearly Identify the Leak
(pprof) top
Showing nodes accounting for 61.30MB, 100% of 61.30MB total
      flat  flat%   sum%        cum   cum%
   61.30MB   100%   100%    61.30MB   100%  main.allocate
         0     0%   100%    61.30MB   100%  main.main
         0     0%   100%    61.30MB   100%  runtime.main
___

### 🔥 Line-level blame
(pprof) list allocate
Total: 61.30MB
ROUTINE ======================== main.allocate in C:\my_data\go_practice\pprof.go        
   61.30MB    61.30MB (flat, cum)   100% of Total
         .          .     27:func allocate() {
         .          .     28:   // Allocate 1MB
   61.30MB    61.30MB     29:   b := make([]byte, 1024*1024)
         .          .     30:
         .          .     31:   // Store reference globally → GC can NEVER free it       
         .          .     32:   leakyGlobal = append(leakyGlobal, b)
         .          .     33:
         .          .     34:   fmt.Println("Allocated MBs:", len(leakyGlobal))

The global slice is holding references forever.
