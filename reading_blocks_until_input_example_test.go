package chinfbuf_test

import (
	"github.com/GodsBoss/chinfbuf"

	"fmt"
	"time"
)

func Example_readingBlocksUntilThereIsInput() {
	input, output := chinfbuf.New()

	go func() {
		// We sleep a *whole* millisecond here.
		time.Sleep(time.Millisecond)
		input <- "Hello, gopher!"
	}()

	fmt.Println(<-output)

	// Output:
	// Hello, gopher!
}
