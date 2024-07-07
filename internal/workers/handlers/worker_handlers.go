package workers

import (
	"fmt"
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/model"
	"github.com/plab0n/search-paste/pkg/helpers"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/plab0n/search-paste/pkg/workerutils"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type WorkerHandler struct {
	Bus bus.Bus
}

func (h *WorkerHandler) NewPasteHandler(message interface{}) error {
	if paste, ok := message.(model.Paste); ok {
		if ok, _ := isUrl(paste.Text); ok {
			cm := model.ScrapingInfo{PasteId: paste.ID, Url: paste.Text}
			err := h.Bus.Publish(workerutils.PasteCrawlTopic(), cm)
			if err != nil {
				logger.Log.Error(err)
			}
		} else {
			//Create embedding
			h.Bus.Publish(workerutils.EmbeddingTopic(), paste.Text)
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
func (h *WorkerHandler) Scrapper(message interface{}) error {
	if cm, ok := message.(model.ScrapingInfo); ok {
		htmlContent, err := fetchContent(cm.Url)
		if err != nil {
			logger.Log.Error(err)
			return err
		}
		rootHtml, err := html.Parse(strings.NewReader(htmlContent))
		if err != nil {
			return err
		}
		plainText := extractText(rootHtml)
		plainText = cancelNoise(plainText)
		h.Bus.Publish(workerutils.EmbeddingTopic(), &model.EmbeddingPayload{PasteId: cm.PasteId, Text: plainText})
	}
	return nil
}

func (h *WorkerHandler) EmbeddingHandler(message interface{}) error {
	payload := message.(model.EmbeddingPayload)
	chunks := chunkTextByTokens(payload.Text, 512)
	for _, chunk := range chunks {
		reqBody := &model.EmbeddingRequestBody{Input: chunk}
		response, err := helpers.GetEmbedding(reqBody)
		fmt.Print(response.Model)
		if err != nil {
			return err
		}
	}
	return nil
}
func tokenize(text string) []string {
	return strings.Fields(text)
}
func chunkTextByTokens(text string, maxTokens int) []string {
	tokens := tokenize(text)
	var chunks []string

	for i := 0; i < len(tokens); i += maxTokens {
		end := i + maxTokens
		if end > len(tokens) {
			end = len(tokens)
		}
		chunk := strings.Join(tokens[i:end], " ")
		chunks = append(chunks, chunk)
	}
	return chunks
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
func extractText(n *html.Node) string {
	var text string

	if n.Type == html.TextNode {
		text = n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}

	return text
}
func cancelNoise(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")     // Replace newline characters with space
	text = strings.ReplaceAll(text, "\r", "")      // Remove carriage return characters
	text = strings.ReplaceAll(text, "\t", " ")     // Replace tabs with space
	text = strings.Join(strings.Fields(text), " ") // Remove extra whitespace
	text = strings.TrimSpace(text)                 // Trim leading and trailing whitespace
	return text
}
