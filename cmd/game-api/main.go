package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// App holds connections, etc.
type App struct {
}

// Handle API requestss
func (a *App) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{Body: "no", StatusCode: 401}, errors.New("not yet implemented")
}

func main() {
	// TODO: stuff to connect to whatever
	app := &App{}
	lambda.Start(app.Handle)
}
