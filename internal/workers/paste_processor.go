package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	"time"
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
	ch, err := b.Subscribe(topic)

	if err != nil {
		return err
	}
	go func() {
		for i := 0; ; i++ {
			// Send data to the channel
			message := <-ch
			if paste, ok := message.(model.Paste); ok {
				if isUrl(paste.Text) {
					//Crawl the url
				} else {
					//Create embedding
				}
			}
			time.Sleep(time.Second) // Simulate some delay
		}
	}()
	return nil
}

func isUrl(text string) bool {
	return false
}
