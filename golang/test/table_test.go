package test_test

import (
	"app/db"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
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

// TestCreateTables はテーブルを作成するテストです。
func TestCreateTables(t *testing.T) {
	svc := db.Init()

	tables := []string{"TypeA", "TypeB", "TypeC"}

	for _, table := range tables {
		// テーブル作成の入力パラメータ
		input := &dynamodb.CreateTableInput{
			TableName: aws.String(table),
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

}
