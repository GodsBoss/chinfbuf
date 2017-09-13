package chinfbuf

// New creates a pair of an input and an output channel which are connected and
// have a storage of infinite size between them.
func New() (chan<- interface{}, <-chan interface{}) {
	buf := &buffer{
		values: []interface{}{},
		input:  make(chan interface{}),
		output: make(chan interface{}),
	}
	buf.state = buf.waitForCloseOrWrite
	go buf.run()
	return buf.input, buf.output
}

type buffer struct {
	input    chan interface{}
	output   chan interface{}
	values   []interface{}
	state    func()
	finished bool
}

func (b *buffer) run() {
	for !b.finished {
		b.state()
	}
}

func (b *buffer) waitForCloseOrWrite() {
	value, ok := <-b.input
	if !ok {
		b.state = b.end
		return
	}
	b.values = append(b.values, value)
	b.state = b.waitForCloseWriteOrRead
}

func (b *buffer) waitForCloseWriteOrRead() {
	select {
	case value, open := <-b.input:
		if open {
			b.values = append(b.values, value)
		} else {
			b.state = b.waitForRead
		}
	case b.output <- b.values[0]:
		b.values = b.values[1:]
		if len(b.values) == 0 {
			b.state = b.waitForCloseOrWrite
		}
	}
}

func (b *buffer) waitForRead() {
	b.output <- b.values[0]
	b.values = b.values[1:]
	if len(b.values) == 0 {
		b.state = b.end
	}
}

func (b *buffer) end() {
	b.finished = true
	close(b.output)
}
