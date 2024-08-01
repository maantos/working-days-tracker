package mongo

import (
	"fmt"
)

type MongoStorage struct {
	// client     *mongo.Client
	// collection *mongo.Collection
	collection []map[string]interface{}
}

func (m *MongoStorage) Initialize() error {
	m.collection = []map[string]interface{}{
		{
			"date":       "01-01-2024",
			"workingDay": true,
			"notes":      "New Year Day",
		},
		{
			"date":       "15-01-2024",
			"workingDay": false,
			"notes":      "Mid-January break",
		},
		{
			"date":       "25-01-2024",
			"workingDay": true,
			"notes":      "Regular working day",
		},
	}

	fmt.Println("DataBase conn initialized...")
	return nil
}

func (m *MongoStorage) GetWorkDays(year string, month string) ([]map[string]interface{}, error) {
	return m.collection, nil
}

func (m *MongoStorage) MarkWorking(year string, month string) error {
	return nil
}
