// source from file,implement source interface
package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/adnpa/fern/pkg/common/config"
	"github.com/adnpa/fern/utils"
)

// NewSource new a file source.
func NewSource(path string) config.Source {
	return &file{path: path}
}

type file struct {
	path string
}

// Load implements the config.Source interface.
func (f *file) Load() ([]*config.KV, error) {
	var res []*config.KV
	if utils.IsDir(f.path) {
		return f.loadDir(f.path)
	}
	kv, err := f.loadFile(f.path)
	if err != nil {
		return nil, err
	}
	res = append(res, kv)
	return res, nil
}

func (f *file) loadFile(filePath string) (*config.KV, error) {
	fi, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(fi)
	if err != nil {
		return nil, err
	}
	return &config.KV{
		Key:    filePath,
		Value:  data,
		Format: utils.FileTypeByExt(filePath),
	}, nil
}

func (f *file) loadDir(path string) (kvs []*config.KV, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, fi := range files {
		if fi.IsDir() || strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		kv, err := f.loadFile(filepath.Join(path, fi.Name()))
		if err != nil {
			return nil, err
		}
		kvs = append(kvs, kv)
	}
	return
}
