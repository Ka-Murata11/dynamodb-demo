package test_test

import (
	"app/db"
	"app/model"
	"context"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// TestInsertData はデータを挿入するテストです。
// func TestInsertData(t *testing.T) {
// 	svc := db.Init()

// 	// insert data
// 	insertData := []model.TimeSeriseValue{
// 		{
// 			UserID:    "1",
// 			Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String(),
// 			Values: []model.TimeSeriseValueList{
// 				{
// 					ValueType: "1",
// 					ValueID:   "1",
// 					Value:     "100",
// 				},
// 				{
// 					ValueType: "2",
// 					ValueID:   "2",
// 					Value:     "200",
// 				},
// 				{
// 					ValueType: "3",
// 					ValueID:   "3",
// 					Value:     "300",
// 				},
// 				{
// 					ValueType: "4",
// 					ValueID:   "4",
// 					Value:     "400",
// 				},
// 				{
// 					ValueType: "5",
// 					ValueID:   "5",
// 					Value:     "500",
// 				},
// 			},
// 		},
// 		{
// 			UserID:    "1",
// 			Timestamp: time.Date(2024, 1, 1, 0, 30, 0, 0, time.UTC).String(),
// 			Values: []model.TimeSeriseValueList{
// 				{
// 					ValueType: "1",
// 					ValueID:   "1",
// 					Value:     "300",
// 				},
// 				{
// 					ValueType: "2",
// 					ValueID:   "2",
// 					Value:     "100",
// 				},
// 				{
// 					ValueType: "3",
// 					ValueID:   "3",
// 					Value:     "600",
// 				},
// 			},
// 		},
// 		{
// 			UserID:    "1",
// 			Timestamp: time.Date(2024, 1, 10, 0, 30, 0, 0, time.UTC).String(),
// 			Values: []model.TimeSeriseValueList{
// 				{
// 					ValueType: "1",
// 					ValueID:   "1",
// 					Value:     "30000000000000",
// 				},
// 				{
// 					ValueType: "2",
// 					ValueID:   "2",
// 					Value:     "100000000000000",
// 				},
// 				{
// 					ValueType: "3",
// 					ValueID:   "3",
// 					Value:     "6000000000000000",
// 				},
// 			},
// 		},
// 		{
// 			UserID:    "2",
// 			Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String(),
// 			Values: []model.TimeSeriseValueList{
// 				{
// 					ValueType: "1",
// 					ValueID:   "1",
// 					Value:     "500",
// 				},
// 				{
// 					ValueType: "2",
// 					ValueID:   "2",
// 					Value:     "400",
// 				},
// 				{
// 					ValueType: "3",
// 					ValueID:   "3",
// 					Value:     "300",
// 				},
// 				{
// 					ValueType: "4",
// 					ValueID:   "4",
// 					Value:     "200",
// 				},
// 				{
// 					ValueType: "5",
// 					ValueID:   "5",
// 					Value:     "100",
// 				},
// 			},
// 		},
// 	}

// 	writeRequests := make([]types.WriteRequest, len(insertData))

// 	for i, data := range insertData {
// 		valuesList := make([]types.AttributeValue, len(data.Values))
// 		for j, v := range data.Values {
// 			valueJSON, err := json.Marshal(v)
// 			if err != nil {
// 				log.Fatalf("Failed to marshal value to JSON: %v", err)
// 			}
// 			valuesList[j] = &types.AttributeValueMemberS{Value: string(valueJSON)}
// 		}

// 		item := map[string]types.AttributeValue{
// 			"user_id":   &types.AttributeValueMemberS{Value: data.UserID},
// 			"timestamp": &types.AttributeValueMemberS{Value: data.Timestamp},
// 			"values":    &types.AttributeValueMemberL{Value: valuesList},
// 		}

// 		writeRequests[i] = types.WriteRequest{
// 			PutRequest: &types.PutRequest{
// 				Item: item,
// 			},
// 		}
// 	}

// 	_, err := svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
// 		RequestItems: map[string][]types.WriteRequest{
// 			"TimeSeriseTable": writeRequests,
// 		},
// 	})
// 	if err != nil {
// 		log.Fatalf("Got error calling BatchWriteItem: %v", err)
// 	}

// 	log.Println("Successfully inserted data into TimeSeriseTable")

// }

// TestInsertDatas はデータを挿入するテストです。
func TestInsertDatas(t *testing.T) {
	svc := db.Init()

	// insert data
	insertData := []model.TypeData{

		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String(),
			Value:     "100",
		},
		{
			UserID:    "2",
			Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String(),
			Value:     "100000",
		},
		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 1, 0, 30, 0, 0, time.UTC).String(),
			Value:     "200",
		},
		{
			UserID:    "2",
			Timestamp: time.Date(2024, 1, 1, 0, 30, 0, 0, time.UTC).String(),
			Value:     "300",
		},
		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC).String(),
			Value:     "500",
		},
		{
			UserID:    "2",
			Timestamp: time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC).String(),
			Value:     "100000",
		},
		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 1, 1, 30, 0, 0, time.UTC).String(),
			Value:     "11200",
		},
		{
			UserID:    "2",
			Timestamp: time.Date(2024, 1, 1, 1, 30, 0, 0, time.UTC).String(),
			Value:     "300",
		},
	}

	writeRequests := make([]types.WriteRequest, len(insertData))

	for i, data := range insertData {

		item := map[string]types.AttributeValue{
			"user_id":   &types.AttributeValueMemberS{Value: data.UserID},
			"timestamp": &types.AttributeValueMemberS{Value: data.Timestamp},
			"values":    &types.AttributeValueMemberS{Value: data.Value},
		}

		writeRequests[i] = types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		}
	}

	tables := []string{"TypeA", "TypeB", "TypeC"}
	for _, table := range tables {
		_, err := svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				table: writeRequests,
			},
		})
		if err != nil {
			log.Fatalf("Got error calling BatchWriteItem: %v", err)
		}
	}

	log.Println("Successfully inserted data")

}
