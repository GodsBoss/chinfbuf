package chinfbuf_test

import (
	"github.com/GodsBoss/chinfbuf"

	"testing"
	"time"
)

func TestDirectlyClosingInputChannelClosesOutputChannel(t *testing.T) {
	input, output := chinfbuf.New()
	close(input)

	timer := time.NewTimer(time.Millisecond)
	select {
	case val, ok := <-output:
		if ok {
			t.Errorf("Expected output not to return a value, but got %+v", val)
		}
	case <-timer.C:
		t.Errorf("Timeout reached when reading from output channel")
	}
}
