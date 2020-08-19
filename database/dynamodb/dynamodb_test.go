package dynamodb

import (
	"fmt"
	"testing"

	awsAuth "github.com/safehousetech/gocloud/auth"
)

func init() {
	awsAuth.LoadConfig()
}

func TestDescribeTables(t *testing.T) {

	var dynamodb Dynamodb

	describetables := map[string]interface{}{
		"Region":    "us-east-1",
		"TableName": "Thread2",
	}

	resp, err := dynamodb.DescribeTables(describetables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}

func TestListTables(t *testing.T) {

	var dynamodb Dynamodb

	listtables := map[string]interface{}{
		"Region":                  "us-east-1",
		"ExclusiveStartTableName": "Thread",
		"Limit":                   1,
	}

	resp, err := dynamodb.ListTables(listtables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}

func TestDeletetables(t *testing.T) {

	var dynamodb Dynamodb

	deletetables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "Thread",
	}

	resp, err := dynamodb.DeleteTables(deletetables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}

/*
func TestCreatetables(t *testing.T) {

	var dynamodb Dynamodb

	keySchema := []map[string]interface{}{
		{
			"AttributeName": "ForumName",
			"KeyType":       "HASH",
		},
		{
			"AttributeName": "Subject",
			"KeyType":       "RANGE",
		},
	}

	attributeDefinitions := []map[string]interface{}{

		{
			"AttributeName": "ForumName",
			"AttributeType": "S",
		},
		{
			"AttributeName": "Subject",
			"AttributeType": "S",
		},
		{
			"AttributeName": "LastPostDateTime",
			"AttributeType": "S",
		},
	}

	projection := map[string]interface{}{
		"ProjectionType": "KEYS_ONLY",
	}

	provisionedThroughput := map[string]interface{}{
		"ReadCapacityUnits":  5,
		"WriteCapacityUnits": 5,
	}

	localSecondaryIndexes := []map[string]interface{}{
		{
			"IndexName":  "LastPostIndex",
			"KeySchema":  keySchema,
			"Projection": projection,
		},
	}

	createtables := map[string]interface{}{
		"Region":                "us-east-1",
		"TableName":             "Thread3",
		"KeySchema":             keySchema,
		"AttributeDefinitions":  attributeDefinitions,
		"LocalSecondaryIndexes": localSecondaryIndexes,
		"ProvisionedThroughput": provisionedThroughput,
	}

	resp, err := dynamodb.Createtables(createtables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
*/
