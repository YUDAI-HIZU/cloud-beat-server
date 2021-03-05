package config

import (
	"fmt"
	"os"
)

var (
	ENV             string
	DatabaseURL     string
	JwtSecret       string
	GCSAccount      string
	FirebaseAccount string
	BucketName      string
	StorageURL      string
	AssetEndpoint   string
)

func init() {
	ENV = os.Getenv("ENV")
	DatabaseURL = os.Getenv("DATABASE_URL")
	JwtSecret = os.Getenv("JWT_SECRET")
	GCSAccount = os.Getenv("GCS_ACCOUNT")
	FirebaseAccount = os.Getenv("FIREBASE_ACCOUNT")
	BucketName = os.Getenv("BUCKET_NAME")
	StorageURL = os.Getenv("STORAGE_URL")
	AssetEndpoint = fmt.Sprintf("%s/%s", StorageURL, BucketName)
}
