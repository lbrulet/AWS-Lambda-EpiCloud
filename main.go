package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lbrulet/AWS-Lambda-EpiCloud/controllers"
	"github.com/lbrulet/AWS-Lambda-EpiCloud/database"
)

type DeleteTripRequest struct {
	ID int `json:"id"`
}

type ReponsePayload struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func handleRequest(ctx context.Context, request DeleteTripRequest) (ReponsePayload, error) {
	db, err := database.InitDB()
	if err != nil {
		return ReponsePayload{Message: "Database initialisation error", Success: false}, err
	}
	defer db.Close()
	err = controllers.DeleteTrip(db, request.ID)
	if err != nil {
		return ReponsePayload{Message: "Database error", Success: false}, err
	}
	return ReponsePayload{Message: "Trip deleted", Success: true}, nil
}
func main() {
	lambda.Start(handleRequest)
}
