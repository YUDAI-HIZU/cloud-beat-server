package config

import (
	"os"
)

var (
	DatabaseURL     string
	JwtSecret       string
	GCSAccount      string
	FirebaseAccount string
)

func init() {
	DatabaseURL = os.Getenv("DATABASE_URL")
	JwtSecret = os.Getenv("JWT_SECRET")
	GCSAccount = os.Getenv("GCS_ACCOUNT")
	FirebaseAccount = os.Getenv("FIREBASE_ACCOUNT")
}
