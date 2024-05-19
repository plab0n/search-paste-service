package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	"github.com/plab0n/search-paste/pkg/logger"
	"net/url"
)

type Root struct {
	Worker
}

var topic string

func Init() {
	topic = "paste.Created"
	p := &Root{}
	p.Start()
}

func (p *Root) Start() error {
	b := bus.New()
	err := b.SubscribeWithHandler(topic, func(message interface{}) error {
		if paste, ok := message.(model.Paste); ok {
			if ok, _ := isUrl(paste.Text); ok {
				err := b.Publish("paste.crawl", paste.Text)
				if err != nil {
					logger.Log.Error(err)
				}
			} else {
				//Create embedding
			}
		}
		return nil
	})
	return err
}

func isUrl(u string) (bool, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false, err
	}
	return parsedUrl.Scheme == "https" && parsedUrl.Host != "", nil
}
