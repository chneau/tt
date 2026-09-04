package tt

import (
	"bytes"
	"io"
	"strings"
	"sync"
	"testing"
	"time"
)

func helperFunction() {
	defer T()()
	time.Sleep(10 * time.Millisecond)
}

func TestTrackAndT(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetFlags(0)
	defer SetOutput(io.Discard)

	start := time.Now().Add(-50 * time.Millisecond)
	Track(start, "customTask")

	out := buf.String()
	if !strings.Contains(out, "<[customTask]>") {
		t.Fatalf("expected log to contain '<[customTask]>', got: %q", out)
	}

	buf.Reset()
	helperFunction()

	out = buf.String()
	if !strings.Contains(out, "helperFunction") {
		t.Fatalf("expected log to contain 'helperFunction', got: %q", out)
	}
}

func TestConcurrency(t *testing.T) {
	SetOutput(io.Discard)

	const goroutines = 16
	const iterations = 500

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			for range iterations {
				defer T()()
				Track(time.Now(), "concurrentOp")
			}
		}()
	}
	wg.Wait()
}

// ====================================================================
// Standard Benchmarks
// ====================================================================

func BenchmarkTrack(b *testing.B) {
	SetOutput(io.Discard)
	now := time.Now()
	b.ReportAllocs()
	for b.Loop() {
		Track(now, "benchOp")
	}
}

func BenchmarkT(b *testing.B) {
	SetOutput(io.Discard)
	b.ReportAllocs()
	for b.Loop() {
		done := T()
		done()
	}
}
