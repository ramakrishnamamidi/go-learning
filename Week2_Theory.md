# Go Lang Week 2 — Theory & Concepts

## 1. Structs and Methods
Structs allow grouping related data fields. Methods defined with a receiver provide behavior to struct types.

Key Points:
- `struct` defines a custom data type.
- Methods can have **value** or **pointer** receivers.
- Value receiver → works on a copy (no modification).
- Pointer receiver → modifies the original instance.

Example Reference: `examples/structs/main.go`

---

## 2. Pointers
Pointers hold memory addresses of variables. They allow modifying original values passed to functions.

Key Points:
- `&x` gives address.
- `*p` dereferences pointer.
- Struct pointer receivers update data directly.

Example Reference: `examples/pointers/main.go`

---

## 3. Interfaces
Interfaces define **behavior (methods)** not data. Any type implementing its methods satisfies the interface.

Key Points:
- Interfaces support **polymorphism**.
- Calling methods without knowing exact type.

Example Reference: `examples/interfaces/main.go`

---

## 4. Error Handling
Go returns errors as values, not exceptions.

Key Points:
- Functions return `(value, error)`
- Always check `err != nil`
- `errors.New()` for custom error

Example Reference: `examples/errors/main.go`

---

## 5. File I/O
Reading and writing files using standard library packages.

Key Points:
- `os.WriteFile()` to write
- `os.ReadFile()` to read
- Always handle errors
- Remove temp files using `os.Remove()`

Example Reference: `examples/fileio/main.go`

---

## 6. Goroutines (Concurrency)
Goroutines are lightweight threads managed by Go runtime.

Key Points:
- Start using `go functionCall()`
- Run concurrently with `main` function
- Prevent premature exit with `time.Sleep()`, WaitGroups later

Example Reference: `examples/concurrency/main.go`

---

## 7. Channels
Channels allow safe communication between goroutines.

Key Points:
- `make(chan int)` creates a channel
- `<-` operator sends and receives
- `range ch` loops until closed

Example Reference: `examples/channels/main.go`

---

## 8. CLI Flags
Command line argument handling using `flag` package.

Key Points:
- `flag.String()` and `flag.Int()`
- `flag.Parse()` required
- Dereference: `*flagVar`

Example Reference: `cmd/greetflags/main.go`

---

## Summary of Week 2
- Structs, methods, and pointers
- Interfaces & polymorphism
- Idiomatic error handling
- File operations
- Concurrency using goroutines and channels
- CLI applications using flags

**Repo folder references:**
- `examples/structs` → structs
- `examples/pointers` → pointers
- `examples/interfaces` → interfaces and polymorphism.
- `examples/errors` → error handling
- `examples/fileio` → file operations
- `examples/concurrency` → goroutines for concurrency
- `examples/channels` → channels usage in concurrency
- `cmd/greetflags` → command-line flag parsing.

---

**End of Week 2 Theory**
