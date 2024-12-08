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
