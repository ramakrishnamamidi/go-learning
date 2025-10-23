# Go Lang Week 1 — Theory & Concepts

This document covers all concepts and theory for **Week 1 Go Lang learning**, aligned with the code examples in this repository.

---

## 1. Introduction to Go (Golang)

Go is a **compiled, statically typed, concurrent programming language** developed by Google.

**Key Features:**
- **Simplicity:** Minimalistic syntax.
- **Fast compilation:** Compiled to machine code.
- **Concurrency:** Built-in goroutines & channels.
- **Strong typing:** Compile-time type safety.
- **Cross-platform:** Windows, Linux, MacOS.
- **Automatic memory management:** Garbage collection.

**Use Cases:** Web servers, CLI tools, cloud services, DevOps, networking.

---

## 2. Go Toolchain

**Installation:** Check version:
```bash
go version
```

**Initialize a module:**
```bash
go mod init github.com/<username>/go-learning
```

**Common Commands:**
| Command | Description |
|---------|-------------|
| `go run main.go` | Compile & run without creating binary |
| `go build` | Build executable binary |
| `go install` | Install binary to `$GOBIN` |
| `go fmt` | Format code automatically |
| `go vet` | Static code analysis |
| `go test` | Run unit tests |
| `go mod tidy` | Clean unused dependencies |

---

## 3. Go Program Structure

Every Go program has:
1. **Package Declaration**
```go
package main
```
- `main` → required for executable programs.
- Libraries → custom package names.

2. **Imports**
```go
import "fmt"
```
- Import standard library or other packages.

3. **Main Function**
```go
func main() {
    fmt.Println("Hello, Go!")
}
```
- Entry point for execution.

---

## 4. Variables, Constants, Types

**Variables**
```go
var name string = "Alice"
age := 25
```
- `var` → explicit type.
- `:=` → type inference.

**Constants**
```go
const country = "India"
```
- Immutable value.

**Basic Types:**
- Integer: `int`, `int8`, `int64`
- Floating point: `float32`, `float64`
- Boolean: `bool`
- String: `string`

**Code reference:** [`examples/basics/main.go`](./examples/basics/main.go)

---

## 5. Control Flow

**If / Else**
```go
if score >= 90 { ... } else if score >= 80 { ... } else { ... }
```

**For Loop**
```go
for i := 1; i <= 5; i++ { ... }
```

**Switch**
```go
switch day {
case "Monday": ...
default: ...
}
```

**Code reference:** [`examples/flow/main.go`](./examples/flow/main.go)

---

## 6. Arrays, Slices, Maps

**Array (fixed size)**
```go
var arr = [3]int{10, 20, 30}
```

**Slice (dynamic size)**
```go
nums := []int{1, 2, 3}
nums = append(nums, 4)
```

**Map (key-value store)**
```go
ages := map[string]int{"Alice": 25, "Bob": 30}
ages["Charlie"] = 28
```

**Code reference:** [`examples/collections/main.go`](./examples/collections/main.go)

---

## 7. Functions

```go
func add(a int, b int) int { return a + b }

func divide(a, b float64) (float64, error) {
    if b == 0 { return 0, fmt.Errorf("division by zero") }
    return a / b, nil
}
```

- Functions can return **multiple values**.
- Idiomatic error handling uses `error` type.

**Code reference:** [`examples/functions/main.go`](./examples/functions/main.go)

---

## 8. Packages

- Organize code into modules for **reusability**.
- Example: `pkg/greet/greet.go`
```go
package greet

import "fmt"

func Hello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```
- Exported functions → capitalized names.
- Use in main program:
```go
import "github.com/<username>/go-learning/pkg/greet"
```

**Code reference:** [`pkg/greet/greet.go`](./pkg/greet/greet.go)

---

## 9. CLI Input/Output

**Input**
```go
reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')
```

**Output**
```go
fmt.Print("Message")
fmt.Println("Message")
fmt.Printf("Formatted %s", var)
```

**Code reference:** [`cmd/greetcli/main.go`](./cmd/greetcli/main.go)

---

## 10. Code Formatting & Static Analysis

```bash
gofmt -w .
go vet ./...
```
- `gofmt` → auto-format.
- `go vet` → detect suspicious code patterns.

---

## 11. Git Workflow for Learning Repo

- Commit **incrementally**:
```bash
git add .
git commit -m "feat: add hello world"
```
- Push to GitHub after completing step 9.

---

## 12. Summary of Week 1

- Go toolchain & workspace setup.
- Writing basic Go programs.
- Variables, constants, types, control flow.
- Arrays, slices, maps.
- Functions & multiple returns.
- Packages & modules.
- CLI input/output.
- Code formatting & vetting.
- Incremental Git commits for professional repo.

**Repo folder references:**
- `cmd/hello` → Hello World
- `examples/basics` → variables/types
- `examples/flow` → control flow
- `examples/collections` → arrays/slices/maps
- `examples/functions` → functions
- `pkg/greet` → custom package
- `cmd/greetcli` → CLI input/output

---

**End of Week 1 Theory**

