package workers

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/plab0n/search-paste/pkg/workerutils"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type WorkerHandler struct {
	Bus *bus.MessageBus
}

func (h *WorkerHandler) NewPasteHandler(message interface{}) error {
	if paste, ok := message.(model.Paste); ok {
		if ok, _ := isUrl(paste.Text); ok {
			cm := &model.CrawlerMessage{PasteId: paste.ID, Url: paste.Text}
			err := h.Bus.Publish(workerutils.PasteCrawlTopic(), cm)
			if err != nil {
				logger.Log.Error(err)
			}
		} else {
			//Create embedding
		}
	}
	return nil
}
func isUrl(u string) (bool, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false, err
	}
	return parsedUrl.Scheme == "https" && parsedUrl.Host != "", nil
}
func (h *WorkerHandler) ScrapeUrl(message interface{}) error {
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
		h.Bus.Publish(workerutils.PasteIndexerTopic(), plainText)
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
