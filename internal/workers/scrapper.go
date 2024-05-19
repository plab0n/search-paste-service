package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/plab0n/search-paste/pkg/workerutils"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

type Scrapper struct {
	Worker
}

var b *bus.MessageBus

func (c *Scrapper) Start() error {
	b = bus.New()
	topic := workerutils.PasteCrawlTopic()
	err := b.SubscribeWithHandler(topic, ScrapUrl)
	return err
}

func ScrapUrl(message interface{}) error {
	if cm, ok := message.(model.CrawlerMessage); ok {
		htmlContent, err := fetchContent(cm.Url)
		if err != nil {
			logger.Log.Error(err)
			return err
		}
		plainText, err := html.Parse(strings.NewReader(htmlContent))
		if err != nil {
			return err
		}
		b.Publish(workerutils.PasteIndexerTopic(), plainText)
	}
	return nil
}

func fetchContent(url string) (string, error) {
	//TODO: Check response error codes
	response, err := http.Get(url)

	if err != nil {
		logger.Log.Error(err)
		return "", err

	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		logger.Log.Error(err)
		return "", err
	}
	htmlContent := string(body)
	return htmlContent, nil
}
