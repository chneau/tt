# tt

A simple, thread-safe execution time tracker in Go in ~50 lines of code.

[![Go Reference](https://pkg.go.dev/badge/github.com/chneau/tt.svg)](https://pkg.go.dev/github.com/chneau/tt)
[![Go Report Card](https://goreportcard.com/badge/github.com/chneau/tt)](https://goreportcard.com/report/github.com/chneau/tt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## ⚡ Features

- **Zero Friction**: One line instrumentation via `defer tt.T()()`.
- **Automatic Function Name Detection**: Leverages `runtime.Caller` to infer caller name automatically.
- **Thread-Safe**: Fully safe for concurrent calls and log configuration updates.
- **Configurable Output**: Direct logs to any `io.Writer` via `tt.SetOutput(w)` or adjust log flags with `tt.SetFlags(flag)`.

---

## 🚀 Installation

```bash
go get github.com/chneau/tt
```

---

## 💡 Usage

Simply add `defer tt.T()()` at the top of your function:

```go
package main

import (
	"time"

	"github.com/chneau/tt"
)

func process() {
	defer tt.T()() // Automatically prints function name and duration upon return

	time.Sleep(50 * time.Millisecond)
}

func customNamed() {
	defer tt.Track(time.Now(), "myCustomOperation")

	time.Sleep(25 * time.Millisecond)
}

func main() {
	process()
	customNamed()
}
```

### Example Output

```text
[TRACK] 2026/09/04 13:36:00 <[main.process]> 50.123456ms
[TRACK] 2026/09/04 13:36:00 <[myCustomOperation]> 25.045612ms
```

---

## 🛠️ Development Commands

```bash
# Run tests with race detector and coverage
go test -v -race -cover ./...

# Run benchmarks
go test -benchmem -bench=. ./...

# Modernize codebase
modernize ./...

# Run linter
golangci-lint run ./...
```

---

## 📄 License

[MIT](LICENSE)
