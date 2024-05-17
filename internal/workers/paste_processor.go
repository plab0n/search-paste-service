package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
)

type PasteProcessor struct {
	Worker
}

var topic string

func Init() {
	topic = "paste.Created"
	p := &PasteProcessor{}
	p.Start()
}

func (p *PasteProcessor) Start() error {
	b := bus.New()
	err := b.SubscribeWithHandler(topic, func(message interface{}) error {
		if paste, ok := message.(model.Paste); ok {
			if isUrl(paste.Text) {
				//Crawl the url
			} else {
				//Create embedding
			}
		}
		return nil
	})
	return err
}

func isUrl(text string) bool {
	return false
}
