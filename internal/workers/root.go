package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Root struct {
	Worker
}

func init() {
	c := &Scrapper{}
	c.Start()

	p := &Root{}
	p.Start()

	e := &Embedder{}
	e.Start()
}

func (p *Root) Start() error {
	b := bus.New()
	topic := workerutils.PasteCreatedTopic()
	h := &workers.WorkerHandler{Bus: b}
	err := b.SubscribeWithHandler(topic, h.NewPasteHandler)
	return err
}
