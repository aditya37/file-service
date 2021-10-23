package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type (
	FirebaseClient struct {
		client *firebase.App
	}
	FirebaseClientConfig struct {
		BucketName     string
		ProjectID      string
		CredentialFile string
	}
)

var ctx context.Context

func NewFirebaseClient(config FirebaseClientConfig) (*FirebaseClient, error) {
	clients, err := firebase.NewApp(ctx, &firebase.Config{
		StorageBucket: config.BucketName,
		ProjectID:     config.ProjectID,
	}, option.WithCredentialsFile(config.CredentialFile), option.WithServiceAccountFile(config.CredentialFile))
	if err != nil {
		return nil, err
	}
	log.Printf("[Firebase] Storage Bucket: %s Project ID: %s", config.BucketName, config.ProjectID)
	return &FirebaseClient{
		client: clients,
	}, nil
}
