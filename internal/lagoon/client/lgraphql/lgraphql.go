// Code generated for package lgraphql by go-bindata DO NOT EDIT. (@generated)
// sources:
// _lgraphql/deployEnvironmentLatest.graphql
// _lgraphql/getDeploymentsByBulkID.graphql
package lgraphql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __lgraphqlDeployenvironmentlatestGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xb1\x6a\xc4\x30\x0c\x86\x77\x3f\xc5\x5f\xe8\x70\x7d\x05\xcf\xbd\x21\xb4\x74\x29\xdc\x52\x3a\xe8\x88\x28\xa2\x8e\x1c\x74\xf2\x41\x28\xf7\xee\x25\x8d\x2f\x71\x3a\xfa\xd7\x67\xe9\xfb\x87\xe2\xe4\x92\x15\x87\x00\x3c\xb2\x5e\xc5\xb2\x0e\xac\x1e\x71\xdc\x1e\x9d\x8e\xc5\x1f\x66\xe2\x5c\xd2\x77\xd7\x47\xbc\xbb\x89\x7e\xdd\x93\x37\x1a\xb8\xcd\x46\x93\x6c\xe2\x53\x44\xa7\xbe\x40\x92\xfa\x13\x99\xd0\x39\xf1\x25\xe2\xe3\xa8\xd7\x17\x9e\x4e\x94\x0a\xff\x2d\xff\x0c\xc0\x13\x7e\x02\x00\xf4\x3c\xa6\x3c\x35\xe7\x5f\xc9\xf9\xe2\x07\x99\xc1\x58\x21\x60\x27\xdb\xaa\xd7\xf9\x5d\xb5\x3a\x37\xe9\xa2\xbb\x9a\xd7\xc9\x26\xbd\xfa\xaf\x7f\xf6\xf6\xff\xea\x54\xca\xd8\x8b\xe9\x33\x39\x45\xb8\x95\x65\xed\x6d\xee\x15\x6e\xbf\x01\x00\x00\xff\xff\xd1\x94\x72\x89\x67\x01\x00\x00")

func _lgraphqlDeployenvironmentlatestGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlDeployenvironmentlatestGraphql,
		"_lgraphql/deployEnvironmentLatest.graphql",
	)
}

func _lgraphqlDeployenvironmentlatestGraphql() (*asset, error) {
	bytes, err := _lgraphqlDeployenvironmentlatestGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/deployEnvironmentLatest.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __lgraphqlGetdeploymentsbybulkidGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2c\x4d\x2d\xaa\x54\xd0\x50\x49\x2a\xcd\xc9\xf6\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x54\xa8\xe6\x52\x50\x50\x50\x48\x49\x2d\xc8\xc9\xaf\xcc\x4d\xcd\x2b\x29\x76\xaa\x74\x02\xab\xd1\x80\x29\x85\xea\xd1\x04\x2b\x84\x28\x07\x81\xcc\x14\x38\x33\x2f\x31\x37\x15\xce\x29\x4a\xcd\xcd\x2f\x49\xf5\x44\xc8\x16\x97\x24\x96\x94\x16\xc3\xb9\x49\xa5\x99\x39\x29\x3e\xf9\xe9\x70\x81\xd4\xbc\xb2\xcc\xa2\xfc\x3c\x90\xe5\x08\xd3\xd1\x6c\xc0\xb0\x05\x04\x0a\x8a\xf2\xb3\x52\x93\xd1\x34\x61\xd1\x88\x55\x73\x2d\x17\x2a\xab\x96\xab\x16\x10\x00\x00\xff\xff\x26\x1b\x9b\xfe\x24\x01\x00\x00")

func _lgraphqlGetdeploymentsbybulkidGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__lgraphqlGetdeploymentsbybulkidGraphql,
		"_lgraphql/getDeploymentsByBulkID.graphql",
	)
}

func _lgraphqlGetdeploymentsbybulkidGraphql() (*asset, error) {
	bytes, err := _lgraphqlGetdeploymentsbybulkidGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_lgraphql/getDeploymentsByBulkID.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"_lgraphql/deployEnvironmentLatest.graphql": _lgraphqlDeployenvironmentlatestGraphql,
	"_lgraphql/getDeploymentsByBulkID.graphql":  _lgraphqlGetdeploymentsbybulkidGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"_lgraphql": &bintree{nil, map[string]*bintree{
		"deployEnvironmentLatest.graphql": &bintree{_lgraphqlDeployenvironmentlatestGraphql, map[string]*bintree{}},
		"getDeploymentsByBulkID.graphql":  &bintree{_lgraphqlGetdeploymentsbybulkidGraphql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
