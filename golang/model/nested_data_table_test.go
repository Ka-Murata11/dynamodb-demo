package model

import (
	"app/db"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNestedDataTable_AddNestedData(t *testing.T) {
	ddb := db.NewDynamoDBClient()
	item := []NestedDataTable{
		{
			ID: "1",
			Data: map[string]Data{
				"key1": {Value1: 10, Value2: 20},
				"key2": {Value1: 30, Value2: 40},
			},
			CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
		{
			ID: "2",
			Data: map[string]Data{
				"key1": {Value1: 50, Value2: 60},
				"key2": {Value1: 70, Value2: 80},
			},
			CreatedAt: time.Date(2025, 1, 1, 0, 30, 0, 0, time.UTC).Format(time.RFC3339),
		},
		{
			ID: "3",
			Data: map[string]Data{
				"key1": {Value1: 90, Value2: 100},
				"key2": {Value1: 110, Value2: 120},
			},
			CreatedAt: time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
	}
	for _, v := range item {
		if err := v.AddNestedData(context.TODO(), ddb.Client); err != nil {
			t.Errorf("failed to add item: %v", err)
		}
	}
}

func TestGetNestedDataTable(t *testing.T) {
	ddb := db.NewDynamoDBClient()
	item := NestedDataTable{
		ID:        "1",
		CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
	}
	if err := item.GetNestedDataTable(context.TODO(), ddb.Client); err != nil {
		t.Errorf("failed to get item: %v", err)
	}
	fmt.Println(item)
}
