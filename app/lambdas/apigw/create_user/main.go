package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fenderdigital/bv-medium-users-service/app/conf"
	"log"
)

func main() {
	f, err := conf.NewCreateUserFeature()
	if err != nil {
		log.Fatal(err, "conf.NewCreateUserFeature failed")
	}

	h := handler{
		feature: f,
	}

	lambda.Start(h.Handle)
}

type Request struct {
	Name  string `json:"name"`
	Email string `json:"emaill"`
}

type UserCreator interface {
	Create(ctx context.Context, name, email string) error
}

type handler struct {
	feature UserCreator
}

func (h handler) Handle(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	if err := json.Unmarshal([]byte(e.Body), &req); err != nil {
		log.Println("json.Unmarshal( failed", err)
		return events.APIGatewayProxyResponse{Body: "bad request", StatusCode: 400}, nil
	}

	if err := h.feature.Create(ctx, req.Name, req.Email); err != nil {
		log.Println("feature.Create failed", err)
		return events.APIGatewayProxyResponse{Body: "server error", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}
