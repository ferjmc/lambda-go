package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func Handler(ctx context.Context) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Println("AwsRequestID: ", lc.AwsRequestID)
	log.Println("IdentityPoolID: ", lc.Identity.CognitoIdentityPoolID)
	log.Println("AppTitle: ", lc.ClientContext.Client.AppTitle)
	log.Println("CognitoIdentityID: ", lc.Identity.CognitoIdentityID)
}

func main() {
	lambda.Start(Handler)
}
