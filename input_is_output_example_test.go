package chinfbuf_test

import (
	"github.com/GodsBoss/chinfbuf"

	"fmt"
)

// All values pushed into the input channel can be pulled from the output channel.
// In addition, if the input channel is closed, the output channel will also
// be closed after the chinfbuf Buffer is drained.
func Example_inputIsOutput() {
	values := []string{"foo", "bar", "baz"}

	input, output := chinfbuf.New()

	for _, value := range values {
		input <- value
	}

	close(input)

	for value := range output {
		fmt.Println(value)
	}

	// Output:
	// foo
	// bar
	// baz
}
