package receiver

import (
	"sync"
	"time"
)

type Receiver struct {
	doneAt time.Time
	mu     sync.Mutex
}

func (r *Receiver) Done() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.doneAt = time.Now()
}

func (r *Receiver) DoneAt() time.Time {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.doneAt
}
