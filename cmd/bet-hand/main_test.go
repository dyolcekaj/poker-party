package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

var a *App

func TestHandle(t *testing) {
	a.Handle(context.Background(), events.APIGatewayProxyRequest{})
}

func setup() {
	a := &App{}
}
