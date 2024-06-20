package vector_storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/plab0n/search-paste/pkg/logger"
	"github.com/sirupsen/logrus"
	"strings"
)

type VectorStorage interface {
	CreateIndex(ctx context.Context, name string) error
	IndexDocument(ctx context.Context, index string, id string, vector []float32) error
	SearchDocument(ctx context.Context, queryVector []float32, k int) error
}

type ElasticSearch struct {
	client *elasticsearch.TypedClient
}

func NewElasticDb() (*ElasticSearch, error) {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses:            []string{"http://localhost:9200"},
		APIKey:               "TFpDX0w1QUJGNm04S2RXWGVkaFA6TUx0RG9YY0pRUzZZY2tMa0I1VUc5Zw==",
		DiscoverNodesOnStart: true,
	})
	return &ElasticSearch{client: client}, err
}

func (e *ElasticSearch) CreateIndex(ctx context.Context, name string) error {
	mapping := `{
		"mappings": {
			"properties": {
				"paste_vector": {
					"type": "dense_vector",
					"dims": 1024
				}
			}
		}
	}`
	req := esapi.IndicesCreateRequest{
		Index: name,
		Body:  strings.NewReader(mapping),
	}
	res, err := req.Do(ctx, e.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}
	logger.Log.Log(logrus.InfoLevel, "Index Created")
	return nil
}
func (e *ElasticSearch) IndexDocument(ctx context.Context, index string, id string, vector []float32) error {
	doc := map[string]interface{}{
		"paste_vector": vector,
	}
	body, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("error marshaling document: %v", err)
	}
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(body),
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.String())
	}
	logger.Log.Log(logrus.InfoLevel, "Document %s indexed", id)
	return nil
}
func (e *ElasticSearch) SearchDocument(ctx context.Context, queryVector []float32, k int) error {
	return nil
}
