package vector_storage_test

import (
	"github.com/plab0n/search-paste/internal/vector_storage"
	"testing"
)

func Test_ElasticConnection(t *testing.T) {
	_, err := vector_storage.NewElasticDb()
	if err != nil {
		t.Error(err)
	}
}
