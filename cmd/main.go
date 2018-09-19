package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kierendavies/alexa-ov/pkg/response"
)

func HandleRequest(ctx context.Context) (response.Body, error) {
	return response.NewText("this is o.v."), nil
}

func main() {
	lambda.Start(HandleRequest)
}
