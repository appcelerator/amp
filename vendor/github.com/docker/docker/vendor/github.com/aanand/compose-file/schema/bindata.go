// Code generated by go-bindata.
// sources:
// data/config_schema_v3.0.json
// DO NOT EDIT!

package schema

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataConfig_schema_v30Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x4d\x93\xdb\x28\x13\xbe\xfb\x57\x4c\x29\xb9\xc5\x33\x93\xaa\x37\xf5\x56\x6d\x6e\x7b\xdc\xd3\xee\x79\x5d\x8a\x0a\x4b\xd8\x26\x23\x09\x02\xc8\x89\x93\xf2\x7f\x5f\x10\x92\x0c\x88\x2f\xdb\x4a\x66\x0f\x3b\x87\xa9\x19\xe8\x6e\xfa\xe3\xa1\x69\x1a\xfd\x58\x3d\x3c\x64\x6f\x59\x79\x80\x0d\xc8\x3e\x3e\x64\x07\xce\xc9\xc7\xe7\xe7\xcf\x0c\xb7\x8f\x6a\xf4\x09\xd3\xfd\x73\x45\xc1\x8e\x3f\xbe\xff\xf0\xac\xc6\xde\x64\x6b\xc9\x87\x2a\xc9\x52\xe2\x76\x87\xf6\x85\x9a\x29\x8e\xff\x7b\x7a\xff\x24\xd9\x15\x09\x3f\x11\x28\x89\xf0\xf6\x33\x2c\xb9\x1a\xa3\xf0\x4b\x87\x28\x94\xcc\x9b\xec\x08\x29\x43\x82\x3a\x5f\xaf\xe4\x1c\xa1\x98\x40\xca\x11\x64\x62\x56\x2a\x27\xc6\x46\x92\x71\x40\x13\xcb\x38\x45\xed\x3e\xeb\x87\xcf\xbd\x04\x31\xc9\x20\x3d\xa2\x52\x93\x30\xa9\xfa\xe6\xf9\x22\xff\x79\x22\x5b\xdb\x52\x35\x65\xfb\x71\x02\x38\x87\xb4\xfd\x6b\xae\x5b\x3f\xfd\x69\x03\x1e\xbf\xff\xfe\xf8\xf7\xfb\xc7\xdf\x9e\x8a\xc7\xfc\xdd\x5b\x63\x5a\xfa\x97\xc2\x9d\x5a\xbe\x82\x3b\xd4\x22\x2e\xac\x99\xd6\xcf\x26\xca\xf3\xf0\xd7\x79\x5a\x18\x54\x55\x4f\x0c\x6a\x63\xed\x1d\xa8\x19\x34\x6d\x6e\x21\xff\x8a\xe9\x4b\xcc\xe6\x89\xec\x95\x6c\x1e\xd6\x77\xd8\x6c\x9a\x73\xc4\x75\xd7\x44\x23\x38\x52\xbd\x92\x31\x6a\xf9\xfb\xe2\xb7\x1a\x8d\x0e\xd2\x2a\x0a\x6d\xed\x5e\x41\x03\xed\x2e\x57\xb9\xd0\xe6\xf7\xd5\xe4\x2c\x8f\x97\x2a\x48\x6a\x7c\x92\x63\x1e\x7f\x28\x82\x06\xb6\x3c\x9b\x5c\x20\xf8\xb6\x1d\xaa\x2b\xdb\xa3\xb8\x85\x7f\x4a\x11\x1b\x6d\xf0\x41\x48\xb6\x36\xb6\x26\xa7\x9f\x37\xfe\xf3\x07\x7c\x9a\xf7\xd8\x32\xcd\x8b\xdc\xc5\xe1\x37\xde\x1b\x15\x5e\x5a\xb9\x00\x97\x2f\x90\xee\x50\x0d\x53\x39\x00\xdd\xb3\x80\xcb\x6a\xc4\x78\x81\x69\x51\x21\xa1\xfd\xd9\x62\x9f\xc9\x8b\xe3\x69\x62\xd5\xfe\xcb\x57\x0e\x81\x59\x09\x48\x21\xc4\x19\x76\x00\x4a\xc1\x29\x5b\x0b\x00\x71\xd8\x30\xb7\x89\x0f\x59\xd7\xa2\x2f\x1d\xfc\x63\x20\xe1\xb4\x83\xb6\xdc\x4a\x28\xb7\xbc\xe0\x3d\xc5\x1d\x29\x08\xa0\x12\x60\x61\xf7\x8b\xb8\x36\x0d\x68\x97\x42\xdd\x35\x76\x24\x78\x5e\x60\x0e\xa0\x16\xd2\xa2\x05\x4d\x0c\x48\x72\xd7\xc1\xb6\x62\x85\x3a\xff\x82\x30\xda\x15\x8a\x9f\x59\x02\xa6\xc3\x70\xd1\x78\x54\x6d\x08\xd8\x4a\x8c\x84\xb6\xd4\x2d\xb3\x18\x0b\x06\x01\x2d\x0f\x37\xf2\xe3\x46\xb8\x2f\xc5\x77\x02\x28\xf4\x44\x30\x52\x78\xf9\xd7\x01\x01\xb6\xc7\x62\xca\x25\x57\xbb\x41\x70\x23\x8a\xdb\x66\xdc\x0d\x29\x09\x66\x4a\xf2\x92\xff\x1b\xc1\x0c\xda\x8e\xb1\x0c\xd4\xa7\x26\x53\x0d\x9f\x8c\x1c\x9b\xd1\x70\xe1\x94\xb6\x6b\xb6\x90\xca\x92\xce\xa0\xdc\x61\xda\x00\xa9\xec\xb8\xb6\x36\x6d\x78\xda\x81\x3c\xdd\x81\xba\x0d\xf2\x58\x07\xb5\xf0\x4e\xfb\xb2\x3c\xc4\x85\x78\x0a\x8a\x03\x66\x3c\x3d\x87\x6b\xec\x07\x08\x6a\x7e\x10\x65\x71\xf9\x12\x60\xd7\xa9\x0c\x6e\xb1\x6c\x0a\xc8\x51\x03\xf6\x71\x22\x52\xc6\x48\x6a\xb0\x85\xf5\x4d\x76\x2e\xea\x7c\x4d\x2c\xde\xef\x25\xa9\x0f\x71\xb3\xca\x65\x98\x8e\x9d\xf9\x15\x45\xe2\x46\x91\x7a\x80\x63\x72\x29\xb8\xec\xc9\x78\x01\xa2\x14\x0a\x56\x9f\x06\xe9\xa7\x27\x55\x7c\x06\x76\x55\xff\x57\x5d\x67\xb9\x5d\x2e\xc8\x9f\xf9\x98\x39\x62\x59\x98\x56\x50\x18\x51\x69\x40\x29\xeb\x06\x0a\x99\x27\xae\x17\xd2\xa1\xd8\x2f\x1a\x5c\xf9\x00\x3a\x23\xb6\x7d\xe3\xcd\xd4\x57\x1f\x84\x3d\xdb\xd5\xf5\x63\x52\xe8\xa2\x17\x88\x88\x35\x3e\xf5\x52\xd5\xbc\xa8\x1b\x87\x58\x4f\x07\x6a\x04\x18\x8c\x6f\x76\xaf\x23\x0d\x69\x88\x1c\x3f\x24\x62\xc2\xc5\xfb\xff\x20\xaf\x87\xd5\x2b\x33\xbd\x46\x8e\x88\xba\xa8\xd2\x6f\x37\x97\x22\x79\x64\xb7\xfd\xe4\x12\x9e\xa0\xca\x9f\x2b\xfa\x0c\xa1\x6f\x30\x82\x29\x9f\xed\xae\x5f\x73\xdc\xab\xa5\xef\x3e\xed\x89\x48\xdc\xa2\x5c\xda\x43\xf3\xd6\xb2\xc5\xb8\x86\xa0\x35\x52\x0f\x85\xa0\x12\x25\x73\x7d\x4a\xa0\x64\x1c\xd0\xe8\x85\x82\xc1\xb2\xa3\x88\x9f\x0a\x71\x1e\x2c\x5e\x67\xb0\x43\x53\x30\xf4\x1d\x9a\xd1\xbc\xe4\xfb\x41\x50\x6e\xf0\xf0\x0a\xb5\x42\x1b\xd8\x46\x4d\x64\x1c\x13\x21\x7f\x2f\x30\x17\x35\x53\x92\xee\x29\x28\x61\x21\xb0\x89\x70\xe5\x62\x58\xeb\xb1\xad\x3a\x0a\x24\x9e\x0d\x31\xbc\x21\xbb\x1b\x6f\x07\x9c\xc7\x63\xd6\xd5\xa8\x41\x7e\x30\x3b\xb2\x64\x42\x22\x57\x49\xdc\x9d\xbb\x03\x79\xfb\xa2\xa9\xb8\x66\x08\x6c\x52\x57\xba\x0b\x94\x0e\xe1\xca\x21\xa1\x64\x38\x00\x6a\x46\x29\xa0\x47\xcf\xc0\xf0\x8e\xbb\x19\x5c\x05\x85\x53\x2f\xa3\x83\xdb\xcb\x5b\x0f\x8a\xe4\x4e\xfa\xab\x72\xb2\xad\x46\xee\x4d\x8b\x67\x67\x5a\xec\x58\xb4\xba\xd3\xfb\x8b\x8b\xee\x64\x59\xc2\x48\x64\x57\xc8\xad\xc2\xca\x52\xf7\x8a\x0e\xaf\x75\x9b\x18\x05\xb8\x7a\x7d\x3a\xa9\xdd\xef\xdb\x4c\x80\x1b\x4f\x89\x4b\x97\xd4\xd3\xf8\x93\xf8\xa0\x47\x23\x79\xb8\x7c\xca\x51\x03\x71\xc7\x23\x54\x14\x8a\x31\xcb\xf3\x43\xa6\x33\x84\x89\xb4\x9c\x5a\x0a\xfe\xd2\x4b\x7b\x85\x18\xd8\x5a\xfd\xbf\x29\x47\xdd\x14\x5e\x25\xf6\xd2\x3b\x8d\x04\x57\xa3\x5c\x20\xb6\x81\xda\x5c\x0b\x19\xa9\x51\x09\x58\x2c\xcb\xdc\x71\x85\xec\x48\x05\x38\x2c\xd4\x53\xd2\x55\x79\x3d\x90\xd0\x09\xa0\xa0\xae\xa1\x58\xb4\x49\x49\x90\x22\x06\x35\x38\xdd\x74\xe0\xf5\xec\x3b\x80\xea\x8e\xc2\x02\x94\x7c\x78\xad\x8a\x20\x53\x38\x5f\x38\x06\x3b\x33\x45\xda\x92\x0d\xf8\x56\x8c\xcb\xf6\x24\xce\x6d\xe5\x2d\xbc\x52\x6f\x7f\x1a\x12\x18\xee\x68\x39\x73\xf6\xcd\x21\xba\x1c\xe4\x1e\xc4\x8c\x2b\xce\x4c\x17\x13\x32\x29\x4d\x97\xf3\x28\x7f\xf4\xdc\x18\x2a\xc1\x82\x60\x81\xf6\xd3\x52\x16\x0a\x48\x2b\x27\xa7\x00\xe2\x4e\x04\x4a\x38\xc8\x3a\xa7\x21\x3c\xba\x59\x7b\x86\xaf\xa8\xad\xf0\xd7\x2b\x16\x5c\x0e\x4a\xa4\x16\x45\xa6\x95\xef\xee\x75\xb4\xd0\x1d\x08\x53\xaf\x3e\xd6\xef\x35\xeb\x8e\x53\x7d\xc2\x67\x24\xeb\x4f\x74\xf1\xb7\x4e\x4f\xa6\x2f\x49\x17\xed\xd8\x34\xb0\xc1\xd4\x09\xc0\x80\x8d\x89\x4f\xd3\x31\x0b\x47\xb2\x05\x0e\xb5\xa4\x0e\xdf\x40\x25\x2f\x74\x8b\xdf\x24\xe2\x5d\xbc\x3c\x9e\x8f\x10\x01\xcd\x52\x9b\x23\xb9\xe7\x99\x39\x8f\x60\x63\xed\x79\xaf\x40\xa9\xeb\xec\x17\xc4\xb4\x8e\xeb\x3e\x50\xb0\x6e\x2b\x10\x12\x82\xe6\xe5\xc7\xf9\x10\x9b\x7e\x05\x39\xfb\x2f\x1c\xf7\xe5\xbc\xf1\xb9\xc2\x13\xd5\xcd\x54\x48\xae\x27\x5f\xe5\xc9\x21\xf6\xbe\x15\x2c\xa7\xff\x95\xf5\xdd\x1d\x69\x71\xf8\xb4\x22\x92\x32\x06\xaa\xff\x32\xc6\x20\xe5\xf5\xf1\x15\x38\x13\x6f\xbc\x1c\x5c\x01\x1a\xab\xab\xa4\x81\x67\x7e\x73\x0c\xc5\x39\xb9\x27\x3e\x70\xe4\xa6\x1a\x36\xd9\xc7\xf9\x67\x6b\x66\x0a\x0d\x35\x1c\x46\x12\x4f\x8f\xd4\x5a\x74\x70\x5e\xd8\xf2\x05\x61\xfb\xf4\x2e\x70\x50\x84\xde\xae\x7e\x52\x86\x5d\xa0\x99\xe3\x8e\xa9\x55\x5c\x8e\xde\x9d\x7f\x7b\xe5\xc9\x54\x1a\xff\xec\x4b\x2c\x69\x67\x7b\x9a\x75\x36\x7e\x98\x5d\x36\xf5\x15\x55\x6e\xf8\xc7\x22\x51\x2f\xc1\x5a\x9e\xc8\xf5\x7a\xdb\x17\x46\xe7\xf7\x59\x76\x8f\x6f\xfc\x4e\x2a\x0f\x6f\xf6\xd5\xf8\xfb\xbc\x3a\xaf\xfe\x09\x00\x00\xff\xff\x37\x89\x5b\xf1\x5b\x2a\x00\x00")

func dataConfig_schema_v30JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v30Json,
		"data/config_schema_v3.0.json",
	)
}

func dataConfig_schema_v30Json() (*asset, error) {
	bytes, err := dataConfig_schema_v30JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.0.json", size: 10843, mode: os.FileMode(420), modTime: time.Unix(1479392593, 0)}
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
	"data/config_schema_v3.0.json": dataConfig_schema_v30Json,
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
	"data": &bintree{nil, map[string]*bintree{
		"config_schema_v3.0.json": &bintree{dataConfig_schema_v30Json, map[string]*bintree{}},
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

