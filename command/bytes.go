package command

import (
	"bytes"
	"sync"
)

// BufferOut holds the bytes of every output command to the stdout
type BufferOut struct {
	stdout bytes.Buffer
	mutex  sync.Mutex
}

// Reset: sets a stdout buffer to an empty buffer.
func (b *BufferOut) Reset() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.stdout.Reset()
}

// BufferErr holds the bytes of every error output command to the stderr
type BufferErr struct {
	stderr bytes.Buffer
	mutex  sync.Mutex
}

// Reset: sets a stderr buffer to an empty buffer.
func (b *BufferErr) Reset() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.stderr.Reset()
}
