// Code generated by go-bindata.
// sources:
// ui/index.html
// ui/work.css
// ui/work.js
// DO NOT EDIT!

package webui

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

var _uiIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x8e\x41\x0e\x02\x31\x08\x45\xf7\x73\x0a\xc2\x01\x24\xee\x69\xef\x62\x2a\x86\x3a\x55\x1b\xc0\xe8\xdc\xde\x3a\x55\x76\xbc\xbc\xe4\x7d\xd6\xb8\xb5\xbc\xb0\xca\xe9\x9c\x17\x18\xc7\x5e\xac\xf6\x00\xb7\x92\x90\x5e\x0f\x5b\x0f\x57\xc7\xcc\x34\xf9\x4f\x6a\xf5\xbe\x82\x9a\x5c\xfe\x4e\x71\x47\x30\x69\x09\x3d\xb6\x26\xae\x22\x81\x10\x5b\x97\x84\x21\xef\xa0\x5d\xa0\x91\xa2\xd9\x62\x3d\x66\x7f\xf6\xf1\xd6\x1d\x7e\x77\x7c\x02\x00\x00\xff\xff\xd0\x47\xe1\x94\x8e\x00\x00\x00")

func uiIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_uiIndexHtml,
		"ui/index.html",
	)
}

func uiIndexHtml() (*asset, error) {
	bytes, err := uiIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/index.html", size: 142, mode: os.FileMode(420), modTime: time.Unix(1466623418, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiWorkCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x30\x54\xa8\xe6\xe2\x4c\xce\xcf\xc9\x2f\xb2\x2a\x4a\x4d\xb1\xe6\xaa\x05\x04\x00\x00\xff\xff\x2c\xdf\x7e\xda\x12\x00\x00\x00")

func uiWorkCssBytes() ([]byte, error) {
	return bindataRead(
		_uiWorkCss,
		"ui/work.css",
	)
}

func uiWorkCss() (*asset, error) {
	bytes, err := uiWorkCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/work.css", size: 18, mode: os.FileMode(420), modTime: time.Unix(1466623227, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiWorkJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xcc\x49\x2d\x2a\xd1\x50\xcf\xc9\x4f\x4c\x49\x4d\x51\xd7\x04\x04\x00\x00\xff\xff\xbf\xb2\x83\x2b\x0f\x00\x00\x00")

func uiWorkJsBytes() ([]byte, error) {
	return bindataRead(
		_uiWorkJs,
		"ui/work.js",
	)
}

func uiWorkJs() (*asset, error) {
	bytes, err := uiWorkJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/work.js", size: 15, mode: os.FileMode(420), modTime: time.Unix(1466623215, 0)}
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
	"ui/index.html": uiIndexHtml,
	"ui/work.css": uiWorkCss,
	"ui/work.js": uiWorkJs,
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
	"ui": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{uiIndexHtml, map[string]*bintree{}},
		"work.css": &bintree{uiWorkCss, map[string]*bintree{}},
		"work.js": &bintree{uiWorkJs, map[string]*bintree{}},
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
