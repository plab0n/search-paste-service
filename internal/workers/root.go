package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Root struct {
	BaseWorker
}

func init() {
	b := bus.New()

	c := &Scrapper{}
	c.B = b
	c.Start()

	e := &Embedder{}
	e.B = b
	e.Start()

	p := &Root{}
	p.B = b
	p.Start()
}

func (p *Root) Start() error {
	topic := workerutils.PasteCreatedTopic()
	h := &workers.WorkerHandler{Bus: p.B}
	err := p.B.SubscribeWithHandler(topic, h.NewPasteHandler)
	return err
}
