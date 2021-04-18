package models

import (
	"bytes"
	"fmt"
)

type Image struct {
	Prefix string
	Name   string
	Buf    *bytes.Buffer
}

func (i *Image) ObjName() string {
	return fmt.Sprintf("%s/%s", i.Prefix, i.Name)
}
