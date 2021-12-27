package command

import (
	"bytes"
	"sync"
)

type BufferOut struct {
	stdout bytes.Buffer
	mutex  sync.Mutex
}

func (b *BufferOut) Reset() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.stdout.Reset()
}

type BufferErr struct {
	stderr bytes.Buffer
	mutex  sync.Mutex
}

func (b *BufferErr) Reset() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.stderr.Reset()
}
