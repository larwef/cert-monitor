package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/larwef/cert-monitor/pkg/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
