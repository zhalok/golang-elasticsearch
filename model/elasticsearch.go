package model

type Hit struct {
    ID     string                 `json:"_id"`
    Source map[string]interface{} `json:"_source"`
}

// SearchResults holds the structure of the search results
type SearchResults struct {
    TotalHits int
    Hits      []Hit
}