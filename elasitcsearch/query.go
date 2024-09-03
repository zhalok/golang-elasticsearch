package elasticsearch

import (
	"context"
	"encoding/json"
	"log"
	"practice/model"
	"strings"
)


func Query(searchTerm string, indexName string,fieldName string)(model.SearchResults, error){
    es:=GetClient()
   
	query := map[string]interface{}{
        "query": map[string]interface{}{
			"match": map[string]interface{}{
                fieldName: searchTerm,
            },
        },
    }
	var buf strings.Builder
    if err := json.NewEncoder(&buf).Encode(query); err != nil {
        log.Fatalf("Error encoding query: %s", err)
    }
	reader := strings.NewReader(buf.String())

	res, err := es.Search(
        es.Search.WithContext(context.Background()),
        es.Search.WithIndex(indexName), // specify the index name
        es.Search.WithBody(reader),
        es.Search.WithTrackTotalHits(true),
        es.Search.WithPretty(),
    )
    if err != nil {
        log.Fatalf("Error getting response: %s", err)
    }
    defer res.Body.Close()
	if res.IsError() {
        var e map[string]interface{}
        if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
            log.Fatalf("Error parsing the response body: %s", err)
        } else {
            log.Fatalf("Error: [%s] %s: %s",
                res.Status(),
                e["error"].(map[string]interface{})["type"],
                e["error"].(map[string]interface{})["reason"],
            )
        }
    }
	var result map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        log.Fatalf("Error parsing the response body: %s", err)
    }
   

	totalHits := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

    // Extract raw hits
    var hits []model.Hit
    for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
        h := hit.(map[string]interface{})
        hits = append(hits, model.Hit{
            ID:     h["_id"].(string),
            Source: h["_source"].(map[string]interface{}),
        })
    }

    return model.SearchResults{
        TotalHits: totalHits,
        Hits:      hits,
    }, nil

}