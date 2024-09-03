package main

import (
	"log"
	elasticsearch "practice/elasitcsearch"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading .env file")
    }
  elasticsearch.Query("Jasmine","persons","name")
}