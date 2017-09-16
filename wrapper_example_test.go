package chinfbuf_test

import (
	"github.com/GodsBoss/chinfbuf"

	"fmt"
)

func newIntWrapper() (chan<- int, <-chan int) {
	input, output := chinfbuf.New()

	intInput := make(chan int)
	intOutput := make(chan int)

	go func() {
		for {
			i, ok := <-intInput
			if !ok {
				close(input)
				return
			}
			input <- i
		}
	}()

	go func() {
		for {
			v, ok := <-output
			if !ok {
				close(intOutput)
				return
			}
			intOutput <- v.(int)
		}
	}()

	return intInput, intOutput
}

// For more type safety, it is advisable to create a wrapper which converts the
// interface{} values going into and coming from the channel to more specific
// types.
func Example_typedWrapper() {
	input, output := newIntWrapper()

	for i := 1; i <= 100; i++ {
		input <- i
	}

	close(input)

	sum := 0

	for v := range output {
		sum = sum + v
	}

	fmt.Printf("Sum of the numbers from 1 to 100 is %d.\n", sum)

	// Output:
	// Sum of the numbers from 1 to 100 is 5050.
}
