package main

import (
	"fmt"
	"log"

	"github.com/maantos/working-days-tracker/internal/storage"
	"github.com/maantos/working-days-tracker/internal/storage/mongo"
)

func main() {
	var db storage.Storage = &mongo.MongoStorage{}

	if err := db.Initialize(); err != nil {
		log.Fatalf("Error initializing storage: %v", err)
	}

	response, _ := db.GetWorkDays("2024", "January")

	for _, item := range response {
		fmt.Printf("Data: %s WorkingData %t \n", item["date"], item["workingDay"])
	}
}
