package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/plab0n/search-paste/pkg/workerutils"
	"net/url"
)

type Root struct {
	Worker
}

func Init() {
	c := &Scrapper{}
	c.Start()

	p := &Root{}
	p.Start()
}

func (p *Root) Start() error {
	b := bus.New()
	topic := workerutils.PasteCreatedTopic()
	err := b.SubscribeWithHandler(topic, func(message interface{}) error {
		if paste, ok := message.(model.Paste); ok {
			if ok, _ := isUrl(paste.Text); ok {
				cm := &model.CrawlerMessage{PasteId: paste.ID, Url: paste.Text}
				err := b.Publish(workerutils.PasteCrawlTopic(), cm)
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
