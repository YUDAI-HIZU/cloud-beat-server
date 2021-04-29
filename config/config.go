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
	AssetsPrifix        string
	AssetsEndpoint      string
)

func init() {
	ENV = os.Getenv("ENV")
	DatabaseURL = os.Getenv("DATABASE_URL")
	CloudStorageAccount = os.Getenv("CLOUD_STORAGE_ACCOUNT")
	FirebaseAccount = os.Getenv("FIREBASE_ACCOUNT")
	AssetsName = os.Getenv("ASSETS_NAME")
	AssetsPrifix = os.Getenv("ASSETS_PRIFIX")
	AssetsEndpoint = fmt.Sprintf("%s/%s", AssetsPrifix, AssetsName)
}
