package main

import (
	"context"
	"log/slog"
	"os"

	"playground/api-auth-verifier/internal/auth"
	"playground/api-auth-verifier/internal/config"
	"playground/api-auth-verifier/internal/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("fatal error loading config: ", "error", err)
		os.Exit(1)
	}

	validator, err := auth.NewTokenValidator(cfg)
	if err != nil {
		slog.Error("fatal error creating validator: ", "error", err)
		os.Exit(1)
	}

	log := slog.New(&slog.JSONHandler{})

	handlerWrapper := func(
		ctx context.Context,
		request events.APIGatewayProxyRequest,
	) (events.APIGatewayProxyResponse, error) {
		return handler.HandlerRequest(ctx, log, cfg, request, validator)
	}

	lambda.Start(handlerWrapper)
}
