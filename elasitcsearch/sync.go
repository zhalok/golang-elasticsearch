package elasticsearch

import (
	"fmt"
	"practice/helper"
)

func Sync(indexName string){

	var myMap []map[string]interface{}

	helper.ReadData("persons", &myMap)

	for _,item := range myMap{
		searchResults,err:= Query(item["name"].(string),indexName,"name")
		if err != nil{
			fmt.Println(err)
		}
		if searchResults.TotalHits == 0{
			Index(indexName,item)
		}
	}

	
	


	
}