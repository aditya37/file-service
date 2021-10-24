package repository

import (
	"context"
	"io"
	"os"

	gcp "cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type (
	firebaseStorage struct {
		bucket *gcp.BucketHandle
	}
	FirebaseStorage interface {
		Upload(ctx context.Context, filetPath, objectName string) (string, error)
		GetObjectAttribute(ctx context.Context, objectName string) (*gcp.ObjectAttrs, error)
		GetObject(ctx context.Context, objectName string) (*gcp.ObjectAttrs, error)
	}
)

func NewFirebaseStorage(app FirebaseClient) (FirebaseStorage, error) {
	// init instance for firebase storage
	storages, err := app.client.Storage(ctx)
	if err != nil {
		return nil, err
	}

	bucket, err := storages.DefaultBucket()
	if err != nil {
		return nil, err
	}

	return &firebaseStorage{
		bucket: bucket,
	}, nil
}

func (g *firebaseStorage) Upload(ctx context.Context, filetPath, objectName string) (string, error) {

	// read file from filepath
	file, err := os.Open(filetPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	id := uuid.New()
	writeBucket := g.bucket.Object(objectName).NewWriter(ctx)
	// set or generate download token
	writeBucket.ObjectAttrs.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": id.String(),
	}
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(writeBucket, file); err != nil {
		return "", err
	}
	if err := writeBucket.Close(); err != nil {
		return "", err
	}

	// set acl
	obj := g.bucket.Object(writeBucket.Name)
	if err := obj.ACL().Set(ctx, gcp.AllUsers, gcp.RoleReader); err != nil {
		return "", err
	}

	return writeBucket.Name, nil
}

func (g *firebaseStorage) GetObjectAttribute(ctx context.Context, objectName string) (*gcp.ObjectAttrs, error) {
	obj, err := g.bucket.Object(objectName).Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (g *firebaseStorage) GetObject(ctx context.Context, objectName string) (*gcp.ObjectAttrs, error) {
	obj, err := g.bucket.Object(objectName).Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
