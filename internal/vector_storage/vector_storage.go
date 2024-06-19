package vector_storage

import "github.com/elastic/go-elasticsearch/v8"

func NewElasticDb() (*elasticsearch.TypedClient, error) {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		APIKey:    "TFpDX0w1QUJGNm04S2RXWGVkaFA6TUx0RG9YY0pRUzZZY2tMa0I1VUc5Zw==",
	})
	return client, err
}
