package packrvfs

import (
	"os"
	"time"
)

type PackrFileInfo struct {
	name string
}

func (finfo PackrFileInfo) Name() string {
	return finfo.name
}

func (finfo PackrFileInfo) Size() int64 {
	return 0
}

func (finfo PackrFileInfo) Mode() os.FileMode {
	return 0777
}

func (finfo PackrFileInfo) ModTime() time.Time {
	return time.Unix(0, 0)
}

func (finfo PackrFileInfo) IsDir() bool {
	return false
}

func (finfo PackrFileInfo) Sys() interface{} {
	return nil
}
