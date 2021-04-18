package models

import (
	"bytes"
	"fmt"
)

type Audio struct {
	Prefix string
	Name   string
	Buf    *bytes.Buffer
}

func (a *Audio) ObjName() string {
	return fmt.Sprintf("%s/%s", a.Prefix, a.Name)
}
