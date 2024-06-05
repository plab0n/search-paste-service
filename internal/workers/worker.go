package workers

import "github.com/plab0n/search-paste/internal/bus"

type Worker interface {
	Start() error
}

type BaseWorker struct {
	B     bus.Bus
	Topic string
}
