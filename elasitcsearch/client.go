package elasticsearch

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func GetClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{os.Getenv("ELASTICSEARCH_NODE")},
			Username: os.Getenv("ELASTICSEARCH_USERNAME"),
			Password: os.Getenv("ELASTICSEARCH_PASSWORD"),
		},
	)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return es
}