package config

import (
	"fmt"
	"os"
)

var (
	ENV                 string
	DatabaseURL         string
	JwtSecret           string
	CloudStorageAccount string
	FirebaseAccount     string
	AssetsName          string
	AssetsPrefix        string
	AssetsEndpoint      string
)

func init() {
	ENV = os.Getenv("ENV")
	DatabaseURL = os.Getenv("DATABASE_URL")
	CloudStorageAccount = os.Getenv("CLOUD_STORAGE_ACCOUNT")
	FirebaseAccount = os.Getenv("FIREBASE_ACCOUNT")
	AssetsName = os.Getenv("ASSETS_NAME")
	AssetsPrefix = os.Getenv("ASSETS_PREFIX")
	AssetsEndpoint = fmt.Sprintf("%s/%s", AssetsPrefix, AssetsName)
}
