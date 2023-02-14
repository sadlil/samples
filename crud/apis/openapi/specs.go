// Code generated for package openapi by go-bindata DO NOT EDIT. (@generated)
// sources:
// apis/openapi/apis/swagger.json
package openapi

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

var _apisOpenapiApisSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\x1b\x37\x10\xbe\xeb\x57\x0c\xb6\x3d\xb4\x80\xa3\x75\xd2\x20\x40\x7c\x8a\x1b\x19\x81\xd1\x40\x36\xfc\x40\x0f\x4d\x10\x50\xcb\x59\x89\x09\x97\x5c\x73\xb8\x72\xd4\x40\xff\xbd\x20\x57\xab\x7d\xcb\x72\x14\xbb\x0a\x60\xdd\x96\x8f\xe1\x37\xc3\x6f\x1e\x24\xf5\x6d\x00\x10\xd0\x2d\x9b\x4e\xd1\x04\x47\x10\xbc\x18\x1e\x06\x07\xae\x4d\xa8\x58\x07\x47\xe0\xfa\x01\x02\x2b\xac\x44\xd7\x7f\xc9\x92\x54\x22\xbc\xbd\xb8\x1e\xc1\xf1\xf9\x29\xf9\xc1\x00\xc1\x1c\x0d\x09\xad\xdc\x90\xf9\xe1\xf0\xf9\x4a\x0a\x40\x10\x69\x65\x59\x64\xd7\xa2\x00\x02\x4c\x98\x90\x6e\xe4\x97\x7f\x85\xb1\xc9\x9b\xa9\xfb\x1e\x46\x3a\x09\xfc\x88\xe5\x00\x60\xe9\x31\x58\x36\xa5\xe0\x08\xfe\xf1\xcd\xeb\xe9\x8a\x25\x1e\xca\x95\xe6\xfa\x12\xcd\x5c\x44\x58\x4e\xfc\xe8\x27\xce\x34\xb9\x15\x83\x99\xb5\xe9\x51\x18\x4a\x1d\x31\xe9\xda\x8e\x5e\xbd\x7e\xfd\x3a\xd7\x8f\xa2\x19\x26\x58\x8a\xf7\x63\x83\xb5\x84\x48\x2b\xca\x6a\xfd\x2c\x4d\xa5\x88\x98\x15\x5a\x85\x9f\x49\xab\x72\x6c\x6a\x34\xcf\xa2\x2d\xc7\x32\x3b\xa3\xd2\xae\x21\x4b\x45\x38\x7f\x1e\x5a\xcd\x75\xd5\x44\x53\xac\x5a\xcc\xc1\xcd\x92\x84\x99\x85\x53\xea\xbd\x20\xeb\x74\x87\x8b\xf3\xb7\x60\xd0\x1a\x81\x73\x24\x60\x20\x05\x59\xd0\x31\xf8\x4e\x3d\xf9\x8c\x91\xa5\xe1\x6a\x1f\xbc\x10\x9d\xa2\xf1\xa0\x4e\x79\xc3\x80\x9f\x0a\xa1\xd5\xe1\x06\x29\xd5\x8a\x90\x6a\x50\x00\x82\x17\x87\x87\x8d\x26\x80\x80\x23\x45\x46\xa4\x76\xc5\x82\x63\xa0\x2c\x8a\x90\x28\xce\x24\x14\x92\xaa\x68\x72\xb5\xdc\x2e\xb0\x96\x30\x80\xe0\x57\x83\xb1\x93\xf3\x4b\xc8\x31\x16\x4a\x38\xb9\x14\xce\x9f\x17\x40\x2f\x56\x22\x83\xda\xc4\x65\xe5\x6b\x59\x5d\x2b\xe0\x18\xb3\x4c\xda\xbb\x71\x2b\xc8\x14\x7e\x4d\x31\xb2\xc8\x01\x8d\xd1\xe6\xc7\xc1\x37\x69\x74\x69\x99\xcd\x68\x03\xea\x41\x07\xfe\x20\x65\x86\x25\x68\xd1\x94\x24\xcb\x7f\x0d\x65\x0a\xdf\x90\x22\x11\xb6\x89\x56\x78\x05\x6f\x32\x34\x8b\x66\x97\xc1\x9b\x4c\x18\x74\xb4\x88\x99\x24\x6c\x74\xdb\x45\xea\xc5\x92\x35\x42\x4d\x9b\x93\x63\x6d\x12\xe6\xfd\x4d\x28\xfb\xea\x65\xd0\xb7\x07\x3d\x58\x75\x1c\x13\xfe\x2c\x60\x49\x1b\x7b\x66\x38\x9a\x07\xc3\x7b\x5f\x40\x39\x9f\x1e\xc9\x7a\xa8\xb2\xa4\xc1\x40\xdf\x7e\x3d\xfe\x6b\x7c\xf6\xf7\xb8\x31\x1c\x20\x38\x3f\x19\x8f\x4e\xc7\xef\xda\x1d\xa7\xe3\x4f\xe7\x17\x67\xef\x2e\x4e\x2e\x2f\xdb\x9d\xa3\xb3\xf1\x49\x47\xeb\xc9\xfb\x93\xab\x93\x51\xbb\xe3\xf8\xcf\xe3\xb1\x9b\x32\xaa\xbb\xd5\xc7\x83\xa6\xa7\x17\x21\x60\x8d\xf7\x9e\xc6\x4e\x8d\xd0\x46\xd8\x96\x4d\x1f\x64\xf3\x07\x1d\x7a\xd4\x53\xe2\xaa\xad\x95\x08\xfd\xa4\x41\x43\xad\x20\xcd\x93\x62\x77\x52\x79\x6b\x90\x59\x5c\xa7\x95\xc8\x7f\xba\xa4\xa2\xf0\xb6\x9a\x50\xb6\xcd\x27\xa5\xbc\xbd\xcf\x28\x25\xd4\xa7\x9c\xe2\x7f\x3d\xe4\x9f\x68\xde\x43\xfc\xae\x9e\x0a\xef\xad\xc9\x9a\xb4\xdf\x6d\x97\x6e\x32\x24\xbb\x8d\xba\xbb\x79\xce\xa0\x62\xb1\x5a\xa5\x16\x7e\x13\x7c\xb9\x6d\xb9\xf6\x0e\xbb\xab\x35\x12\x6a\x2a\xb1\xea\x5b\x30\x61\x84\x1c\xb4\x02\x61\x09\x4e\x47\xdb\xfa\xda\x6a\x85\xbd\x77\xb4\x15\xce\x27\x2f\xf3\xbf\x1e\x2f\x13\xbc\xdb\xc7\xdc\xc1\xe1\x7e\x3e\xf6\x88\xa9\x85\xa3\x44\x8b\xbd\x2e\x70\x35\x43\x18\xf9\x21\x6b\x4f\xc8\x67\x38\x3f\xf8\x01\x0e\x50\xca\xde\x1b\x1f\x48\x8d\xc3\x6b\x45\xbe\xfa\xf2\x89\xea\x95\x65\x7f\x6a\xaa\xa7\x59\x7f\xa8\xbf\x4e\x79\xb5\x88\xca\xfc\x27\x01\x53\x80\x5f\x05\x59\xa1\xa6\xdf\x53\x4a\x95\x52\xf7\x86\xdd\x7d\x11\xbe\x84\xfa\x14\xe4\xfd\xef\xff\x64\xfe\x16\xe7\x99\xc7\x29\xe9\x0a\x8c\x39\xed\xdb\xc7\xb7\x7a\xac\x6c\xf4\x7a\xeb\x2e\xa4\x66\xbc\xb3\x73\x03\x19\xbd\xc7\xb4\x26\x2c\x07\x9b\xbe\x1f\xae\x8c\x5c\x5f\x6a\x56\x40\x96\xf7\x80\xa9\xd1\x56\x4f\xb2\xf8\x58\x2d\xaa\x75\x65\x8f\xe5\xfa\x2c\x16\xbc\x59\x4d\xa8\x05\x84\x3e\x8a\x2c\x5b\xb1\x8d\x71\xee\x81\x31\x79\xde\x91\xbe\x8a\x4a\xb8\xf4\x92\x1d\x90\x46\x9a\xf7\x02\x15\xca\xe2\xb4\x71\xc3\x52\xbf\xb7\xf9\xe3\x45\xd0\xe9\x87\x09\x12\xb1\xe9\xf6\x16\xa8\x4c\xe5\x68\x99\x90\xad\x68\x5a\x4c\x65\xc6\xb0\xba\x43\x04\xc2\x62\xd2\x66\x6c\x0f\x1d\xab\x1b\xdc\x1d\x4e\xea\x74\x29\xae\xd5\x3b\xce\x3d\x3b\x58\xbd\x71\xcd\xbc\x01\x6f\xd3\x7d\xb6\x81\xb7\x8a\xf8\x7b\x85\xaf\x79\xe6\xd8\x2b\x70\xad\xbb\xec\x1d\xd1\x3d\x18\x7b\xdb\xc1\xf4\x6e\xe2\x5e\x35\x9e\x34\xee\xab\x90\x68\x06\xfc\x4d\x57\x93\xcd\x32\x21\xcb\x04\x87\x4c\x89\x9b\x0c\xe5\x02\x0c\xa6\x06\x09\x95\x75\xb5\xd8\xa2\x76\xe2\x10\x0a\xec\x0c\x81\xac\x36\x6c\x8a\x43\x70\x47\x95\x39\x93\x19\x82\x8e\xe1\x83\x12\x1c\x6e\x85\x94\x30\x41\x98\xa2\x72\x25\x1a\x72\x98\x2c\xf2\x39\x0b\xb2\x98\xc0\xed\x4c\x48\xf4\x0d\x6e\x0b\x40\x10\x4c\xd0\x55\x7a\x91\x1f\x3c\xec\x8e\x36\xab\x04\xfc\xbd\xea\x5d\x13\x1a\x48\x8d\x9e\x0b\x8e\x1c\xb8\xa0\x54\xb2\x05\x38\xa1\x0e\x76\x81\xa5\x67\xed\xba\xac\x7b\x07\xca\xf5\x1d\xe8\x8e\xf0\x3d\xbf\x90\x43\x21\xaf\x40\x7e\x75\x36\x3a\x3b\x80\x88\x29\xd0\x4a\x2e\x9c\xe5\x09\x2d\x30\xc8\xe5\x96\x9b\x73\xfe\xdb\xe1\xb3\x97\xbf\xf7\xe8\x98\xdf\x5f\xf2\xe3\x66\x41\xb9\x11\x65\x99\x62\x5c\x15\xfb\xcc\x8a\x04\x37\xaa\x91\xc7\x3e\x0e\xcc\x82\x1b\x4b\x96\x25\xe9\x41\x95\x6c\xb7\x33\xcc\xe9\x85\xca\x9a\x45\x95\x1a\x58\x25\x52\x46\x68\x7a\xf7\x8a\x71\x29\xd4\xce\x5c\x29\x8c\x5d\xc8\x03\x4a\x31\x12\xb1\x40\xee\x3c\x80\x67\xf9\xd9\xe3\x00\x84\xe2\x22\xf2\x87\x16\x4f\x22\x91\xa0\x03\xb9\xd6\xc3\x32\xfa\xf2\x41\xd1\x4c\x67\x92\xbb\x9d\x89\x74\x92\xba\xa3\x77\x1f\xcd\xa9\x59\x27\xc0\x5d\x31\xe6\xb2\xfd\x9c\xd1\x54\x28\x1f\xd2\x47\xf4\x56\x51\xd3\x98\xed\x9d\x7f\x55\x27\xd4\x02\x43\xfd\x68\xd6\x11\xcc\xfa\x8b\x9e\xc6\x2e\xb4\x5f\x49\x3a\xde\x47\x3a\x5e\x46\x7a\xde\x44\x9a\xaf\x21\x1d\xef\x20\xed\x17\x90\x8f\x15\xfd\x5b\xaf\x1e\x9b\x6c\xb3\xb2\xae\x53\x61\xc5\x9a\x9c\x0b\xa9\x26\x12\x13\xe9\x42\xa5\xe7\x87\x8e\xef\x30\x59\xc7\x59\x70\x1f\x92\xef\xba\x0e\xc7\xaf\x16\x8d\x62\x72\xa4\xa3\x4a\x21\x9e\x19\x59\xfc\x71\x80\x8e\xc2\x70\x2a\xec\x2c\x9b\x0c\x23\x9d\x84\xe4\x3c\x47\x86\xe4\xff\x00\x41\xa1\x35\x88\x61\xc2\x84\x0a\x23\x93\x71\xb7\xdc\x72\xb0\x1c\xfc\x17\x00\x00\xff\xff\x28\xfa\x0b\x94\x4f\x21\x00\x00")

func apisOpenapiApisSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apisOpenapiApisSwaggerJson,
		"apis/openapi/apis/swagger.json",
	)
}

func apisOpenapiApisSwaggerJson() (*asset, error) {
	bytes, err := apisOpenapiApisSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "apis/openapi/apis/swagger.json", size: 8527, mode: os.FileMode(420), modTime: time.Unix(1676334762, 0)}
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
	"apis/openapi/apis/swagger.json": apisOpenapiApisSwaggerJson,
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
	"apis": &bintree{nil, map[string]*bintree{
		"openapi": &bintree{nil, map[string]*bintree{
			"apis": &bintree{nil, map[string]*bintree{
				"swagger.json": &bintree{apisOpenapiApisSwaggerJson, map[string]*bintree{}},
			}},
		}},
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
