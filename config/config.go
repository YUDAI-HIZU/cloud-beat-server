package config

import (
	"os"
)

var (
	ENV             string
	DatabaseURL     string
	JwtSecret       string
	GCSAccount      string
	FirebaseAccount string
	BucketName      string
)

func init() {
	ENV = os.Getenv("ENV")
	DatabaseURL = os.Getenv("DATABASE_URL")
	JwtSecret = os.Getenv("JWT_SECRET")
	GCSAccount = os.Getenv("GCS_ACCOUNT")
	FirebaseAccount = os.Getenv("FIREBASE_ACCOUNT")
	BucketName = os.Getenv("BUCKET_NAME")
}
