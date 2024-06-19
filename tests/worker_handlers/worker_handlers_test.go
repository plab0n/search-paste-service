package worker_handlers_test

import (
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	workers "github.com/plab0n/search-paste/internal/workers/handlers"
	"testing"
)

func Test_Scrapper(t *testing.T) {
	h := &workers.WorkerHandler{Bus: bus.New()}
	err := h.Scrapper(model.ScrapingInfo{PasteId: 1, Url: "https://programmer.ink/think/go-scheduler-series-source-code-reading-and-exploration.html"})
	if err != nil {
		t.Error(err)
	}
}

func Test_Embedder(t *testing.T) {
	h := &workers.WorkerHandler{Bus: bus.New()}
	err := h.EmbeddingHandler("Hi! Please embed me")
	if err != nil {
		t.Error(err)
	}
}
