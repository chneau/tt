package tt

import (
	"io"
	logger "log"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	mu  sync.RWMutex
	log = logger.New(os.Stdout, "[TRACK] ", logger.LstdFlags)
)

// SetOutput sets the output destination for the standard logger in tt.
// It is safe for concurrent use.
func SetOutput(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	log.SetOutput(w)
}

// SetFlags sets the output flags for the logger in tt.
// It is safe for concurrent use.
func SetFlags(flag int) {
	mu.Lock()
	defer mu.Unlock()
	log.SetFlags(flag)
}

// Track allows you to track timing of a function.
// Just add:
//
//	defer tt.Track(time.Now(), "nameOfFunction")
//
// at the top of your function.
func Track(start time.Time, name string) {
	mu.RLock()
	defer mu.RUnlock()
	log.Printf("<[%s]> %s", name, time.Since(start))
}

// T allows you to track execution time of a function.
// It automatically retrieves the caller's function name using runtime.
// Just add:
//
//	defer tt.T()()
//
// at the top of your function.
func T() func() {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name() // nil is ok, see impl of Name()
	start := time.Now()
	return func() {
		Track(start, name)
	}
}
