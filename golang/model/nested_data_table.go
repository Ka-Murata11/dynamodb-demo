package model

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type NestedDataTable struct {
	ID        string          `dynamodbav:"id"`
	Data      map[string]Data `dynamodbav:"data"`
	CreatedAt string          `dynamodbav:"createdAt"`
}

type NestedDataTablePrepare struct {
	ID        string                 `dynamodbav:"id"`
	Data      map[string]interface{} `dynamodbav:"data"`
	CreatedAt string                 `dynamodbav:"createdAt"`
}

type Data struct {
	Value1 float64 `dynamodbav:"value1"`
	Value2 float64 `dynamodbav:"value2"`
}

func (nestedDataTable *NestedDataTable) AddNestedData(ctx context.Context, client *dynamodb.Client) error {
	item, err := attributevalue.MarshalMap(nestedDataTable)
	if err != nil {
		return err
	}
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("NestedDataTable"),
		Item:      item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

func (nestedDataTable *NestedDataTable) GetNestedDataTable(ctx context.Context, client *dynamodb.Client) error {
	result, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("NestedDataTable"),
		Key:       nestedDataTable.GetKey(),
	})
	if err != nil {
		log.Printf("Got error calling GetItem: %v", err)
	}
	var prepare NestedDataTablePrepare
	if err := attributevalue.UnmarshalMap(result.Item, &prepare); err != nil {
		log.Printf("Got error unmarshalling: %v", err)
	}

	nestedDataTable.Data, err = UnmarshalNestedData[Data](prepare.Data)

	return err
}

func UnmarshalNestedData[T any](data map[string]interface{}) (map[string]T, error) {
	result := make(map[string]T, len(data))
	for key, v := range data {
		var item T
		av, err := attributevalue.MarshalMap(v)
		if err != nil {
			return nil, err
		}
		err = attributevalue.UnmarshalMap(av, &item)
		if err != nil {
			return nil, err
		}
		result[key] = item
	}
	return result, nil
}

func (nestedDataTable *NestedDataTable) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(nestedDataTable.ID)
	if err != nil {
		panic(err)
	}
	createdAt, err := attributevalue.Marshal(nestedDataTable.CreatedAt)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id, "createdAt": createdAt}
}
