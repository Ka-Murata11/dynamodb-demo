package model

type TimeSeriseValue struct {
	UserID    string                `dynamodbav:"user_id"`
	Timestamp string                `dynamodbav:"timestamp"`
	Values    []TimeSeriseValueList `dynamodbav:"values"`
}

type TimeSeriseValueList struct {
	ValueType string `dynamodbav:"value_type"`
	ValueID   string `dynamodbav:"value_id"`
	Value     string `dynamodbav:"value"`
}

type TimeSeriseValuePrepare struct {
	UserID    string   `dynamodbav:"user_id"`
	Timestamp string   `dynamodbav:"timestamp"`
	Values    []string `dynamodbav:"values"`
}
