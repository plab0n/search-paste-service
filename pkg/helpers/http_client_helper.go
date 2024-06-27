package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/plab0n/search-paste/internal/model"
	"io"
	"net/http"
	"os"
)

func GetEmbedding(reqBody *model.EmbeddingRequestBody) (*model.EmbeddingResponse, error) {
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	embeddingApi := os.Getenv("EMBEDDING_API")
	if len(embeddingApi) == 0 {
		embeddingApi = "http://localhost:8000/v1/embeddings"
	}
	embeddingReq, err := http.NewRequest("POST", embeddingApi, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	res, err := httpClient.Do(embeddingReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Http request failed. StatusCode: %d", res.StatusCode))
	}
	embeddingResponse := &model.EmbeddingResponse{}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resBody, embeddingResponse)
	if err != nil {
		return nil, err
	}
	return embeddingResponse, nil
}
