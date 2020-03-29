package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

var a *App

func TestHandle(t *testing.T) {
	a.Handle(context.Background(), events.APIGatewayProxyRequest{})
}

func setup() {
	a := &App{}
}