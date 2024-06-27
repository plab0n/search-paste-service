package vector_storage_test

import (
	"context"
	"github.com/plab0n/search-paste/internal/model"
	"github.com/plab0n/search-paste/internal/vector_storage"
	"github.com/plab0n/search-paste/pkg/helpers"
	"testing"
)

func Test_CreateConnection(t *testing.T) {
	_, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
}
func Test_CreateIndex(t *testing.T) {
	es, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	err = es.CreateIndex(ctx, "test_vector_index")
	if err != nil {
		t.Errorf(err.Error())
	}
}
func Test_IndexDocument(t *testing.T) {
	es, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	var vector []float64
	for i := 0; i < 1024; i++ {
		vector = append(vector, 0.3424532434243424)
	}
	err = es.IndexDocument(ctx, "test_vector_index", "n1", vector)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func Test_SearchDocument(t *testing.T) {
	es, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	var vector []float64
	for i := 0; i < 1024; i++ {
		vector = append(vector, 1.0)
	}
	err = es.SearchDocument(ctx, "test_vector_index", vector, 2)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func Test_IndexText(t *testing.T) {
	ctx := context.Background()
	req := &model.EmbeddingRequestBody{Input: "Golang is well known for concurrency."}
	res, err := helpers.GetEmbedding(req)
	if err != nil {
		t.Error(err)
	}
	var vec []float64
	for _, v := range res.Data[0].Embedding {
		vec = append(vec, v)
	}
	es, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
	err = es.IndexDocument(ctx, "test_vector_index", "n1", vec)
	if err != nil {
		t.Errorf(err.Error())
	}
}
func Test_SearchText(t *testing.T) {
	ctx := context.Background()
	req := &model.EmbeddingRequestBody{Input: "java"}
	res, err := helpers.GetEmbedding(req)
	if err != nil {
		t.Error(err)
	}
	var vec []float64
	for _, v := range res.Data[0].Embedding {
		vec = append(vec, v)
	}
	es, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
	err = es.SearchDocument(ctx, "test_vector_index", vec, 2)
	if err != nil {
		t.Errorf(err.Error())
	}
}
