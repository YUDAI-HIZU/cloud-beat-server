package persistence

// import (
// 	"app/config"
// 	"app/infrastructure/gcp"
// 	"context"
// 	"fmt"
// 	"io"

// 	"cloud.google.com/go/storage"
// )

// type FilePersistence struct {
// 	ctx        context.Context
// 	bucketName string
// 	rootPrefix string
// 	client     *storage.Client
// }

// func NewFilePersistence(ctx context.Context) *FilePersistence {
// 	client := gcp.InitClient()
// 	return &FilePersistence{
// 		ctx:        ctx,
// 		bucketName: config.BucketName,
// 		rootPrefix: config.ENV,
// 		client:     client,
// 	}
// }

// func (f *FilePersistence) Upload(file io.Reader, key string) {
// 	fmt.Println("=======file upload======")
// 	path := fmt.Sprintf("%s/%s", f.rootPrefix, key)
// 	writer := f.client.Bucket(f.bucketName).Object(path).NewWriter(f.ctx)
// 	writer.ContentType = "image/png"
// 	_, err := io.Copy(writer, file)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = writer.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
