package packrvfs

import (
	"github.com/gobuffalo/packr/v2"
	"golang.org/x/tools/godoc/vfs"
	"os"
	"path/filepath"
)

type PackrVfs struct {
	packrBox *packr.Box
}

func NewPackrVfs(packrBox *packr.Box) *PackrVfs {
	return &PackrVfs{
		packrBox: packrBox,
	}
}

func (pvfs *PackrVfs) PackrBox() *packr.Box {
	return pvfs.packrBox
}

func (pvfs *PackrVfs) Open(name string) (vfs.ReadSeekCloser, error) {
	return pvfs.packrBox.Open(name)
}

func (pvfs *PackrVfs) Lstat(path string) (os.FileInfo, error) {
	return PackrFileInfo{
		name: path,
	}, nil
}

func (pvfs *PackrVfs) Stat(path string) (os.FileInfo, error) {
	return PackrFileInfo{
		name: path,
	}, nil
}

func (pvfs *PackrVfs) ReadDir(path string) ([]os.FileInfo, error) {
	files := pvfs.packrBox.List()
	fileInfos := make([]os.FileInfo, 0)

	for _, filePath := range files {
		if filepath.Dir(filePath) == filepath.Clean(path) {
			fileInfos = append(fileInfos, PackrFileInfo{
				name: filePath,
			})
		}
	}

	return fileInfos, nil
}

func (pvfs *PackrVfs) RootType(path string) vfs.RootType {
	return ""
}

func (pvfs *PackrVfs) String() string {
	return ""
}
