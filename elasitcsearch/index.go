package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

func doesIndexExists(indexName string,es *elasticsearch.Client)bool{
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
        log.Fatalf("Error checking if the index exists: %s", err)
    }
	defer res.Body.Close()
	return res.StatusCode != 404
	
}

func createIndex(indexName string,es *elasticsearch.Client) {
 // Index does not exist, so create it
 createIndexRes, err := es.Indices.Create(indexName)
 if err != nil {
	 log.Fatalf("Error creating the index: %s", err)
 }
 defer createIndexRes.Body.Close()

 if createIndexRes.IsError() {
	 log.Fatalf("Error response from Elasticsearch: %s", createIndexRes.String())
 }
 fmt.Printf("Index %s created.\n", indexName)

}

func Index(indexName string, document map[string]interface{}){
	es:= GetClient()
	
	if !doesIndexExists(indexName,es) {
       createIndex(indexName,es)
    }  else {
        fmt.Printf("Index %s already exists.\n", indexName)
    }
	// replace with the data read with file reader from the json file

    docJSON, err := json.Marshal(document)
    if err != nil {
        log.Fatalf("Error marshaling document: %s", err)
    }

    req := esapi.IndexRequest{
        Index:      indexName,
        DocumentID: document["id"].(string),
        Body:       strings.NewReader(string(docJSON)),
        Refresh:    "true",
    }

	insertRes, err := req.Do(context.Background(), es)
    if err != nil {
        log.Fatalf("Error indexing document: %s", err)
    }
    defer insertRes.Body.Close()

	if insertRes.IsError() {
        log.Fatalf("Error response from Elasticsearch: %s", insertRes.String())
    } else {
        fmt.Printf("Document ID %s inserted into index %s.\n", document["id"], indexName)
    }

}