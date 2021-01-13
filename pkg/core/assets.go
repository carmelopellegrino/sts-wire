// Code generated for package core by go-bindata DO NOT EDIT. (@generated)
// sources:
// data/html/errorNoCred.html
// data/html/errorNoSaveToken.html
// data/html/errorNoStsCred.html
// data/html/errorNoToken.html
// data/html/errorTokenExpired.html
// data/html/mountingPage.html
package core

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

var _htmlErrornocredHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x51\x6f\xd3\x30\x10\xc7\x9f\xd3\x4f\x71\xeb\xc4\x1b\x6e\xda\x49\xc0\xc8\xdc\x48\xd3\x36\x04\x12\xa8\x88\x16\xa1\x3d\xba\xf6\xb5\x36\x73\xed\x60\x5f\xd8\x22\xb4\xef\x8e\x9c\x34\x6d\x8a\xf6\xc0\x4b\xd3\x5c\xee\xff\xf3\xdd\xff\xce\xfc\xec\x76\x71\xb3\xba\xff\x7a\x07\x1f\x57\x5f\x3e\x97\x23\xae\x69\x67\xcb\x51\xc6\x35\x0a\x95\x9e\x64\xc8\x62\xb9\x5c\x2d\xd9\x8f\x4f\xdf\xee\x78\xde\xbd\x8f\x32\xbe\x43\x12\xa0\x89\x2a\x86\xbf\x6a\xf3\x7b\x3e\xbe\xf1\x8e\xd0\x11\x5b\x35\x15\x8e\x41\x76\x6f\xf3\x31\xe1\x13\xe5\x89\x7a\x05\x52\x8b\x10\x91\xe6\xdf\x57\x1f\xd8\xe5\x18\xf2\x84\x89\xd4\xb4\xbc\x2c\xa5\xbc\x1e\x65\xd9\xda\xab\x06\xfe\xa4\x00\x9a\xad\xa6\x02\x66\xd3\xe9\xab\xab\x51\x96\x3d\x0f\x3f\x2a\x13\x2b\x2b\x9a\x02\x36\x16\x9f\xd2\xd7\xf4\x64\x8f\x41\x54\x05\xa4\xdf\x14\xda\x89\xb0\x35\xae\x80\x69\xaf\x9e\xa4\xa6\x30\xb0\x1d\xba\x3a\x1d\xb5\xf1\x9e\x30\xbc\xcc\x13\xd6\x6c\x1d\x33\x84\xbb\x58\x80\x44\x47\x18\x52\xf8\xd1\x28\xd2\xa7\x35\x0d\xa9\x2d\xeb\x67\x1d\xc9\x6c\x1a\xb6\xb7\x60\x28\xef\x7b\x7a\x3b\xad\xda\x53\xd6\x42\x3e\x6c\x83\xaf\x9d\x2a\xe0\x7c\x26\x2f\xdf\xc9\xf7\x29\x2c\xbd\xf5\xa1\x80\xf3\xcd\x66\xd3\x9f\xa2\x2f\x5a\xf6\xa1\x27\x98\xc2\xe5\xbf\x0c\xb6\xd7\x05\x54\xbd\xac\xb6\x60\xcd\x69\x87\xc6\x59\xe3\x90\xad\xad\x97\x0f\x29\xad\x12\x4a\x19\xb7\x4d\xcc\xd9\xbe\x2c\x6b\x22\xb1\x76\x32\x05\x38\xef\xb0\xa7\x45\x94\x64\xbc\x6b\x79\xc9\xa8\x02\x66\x03\x53\xde\x74\x9e\x1c\x78\x3d\x2d\x29\x45\x20\x23\x2d\x9e\xf4\x20\x6a\xf2\xff\x21\x4f\x0b\xc4\xda\x71\x0c\x9d\x7c\x3e\x9d\xdf\x0b\x4d\x9c\x78\xab\x54\x6f\x09\xcf\xfb\x95\xe3\x79\xbf\xe4\x69\xad\xfa\xa5\xc7\x00\xd2\x8a\x18\xe7\xe3\xc1\x5c\xc7\x69\x43\xb9\x9e\x0d\x2e\x82\x9e\x1d\x10\x18\xda\x4d\xee\xcc\x69\x33\xf7\xed\x76\xaa\x3e\x23\xfd\xbf\x28\xef\x42\xf0\x81\xe7\xfa\xa2\x8b\xc4\x9d\xb0\xb6\x5c\x5c\xd7\xa4\xa1\x0a\x5e\x62\x8c\x3c\xef\x82\x49\x7c\xe4\x67\xbc\x2a\x6f\x84\x73\x9e\x20\xa0\x50\x20\x7d\x08\x28\xc9\x36\x40\x1a\x41\x06\x54\xe8\xc8\x08\x1b\x27\x93\x09\x5f\x87\xf2\xde\xd7\x20\x85\x03\x69\x7d\x44\x20\x6d\x22\x90\x58\x9f\xf1\xbc\xea\xc0\xc7\x12\x79\x7e\x2c\x9d\x77\x9e\xb6\x29\x5d\x15\xb7\x8b\xdb\xeb\x25\x5b\x2d\x8f\x55\xf1\xfc\x90\xc4\xf3\xce\x3a\xde\xde\xf0\xf2\x6f\x00\x00\x00\xff\xff\x59\x8a\x6e\xd7\x4e\x04\x00\x00")

func htmlErrornocredHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlErrornocredHtml,
		"html/errorNoCred.html",
	)
}

func htmlErrornocredHtml() (*asset, error) {
	bytes, err := htmlErrornocredHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/errorNoCred.html", size: 1102, mode: os.FileMode(420), modTime: time.Unix(1610538990, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlErrornosavetokenHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x41\x6f\xdb\x3c\x0c\x86\xcf\xce\xaf\x60\x53\x7c\xb7\x4f\x71\x52\x60\x5b\xe7\x2a\x06\x8a\xb6\xc3\x06\x6c\xe8\xb0\x78\x18\x7a\x54\x64\x3a\xd2\x22\x4b\x9e\x44\xb7\x0d\x86\xfe\xf7\x41\x76\x9c\x38\x43\x0f\xbb\xc4\x31\xcd\xf7\x11\xf9\x92\xe2\x67\xb7\xf7\x37\xc5\xc3\xd7\x3b\xf8\x58\x7c\xf9\x9c\x4f\xb8\xa2\xda\xe4\x93\x84\x2b\x14\x65\x7c\x92\x26\x83\xf9\xaa\x58\xb1\x1f\x9f\xbe\xdd\xf1\xb4\x7f\x9f\x24\xbc\x46\x12\xa0\x88\x1a\x86\xbf\x5a\xfd\xb8\x9c\xde\x38\x4b\x68\x89\x15\xbb\x06\xa7\x20\xfb\xb7\xe5\x94\xf0\x99\xd2\x48\xbd\x02\xa9\x84\x0f\x48\xcb\xef\xc5\x07\x76\x39\x85\x34\x62\x02\xed\x3a\x5e\x12\x53\xfe\x9f\x24\xc9\xda\x95\x3b\xf8\x1d\x03\xa8\x37\x8a\x32\x58\xcc\xe7\xff\x5d\x4d\x92\xe4\x65\xfc\xb1\xd4\xa1\x31\x62\x97\x41\x65\xf0\x39\x7e\x8d\x4f\xf6\xe4\x45\x93\x41\xfc\x8d\xa1\x5a\xf8\x8d\xb6\x19\xcc\x07\xf5\x2c\x36\x85\x9e\xd5\x68\xdb\x78\x54\xe5\x1c\xa1\x7f\x9d\x27\x8c\xde\x58\xa6\x09\xeb\x90\x81\x44\x4b\xe8\x63\xf8\x49\x97\xa4\x4e\x6b\x1a\x53\x3b\xd6\xcf\x36\x90\xae\x76\x6c\x6f\xc1\x58\x3e\xf4\xf4\x76\xde\x74\xa7\xac\x85\xdc\x6e\xbc\x6b\x6d\x99\xc1\xf9\x42\x5e\xbe\x93\xef\x63\x58\x3a\xe3\x7c\x06\xe7\x55\x55\x0d\xa7\xa8\x8b\x8e\x7d\xe8\x09\xe6\x70\xf9\x37\x83\xed\x75\x1e\xcb\x41\xd6\x1a\x30\xfa\xb4\x43\x6d\x8d\xb6\xc8\xd6\xc6\xc9\x6d\x4c\x6b\x44\x59\x6a\xbb\x89\xcc\xc5\xbe\x2c\xa3\x03\xb1\x6e\x32\x19\x58\x67\x71\xa0\x05\x94\xa4\x9d\xed\x78\xd1\xa8\x0c\x16\x23\x53\xde\xf4\x9e\x1c\x78\x03\x2d\x2a\x85\x27\x2d\x0d\x9e\xf4\x20\x5a\x72\xff\x20\x8f\x0b\xc4\xba\x71\x8c\x9d\x7c\x39\x9d\xdf\x2b\x4d\x9c\x78\x5b\x96\x83\x25\x3c\x1d\x56\x8e\xa7\xc3\x92\xc7\xb5\x1a\x96\x1e\x3d\x48\x23\x42\x58\x4e\x47\x73\x9d\xc6\x0d\xe5\x6a\x31\xba\x08\x6a\x71\x40\xa0\xef\x36\xb9\x37\xa7\xcb\xdc\xb7\xdb\xab\x86\x8c\xf8\xff\x22\xbf\xf3\xde\x79\x9e\xaa\x8b\x3e\x12\x6a\x61\x4c\x7e\x7f\xdd\x92\x82\xc6\x3b\x89\x21\xf0\xb4\x0f\x46\xf1\x91\x9f\xf0\x26\xbf\x71\xad\x29\xc1\x3a\x82\x20\x1e\x11\xc8\x6d\xd1\x42\xa5\x0d\xce\x66\x33\xbe\xf6\xf9\x83\x6b\x41\x0a\x0b\xd2\xb8\x80\x40\x4a\x07\x20\xb1\x3e\xe3\x69\xd3\xc3\x8e\x65\xf1\xf4\x58\x2e\xef\x7d\xec\x52\xfa\x93\x6f\xef\x6f\xaf\x57\xac\x58\x1d\x2b\xe1\xe9\x21\x89\xa7\xbd\x5d\xbc\xbb\xd5\xf9\x9f\x00\x00\x00\xff\xff\x16\x19\xf6\xe6\x42\x04\x00\x00")

func htmlErrornosavetokenHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlErrornosavetokenHtml,
		"html/errorNoSaveToken.html",
	)
}

func htmlErrornosavetokenHtml() (*asset, error) {
	bytes, err := htmlErrornosavetokenHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/errorNoSaveToken.html", size: 1090, mode: os.FileMode(420), modTime: time.Unix(1610538818, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlErrornostscredHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x51\x6f\xd3\x30\x10\xc7\x9f\xd3\x4f\x71\xeb\xc4\x1b\x6e\xda\x49\xc0\xc8\xdc\x48\xd3\x36\x04\x12\xa8\x88\x06\xa1\x3d\xba\xf6\xb5\x36\x73\xed\x60\x5f\xd8\x2a\xb4\xef\x8e\x9c\x34\x6d\x8a\xf6\xc0\x4b\xd3\x5c\xee\xff\xf3\xdd\xff\xce\xfc\xec\x76\x71\x53\xdd\x7f\xbd\x83\x8f\xd5\x97\xcf\xe5\x88\x6b\xda\xda\x72\x94\x71\x8d\x42\xa5\x27\x19\xb2\x58\x2e\xab\x25\xfb\xf1\xe9\xdb\x1d\xcf\xbb\xf7\x51\xc6\xb7\x48\x02\x34\x51\xcd\xf0\x57\x63\x7e\xcf\xc7\x37\xde\x11\x3a\x62\xd5\xae\xc6\x31\xc8\xee\x6d\x3e\x26\x7c\xa2\x3c\x51\xaf\x40\x6a\x11\x22\xd2\xfc\x7b\xf5\x81\x5d\x8e\x21\x4f\x98\x48\xbb\x96\x97\xa5\x94\xd7\xa3\x2c\x5b\x79\xb5\x83\x3f\x29\x80\x66\xa3\xa9\x80\xd9\x74\xfa\xea\x6a\x94\x65\xcf\xc3\x8f\xca\xc4\xda\x8a\x5d\x01\x6b\x8b\x4f\xe9\x6b\x7a\xb2\xc7\x20\xea\x02\xd2\x6f\x0a\x6d\x45\xd8\x18\x57\xc0\xb4\x57\x4f\x52\x53\x18\xd8\x16\x5d\x93\x8e\x5a\x7b\x4f\x18\x5e\xe6\x09\x6b\x36\x8e\x19\xc2\x6d\x2c\x40\xa2\x23\x0c\x29\xfc\x68\x14\xe9\xd3\x9a\x86\xd4\x96\xf5\xb3\x89\x64\xd6\x3b\xb6\xb7\x60\x28\xef\x7b\x7a\x3b\xad\xdb\x53\x56\x42\x3e\x6c\x82\x6f\x9c\x2a\xe0\x7c\x26\x2f\xdf\xc9\xf7\x29\x2c\xbd\xf5\xa1\x80\xf3\xf5\x7a\xdd\x9f\xa2\x2f\x5a\xf6\xa1\x27\x98\xc2\xe5\xbf\x0c\xb6\xd7\x05\x54\xbd\xac\xb1\x60\xcd\x69\x87\xc6\x59\xe3\x90\xad\xac\x97\x0f\x29\xad\x16\x4a\x19\xb7\x49\xcc\xd9\xbe\x2c\x6b\x22\xb1\x76\x32\x05\x38\xef\xb0\xa7\x45\x94\x64\xbc\x6b\x79\xc9\xa8\x02\x66\x03\x53\xde\x74\x9e\x1c\x78\x3d\x2d\x29\x45\x20\x23\x2d\x9e\xf4\x20\x1a\xf2\xff\x21\x4f\x0b\xc4\xda\x71\x0c\x9d\x7c\x3e\x9d\xdf\x0b\x4d\x9c\x78\xab\x54\x6f\x09\xcf\xfb\x95\xe3\x79\xbf\xe4\x69\xad\xfa\xa5\xc7\x00\xd2\x8a\x18\xe7\xe3\xc1\x5c\xc7\x69\x43\xb9\x9e\x0d\x2e\x82\x9e\x1d\x10\x18\xda\x4d\xee\xcc\x69\x33\xf7\xed\x76\xaa\x3e\x23\xfd\xbf\x28\xef\x42\xf0\x81\xe7\xfa\xa2\x8b\xc4\xad\xb0\xb6\x5c\x5c\x37\xa4\xa1\x0e\x5e\x62\x8c\x3c\xef\x82\x49\x7c\xe4\x67\xbc\x2e\x6f\x7c\x63\x15\x38\x4f\xb0\x41\x82\x65\xb5\x04\x19\x50\xa1\x23\x23\x6c\x9c\x4c\x26\x7c\x15\xca\x7b\xdf\x80\x14\x0e\xa4\xf5\x11\x81\xb4\x89\x40\x62\x75\xc6\xf3\xba\x03\x1e\x4b\xe3\xf9\xb1\x64\xde\x79\xd9\xa6\x74\xa7\xdf\x2e\x6e\xaf\x97\xac\x5a\x1e\xab\xe1\xf9\x21\x89\xe7\x9d\x65\xbc\xbd\xd9\xe5\xdf\x00\x00\x00\xff\xff\xd3\xce\x66\x6d\x46\x04\x00\x00")

func htmlErrornostscredHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlErrornostscredHtml,
		"html/errorNoStsCred.html",
	)
}

func htmlErrornostscredHtml() (*asset, error) {
	bytes, err := htmlErrornostscredHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/errorNoStsCred.html", size: 1094, mode: os.FileMode(420), modTime: time.Unix(1610538876, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlErrornotokenHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x51\x6f\xd3\x30\x10\xc7\x9f\xd3\x4f\x71\xeb\xc4\x1b\x6e\xda\x49\xc0\xc8\xdc\x48\xd3\x36\x04\x12\x68\x88\x16\xa1\x3d\xba\xce\x25\x36\x75\xec\x60\x5f\xd8\x2a\xb4\xef\x8e\x9c\x34\x6d\x8a\xf6\xc0\x4b\xd3\x5c\xee\xff\xf3\xdd\xff\xce\xfc\xec\xf6\xfe\x66\xfd\xf0\xf5\x0e\x3e\xae\xbf\x7c\xce\x27\x5c\x51\x6d\xf2\x49\xc2\x15\x8a\x22\x3e\x49\x93\xc1\x7c\xb5\x5e\xb1\x1f\x9f\xbe\xdd\xf1\xb4\x7f\x9f\x24\xbc\x46\x12\xa0\x88\x1a\x86\xbf\x5a\xfd\x7b\x39\xbd\x71\x96\xd0\x12\x5b\xef\x1a\x9c\x82\xec\xdf\x96\x53\xc2\x27\x4a\x23\xf5\x0a\xa4\x12\x3e\x20\x2d\xbf\xaf\x3f\xb0\xcb\x29\xa4\x11\x13\x68\xd7\xf1\x92\x98\xf2\x7a\x92\x24\x1b\x57\xec\xe0\x4f\x0c\xa0\xae\x14\x65\xb0\x98\xcf\x5f\x5d\x4d\x92\xe4\x79\xfc\xb1\xd0\xa1\x31\x62\x97\x41\x69\xf0\x29\x7e\x8d\x4f\xf6\xe8\x45\x93\x41\xfc\x8d\xa1\x5a\xf8\x4a\xdb\x0c\xe6\x83\x7a\x16\x9b\x42\xcf\x6a\xb4\x6d\x3c\xaa\x74\x8e\xd0\xbf\xcc\x13\x46\x57\x96\x69\xc2\x3a\x64\x20\xd1\x12\xfa\x18\x7e\xd4\x05\xa9\xd3\x9a\xc6\xd4\x8e\xf5\xb3\x0d\xa4\xcb\x1d\xdb\x5b\x30\x96\x0f\x3d\xbd\x9d\x37\xdd\x29\x1b\x21\xb7\x95\x77\xad\x2d\x32\x38\x5f\xc8\xcb\x77\xf2\x7d\x0c\x4b\x67\x9c\xcf\xe0\xbc\x2c\xcb\xe1\x14\x75\xd1\xb1\x0f\x3d\xc1\x1c\x2e\xff\x65\xb0\xbd\xce\x63\x31\xc8\x5a\x03\x46\x9f\x76\xa8\xad\xd1\x16\xd9\xc6\x38\xb9\x8d\x69\x8d\x28\x0a\x6d\xab\xc8\x5c\xec\xcb\x32\x3a\x10\xeb\x26\x93\x81\x75\x16\x07\x5a\x40\x49\xda\xd9\x8e\x17\x8d\xca\x60\x31\x32\xe5\x4d\xef\xc9\x81\x37\xd0\xa2\x52\x78\xd2\xd2\xe0\x49\x0f\xa2\x25\xf7\x1f\xf2\xb8\x40\xac\x1b\xc7\xd8\xc9\xe7\xd3\xf9\xbd\xd0\xc4\x89\xb7\x45\x31\x58\xc2\xd3\x61\xe5\x78\x3a\x2c\x79\x5c\xab\x61\xe9\xd1\x83\x34\x22\x84\xe5\x74\x34\xd7\x69\xdc\x50\xae\x16\xa3\x8b\xa0\x16\x07\x04\xfa\x6e\x93\x7b\x73\xba\xcc\x7d\xbb\xbd\x6a\xc8\x88\xff\x2f\xf2\x3b\xef\x9d\xe7\xa9\xba\xe8\x23\xa1\x16\xc6\xe4\xf7\xd7\x2d\x29\x68\xbc\x93\x18\x02\x4f\xfb\x60\x14\x1f\xf9\x09\x6f\xf2\x1b\x61\xad\x23\xa8\x90\x80\xdc\x16\xed\x6c\x36\xe3\x1b\x9f\x3f\xb8\x16\xa4\xb0\x20\x8d\x0b\x08\xa4\x74\x00\x12\x9b\x33\x9e\x36\x3d\xe3\x58\x0d\x4f\x8f\x55\xf2\xde\xbe\x2e\xa5\x3f\xf0\xf6\xfe\xf6\x7a\xc5\xd6\xab\x63\x01\x3c\x3d\x24\xf1\xb4\x77\x89\x77\x97\x39\xff\x1b\x00\x00\xff\xff\xe2\x7c\xdd\x3b\x39\x04\x00\x00")

func htmlErrornotokenHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlErrornotokenHtml,
		"html/errorNoToken.html",
	)
}

func htmlErrornotokenHtml() (*asset, error) {
	bytes, err := htmlErrornotokenHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/errorNoToken.html", size: 1081, mode: os.FileMode(420), modTime: time.Unix(1610538488, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlErrortokenexpiredHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x51\x6f\xd3\x30\x10\xc7\x9f\xd3\x4f\x71\xeb\xc4\x1b\x6e\xda\x49\xc0\xc8\xdc\x48\xd3\x36\x04\x12\x68\x88\x06\xa1\x3d\xba\xce\xa5\x36\x73\xec\x60\x5f\x58\x2b\xb4\xef\x8e\x9c\x34\x6d\x8a\xf6\xc0\x4b\xd3\x5c\xee\xff\x3b\xdf\xff\xce\xfc\xec\xf6\xfe\xa6\x78\xf8\x7a\x07\x1f\x8b\x2f\x9f\xf3\x09\x57\x54\x9b\x7c\x92\x70\x85\xa2\x8c\x4f\xd2\x64\x30\x5f\x15\x2b\xf6\xe3\xd3\xb7\x3b\x9e\xf6\xef\x93\x84\xd7\x48\x02\x14\x51\xc3\xf0\x57\xab\x7f\x2f\xa7\x37\xce\x12\x5a\x62\xc5\xae\xc1\x29\xc8\xfe\x6d\x39\x25\xdc\x52\x1a\xa9\x57\x20\x95\xf0\x01\x69\xf9\xbd\xf8\xc0\x2e\xa7\x90\x46\x4c\xa0\x5d\xc7\x4b\x62\xca\xeb\x49\x92\xac\x5d\xb9\x83\x3f\x31\x80\x7a\xa3\x28\x83\xc5\x7c\xfe\xea\x6a\x92\x24\xcf\xe3\x8f\xa5\x0e\x8d\x11\xbb\x0c\x2a\x83\xdb\xf8\x35\x3e\xd9\x93\x17\x4d\x06\xf1\x37\x86\x6a\xe1\x37\xda\x66\x30\x1f\xd4\xb3\xd8\x14\x7a\x56\xa3\x6d\x63\xa9\xca\x39\x42\xff\x32\x4f\x18\xbd\xb1\x4c\x13\xd6\x21\x03\x89\x96\xd0\xc7\xf0\x93\x2e\x49\x9d\x9e\x69\x4c\xed\x58\x3f\xdb\x40\xba\xda\xb1\xbd\x05\x63\xf9\xd0\xd3\xdb\x79\xd3\x55\x59\x0b\xf9\xb8\xf1\xae\xb5\x65\x06\xe7\x0b\x79\xf9\x4e\xbe\x8f\x61\xe9\x8c\xf3\x19\x9c\x57\x55\x35\x54\x51\x17\x1d\xfb\xd0\x13\xcc\xe1\xf2\x5f\x06\xdb\xeb\x3c\x96\x83\xac\x35\x60\xf4\x69\x87\xda\x1a\x6d\x91\xad\x8d\x93\x8f\x31\xad\x11\x65\xa9\xed\x26\x32\x17\xfb\x63\x19\x1d\x88\x75\x93\xc9\xc0\x3a\x8b\x03\x2d\xa0\x24\xed\x6c\xc7\x8b\x46\x65\xb0\x18\x99\xf2\xa6\xf7\xe4\xc0\x1b\x68\x51\x29\x3c\x69\x69\xf0\xa4\x07\xd1\x92\xfb\x0f\x79\x5c\x20\xd6\x8d\x63\xec\xe4\xf3\xe9\xfc\x5e\x68\xe2\xc4\xdb\xb2\x1c\x2c\xe1\xe9\xb0\x72\x3c\x1d\x96\x3c\xae\xd5\xb0\xf4\xe8\x41\x1a\x11\xc2\x72\x3a\x9a\xeb\x34\x6e\x28\x57\x8b\xd1\x45\x50\x8b\x03\x02\x7d\xb7\xc9\xbd\x39\x5d\xe6\xbe\xdd\x5e\x35\x64\xc4\xff\x17\xf9\x9d\xf7\xce\xf3\x54\x5d\xf4\x91\x50\x0b\x63\xf2\xfb\xeb\x96\x14\x34\xde\x49\x0c\x81\xa7\x7d\x30\x8a\x8f\xfc\x84\x37\x79\xe1\x1e\xd1\x02\x6e\x1b\xed\xb1\x9c\xcd\x66\x7c\xed\xf3\x07\xd7\x82\x14\x16\xa4\x71\x01\x81\x94\x0e\x40\x62\x7d\xc6\xd3\xa6\x07\x1c\x8f\xc2\xd3\xe3\x11\x79\xef\x5d\x97\xd2\x57\xbb\xbd\xbf\xbd\x5e\xb1\x62\x75\xac\xce\xd3\x43\x12\x4f\x7b\x8b\x78\x77\x93\xf3\xbf\x01\x00\x00\xff\xff\xb8\xb3\xd3\xf7\x36\x04\x00\x00")

func htmlErrortokenexpiredHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlErrortokenexpiredHtml,
		"html/errorTokenExpired.html",
	)
}

func htmlErrortokenexpiredHtml() (*asset, error) {
	bytes, err := htmlErrortokenexpiredHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/errorTokenExpired.html", size: 1078, mode: os.FileMode(420), modTime: time.Unix(1610538533, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlMountingpageHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xc1\x6e\xdb\x3a\x10\x3c\xcb\x5f\xb1\x71\xf0\x6e\x4f\x96\x1d\xa0\x6d\xaa\xd0\x02\x8a\x24\x45\x03\x34\x48\x51\xbb\x2d\x72\xa4\xa9\x95\xc4\x86\x22\x55\x72\x55\xc7\x28\xf2\xef\xc5\x4a\x96\x2d\x03\x39\xf4\x62\x59\xcb\x9d\xd9\x9d\xe1\x48\x9c\xdd\x3c\x5c\xaf\x1f\xbf\xdc\xc2\xa7\xf5\xfd\xe7\x6c\x22\x2a\xaa\x4d\x36\x89\x44\x85\x32\xe7\x27\x69\x32\x98\xad\xd6\xab\xf8\xc7\xdd\xd7\x5b\x91\xf4\xef\x93\x48\xd4\x48\x12\x2a\xa2\x26\xc6\x5f\xad\xfe\xbd\x9c\x5e\x3b\x4b\x68\x29\x5e\xef\x1a\x9c\x82\xea\xdf\x96\x53\xc2\x67\x4a\x98\xf5\x0a\x54\x25\x7d\x40\x5a\x7e\x5b\x7f\x8c\x2f\xa7\x90\x30\x4d\xa0\x5d\xc7\x17\x71\xcb\xff\x93\x28\xda\xb8\x7c\x07\x7f\xb8\x80\xba\xac\x28\x85\xc5\x7c\xfe\xdf\xd5\x24\x8a\x5e\xc6\x87\xb9\x0e\x8d\x91\xbb\x14\x0a\x83\xcf\x7c\xca\xcf\x78\xeb\x65\x93\x02\xff\x72\xa9\x96\xbe\xd4\x36\x85\xf9\x80\x9e\xb1\x28\xf4\x71\x8d\xb6\xe5\x51\x85\x73\x84\xfe\x75\x3e\x69\x74\x69\x63\x4d\x58\x87\x14\x14\x5a\x42\xcf\xe5\xad\xce\xa9\x3a\xdd\x69\xcc\xda\x71\xfd\x6c\x03\xe9\x62\x17\xef\x2d\x18\xc3\x07\x4d\x6f\xe7\x4d\x37\x65\x23\xd5\x53\xe9\x5d\x6b\xf3\x14\xce\x17\xea\xf2\x9d\x7a\xcf\x65\xe5\x8c\xf3\x29\x9c\x17\x45\x31\x4c\xa9\x2e\x3a\xee\x83\x26\x98\xc3\x65\xcf\xc1\xa7\xad\x01\xa3\x4f\x85\x68\x6b\xb4\xc5\x78\x63\x9c\x7a\xe2\xb6\x46\xe6\xb9\xb6\x25\x43\x17\xfb\xe9\x46\x07\x8a\xbb\x0b\x48\xc1\x3a\x8b\x03\x5b\x40\x45\xda\xd9\x8e\x8f\xfd\x48\x61\x31\xd2\xfe\xa6\x97\x7e\xe0\x1b\xd8\x18\x29\x3d\x69\x65\xf0\x64\x55\xd9\x92\xfb\x07\x38\xe7\x24\xee\x5c\x1f\x1b\xf6\x72\x7a\x4d\xaf\x88\x38\xb1\x30\xcf\xf3\x3d\x4a\x24\x43\xb2\x44\x32\x64\x99\xd3\x33\x64\x1b\x3d\x28\x23\x43\x58\x4e\x47\xd7\x37\xe5\x20\x8a\x6a\x31\xca\x7b\xb5\x38\x50\xa0\xef\x02\xdb\x9b\xd3\x75\xee\xe5\xf6\xa8\xa1\x83\xff\x5f\x64\x77\xb6\x70\x22\xa9\x2e\xfa\x42\xa8\xa5\x31\xd9\xbd\x6b\x2d\x69\x5b\x42\xe3\x9d\xc2\x10\x44\xd2\xd7\x19\x7e\x9c\x10\x89\x26\xfb\xee\x4c\x5b\x23\x6c\xb5\x31\xb0\x41\xa8\x19\x88\x39\x68\x0b\x05\x6e\x21\xa0\x72\x36\x0f\xb3\xd9\x4c\x6c\x7c\xf6\xe8\x5a\x50\xd2\x82\x32\x2e\x20\x50\xa5\x03\x90\xdc\x9c\x89\xa4\xe9\x89\x8f\x4b\x8a\xe4\xb8\xbc\xe8\x5d\xed\x5a\xfa\x2d\x6e\x1e\x6e\x3e\xac\xe2\xf5\xea\xb8\x95\x48\x0e\x4d\x22\xe9\xcd\x13\xdd\xa7\x9c\xfd\x0d\x00\x00\xff\xff\x36\xd8\xb9\x43\x37\x04\x00\x00")

func htmlMountingpageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlMountingpageHtml,
		"html/mountingPage.html",
	)
}

func htmlMountingpageHtml() (*asset, error) {
	bytes, err := htmlMountingpageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/mountingPage.html", size: 1079, mode: os.FileMode(420), modTime: time.Unix(1609243665, 0)}
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
	"html/errorNoCred.html":       htmlErrornocredHtml,
	"html/errorNoSaveToken.html":  htmlErrornosavetokenHtml,
	"html/errorNoStsCred.html":    htmlErrornostscredHtml,
	"html/errorNoToken.html":      htmlErrornotokenHtml,
	"html/errorTokenExpired.html": htmlErrortokenexpiredHtml,
	"html/mountingPage.html":      htmlMountingpageHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"errorNoCred.html":       &bintree{htmlErrornocredHtml, map[string]*bintree{}},
		"errorNoSaveToken.html":  &bintree{htmlErrornosavetokenHtml, map[string]*bintree{}},
		"errorNoStsCred.html":    &bintree{htmlErrornostscredHtml, map[string]*bintree{}},
		"errorNoToken.html":      &bintree{htmlErrornotokenHtml, map[string]*bintree{}},
		"errorTokenExpired.html": &bintree{htmlErrortokenexpiredHtml, map[string]*bintree{}},
		"mountingPage.html":      &bintree{htmlMountingpageHtml, map[string]*bintree{}},
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
