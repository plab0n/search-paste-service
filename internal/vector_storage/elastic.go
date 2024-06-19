package vector_storage

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
)

type VectorStorage interface {
	CreateIndex(ctx context.Context, name string) error
	IndexDocument(ctx context.Context, id string, vector []float32) error
	SearchDocument(ctx context.Context, queryVector []float32, k int) error
}

type ElasticSearch struct {
	client *elasticsearch.TypedClient
}

func NewElasticDb() (*ElasticSearch, error) {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		APIKey:    "TFpDX0w1QUJGNm04S2RXWGVkaFA6TUx0RG9YY0pRUzZZY2tMa0I1VUc5Zw==",
	})
	return &ElasticSearch{client: client}, err
}

func (e *ElasticSearch) CreateIndex(ctx context.Context, name string) error {

	return nil
}
func (e *ElasticSearch) IndexDocument(ctx context.Context, id string, vector []float32) error {
	return nil
}
func (e *ElasticSearch) SearchDocument(ctx context.Context, queryVector []float32, k int) error {
	return nil
}
