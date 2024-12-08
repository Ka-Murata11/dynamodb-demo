package test_test

import (
	"app/db"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gocarina/gocsv"
)

type ValueForCSV struct {
	UserID    string `csv:"ユーザID" dynamodbav:"user_id"`
	TableName string `csv:"テーブル名"`
	Timestamp string `csv:"作成日時" dynamodbav:"timestamp"`
	Value     string `csv:"値" dynamodbav:"values"`
}

// TestQueryDatas はデータを取得するテストです。
func TestQueryDatas(t *testing.T) {
	svc := db.Init()

	userIDs := []string{"1", "2"}

	tables := []string{"TypeA", "TypeB", "TypeC"}

	var results []ValueForCSV

	for _, userID := range userIDs {
		for _, table := range tables {
			// クエリの入力パラメータ
			input := &dynamodb.QueryInput{
				TableName: aws.String(table),
				KeyConditions: map[string]types.Condition{
					"user_id": {
						ComparisonOperator: types.ComparisonOperatorEq,
						AttributeValueList: []types.AttributeValue{
							&types.AttributeValueMemberS{Value: userID},
						},
					},
					"timestamp": {
						ComparisonOperator: types.ComparisonOperatorBetween,
						AttributeValueList: []types.AttributeValue{
							&types.AttributeValueMemberS{Value: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).String()},
							&types.AttributeValueMemberS{Value: time.Date(2024, 1, 1, 1, 30, 0, 0, time.UTC).String()},
						},
					},
				},
			}

			// クエリを実行
			result, err := svc.Query(context.TODO(), input)
			if err != nil {
				log.Fatalf("Got error calling Query: %v", err)
			}

			// 結果をアンマーシャル
			for _, item := range result.Items {
				var data ValueForCSV
				err = attributevalue.UnmarshalMap(item, &data)
				if err != nil {
					log.Fatalf("Failed to unmarshal data: %v", err)
				}
				data.TableName = table
				results = append(results, data)
			}

		}
	}

	// tmpフォルダにCSVファイルを作成
	file, err := os.Create("/tmp/output.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// CSVファイルに書き込む
	err = gocsv.MarshalFile(&results, file)
	if err != nil {
		log.Fatalf("Failed to write data to CSV: %v", err)
	}

	// 結果を表示
	fmt.Println("Successfully wrote data to CSV")
}
