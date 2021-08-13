package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	fmt.Printf("Trigger Func 1 received S3 event: %+v", s3Event)

	return nil
}

func main() {
	lambda.Start(handler)
}
