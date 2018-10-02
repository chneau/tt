package tt

import (
	logger "log"
	"os"
	"runtime"
	"time"
)

var log = logger.New(os.Stdout, "[TRACK] ", logger.LstdFlags)

// Track allows you to track timing of a function
// Just add defer tt.Track(time.New, "nameOfFunction")
// at the top of your function.
func Track(start time.Time, name string) {
	log.Printf("<[%s]> %s", name, time.Since(start))
}

// T allows you to track time of a function
// It will get the name of the function via runtime
// It will print the time used by the function
func T() func() {
	name := "N/A"
	{
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			name = details.Name()
		}
	}
	start := time.Now()
	return func() {
		Track(start, name)
	}
}
