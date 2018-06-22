package tt

import (
	"log"
	"time"
)

// Track allow you to track timing of a function
// Just add defer tt.Track(time.New, "nameOfFunction")
// at the top of your function.
func Track(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
