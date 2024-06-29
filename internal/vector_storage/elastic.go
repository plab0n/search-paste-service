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
	"math"
	"strings"
)

type VectorStorage interface {
	CreateIndex(ctx context.Context, name string) error
	IndexDocument(ctx context.Context, index string, id string, vector []float64) error
	SearchDocument(ctx context.Context, index string, queryVector []float64, k int) error
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
					"dims": 1024,
					"index": true,
					"similarity": "dot_product",
					"index_options": {
    					"type": "hnsw",
    					"ef_construction": 128,
    					"m": 24    
  					}
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
func (e *ElasticSearch) IndexDocument(ctx context.Context, index string, id string, vector []float64) error {
	vector = normalizeVector(vector)
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
		Refresh:    "true",
	}

	res, err := req.Do(ctx, e.client)
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
func (e *ElasticSearch) SearchDocument(ctx context.Context, index string, queryVector []float64, k int) error {
	queryVector = normalizeVector(queryVector)
	query := map[string]interface{}{
		"knn": map[string]interface{}{
			"field":          "paste_vector",
			"query_vector":   queryVector,
			"k":              k,
			"num_candidates": 100,
		},
	}
	body, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("error marshaling query: %v", err)
	}

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(body),
	}
	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error performing search: %s", res.String())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	logger.Log.Log(logrus.InfoLevel, "Found %d documents\n", int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)))

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		logger.Log.Log(logrus.InfoLevel, "Document ID: %s, Score: %f\n", doc["_id"], doc["_score"].(float64))
	}
	return nil
}

func normalizeVector(vector []float64) []float64 {
	var magnitude float64
	for _, v := range vector {
		magnitude += v * v
	}
	magnitude = math.Sqrt(magnitude)

	normalized := make([]float64, len(vector))
	for i, v := range vector {
		normalized[i] = v / magnitude
	}
	return normalized
}
