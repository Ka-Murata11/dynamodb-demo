package test_test

import (
	"app/db"
	"app/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// TestCreateTable はテーブルを作成するテストです。
func TestCreateTable(t *testing.T) {
	svc := db.Init()

	// テーブル作成の入力パラメータ
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("TimeSeriseTable"),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("user_id"),
				KeyType:       types.KeyTypeHash, // パーティションキー
			},
			{
				AttributeName: aws.String("timestamp"),
				KeyType:       types.KeyTypeRange, // ソートキー
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("user_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("timestamp"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	// テーブルを作成
	_, err := svc.CreateTable(context.TODO(), input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %v", err)
	}

	fmt.Println("Created the table", *input.TableName)
}

// TestInsertData はデータを挿入するテストです。
func TestInsertData(t *testing.T) {
	svc := db.Init()

	// insert data
	insertData := []model.TimeSeriseValue{
		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String(),
			Values: []model.TimeSeriseValueList{
				{
					ValueType: "1",
					ValueID:   "1",
					Value:     "100",
				},
				{
					ValueType: "2",
					ValueID:   "2",
					Value:     "200",
				},
				{
					ValueType: "3",
					ValueID:   "3",
					Value:     "300",
				},
				{
					ValueType: "4",
					ValueID:   "4",
					Value:     "400",
				},
				{
					ValueType: "5",
					ValueID:   "5",
					Value:     "500",
				},
			},
		},
		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 1, 0, 30, 0, 0, time.UTC).String(),
			Values: []model.TimeSeriseValueList{
				{
					ValueType: "1",
					ValueID:   "1",
					Value:     "300",
				},
				{
					ValueType: "2",
					ValueID:   "2",
					Value:     "100",
				},
				{
					ValueType: "3",
					ValueID:   "3",
					Value:     "600",
				},
			},
		},
		{
			UserID:    "1",
			Timestamp: time.Date(2024, 1, 10, 0, 30, 0, 0, time.UTC).String(),
			Values: []model.TimeSeriseValueList{
				{
					ValueType: "1",
					ValueID:   "1",
					Value:     "30000000000000",
				},
				{
					ValueType: "2",
					ValueID:   "2",
					Value:     "100000000000000",
				},
				{
					ValueType: "3",
					ValueID:   "3",
					Value:     "6000000000000000",
				},
			},
		},
		{
			UserID:    "2",
			Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String(),
			Values: []model.TimeSeriseValueList{
				{
					ValueType: "1",
					ValueID:   "1",
					Value:     "500",
				},
				{
					ValueType: "2",
					ValueID:   "2",
					Value:     "400",
				},
				{
					ValueType: "3",
					ValueID:   "3",
					Value:     "300",
				},
				{
					ValueType: "4",
					ValueID:   "4",
					Value:     "200",
				},
				{
					ValueType: "5",
					ValueID:   "5",
					Value:     "100",
				},
			},
		},
	}

	writeRequests := make([]types.WriteRequest, len(insertData))

	for i, data := range insertData {
		valuesList := make([]types.AttributeValue, len(data.Values))
		for j, v := range data.Values {
			valueJSON, err := json.Marshal(v)
			if err != nil {
				log.Fatalf("Failed to marshal value to JSON: %v", err)
			}
			valuesList[j] = &types.AttributeValueMemberS{Value: string(valueJSON)}
		}

		item := map[string]types.AttributeValue{
			"user_id":   &types.AttributeValueMemberS{Value: data.UserID},
			"timestamp": &types.AttributeValueMemberS{Value: data.Timestamp},
			"values":    &types.AttributeValueMemberL{Value: valuesList},
		}

		writeRequests[i] = types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		}
	}

	_, err := svc.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			"TimeSeriseTable": writeRequests,
		},
	})
	if err != nil {
		log.Fatalf("Got error calling BatchWriteItem: %v", err)
	}

	log.Println("Successfully inserted data into TimeSeriseTable")

}

// TestQueryData はデータを取得するテストです。
func TestQueryData(t *testing.T) {
	svc := db.Init()

	// クエリの入力パラメータ
	input := &dynamodb.QueryInput{
		TableName: aws.String("TimeSeriseTable"),
		KeyConditions: map[string]types.Condition{
			"user_id": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: "1"},
				},
			},
			"timestamp": {
				ComparisonOperator: types.ComparisonOperatorBetween,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String()},
					&types.AttributeValueMemberS{Value: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC).String()},
				},
			},
		},
	}

	// クエリを実行
	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		log.Fatalf("Got error calling Query: %v", err)
	}

	// 結果を表示
	var timeSeriseValue []model.TimeSeriseValue
	for _, item := range result.Items {
		data, err := TimeSeriseValuePrepare(item)
		if err != nil {
			log.Fatalf("Failed to prepare TimeSeriseValue: %v", err)
		}

		timeSeriseValue = append(timeSeriseValue, data)
	}

	fmt.Println(timeSeriseValue)
}

// TimeSeriseValuePrepare はDynamoDBのアイテムをTimeSeriseValueに変換します。
func TimeSeriseValuePrepare(item map[string]types.AttributeValue) (model.TimeSeriseValue, error) {
	var data model.TimeSeriseValuePrepare
	err := attributevalue.UnmarshalMap(item, &data)
	if err != nil {
		return model.TimeSeriseValue{}, fmt.Errorf("failed to unmarshal item: %w", err)
	}

	var values []model.TimeSeriseValueList
	for _, v := range data.Values {
		var value model.TimeSeriseValueList
		err := json.Unmarshal([]byte(v), &value)
		if err != nil {
			return model.TimeSeriseValue{}, fmt.Errorf("failed to unmarshal Values attribute: %w", err)
		}
		values = append(values, value)
	}
	timeSeriesValue := model.TimeSeriseValue{
		UserID:    data.UserID,
		Timestamp: data.Timestamp,
		Values:    values,
	}

	return timeSeriesValue, nil
}
