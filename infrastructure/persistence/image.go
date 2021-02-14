package persistence

import (
	"context"
	"fmt"
	"io"
)

type ImagePersistence struct {
	file      *FilePersistence
	keyPrefix string
}

func NewImagePersistence(keyPrefix string) *ImagePersistence {
	return &ImagePersistence{
		file:      NewFilePersistence(context.Background()),
		keyPrefix: fmt.Sprintf("images/%s", keyPrefix),
	}
}

func (i *ImagePersistence) Upload(fileName string, imageFile io.Reader) {
	key := fmt.Sprintf("%s/%s", i.keyPrefix, fileName)
	i.file.Upload(imageFile, key)
}
