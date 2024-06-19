package vector_storage_test

import (
	"context"
	"github.com/plab0n/search-paste/internal/vector_storage"
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
}
