package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/vector_storage"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"github.com/plab0n/search-paste/pkg/workerutils"
)

type Root struct {
	BaseWorker
}

func init() {
	b := bus.New()
	es, _ := vector_storage.NewElasticDb()
	c := &Scrapper{}
	c.B = b
	c.Start()

	e := &Embedder{}
	e.B = b
	e.Start()

	i := &Indexer{}
	i.B = b
	i.VectorStorage = es
	i.Start()
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
