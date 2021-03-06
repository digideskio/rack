// Code generated by go-bindata.
// sources:
// templates/init/rails/.dockerignore
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/.dockerignore
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/.dockerignore
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/init/unknown/.dockerignore
// templates/init/unknown/Dockerfile
// templates/init/unknown/docker-compose.yml
// templates/templates.go
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _initRailsDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x97\x7e\x4a\x92\xbe\x96\x5e\x71\x61\x4e\x66\x49\xaa\x31\x2a\x4f\x37\x2b\xbf\xb4\x28\x2f\x31\x87\x4b\x3f\x27\x3f\x5d\x5f\x8b\x4b\xbf\x24\xb7\x80\x0b\x10\x00\x00\xff\xff\xa0\x04\x95\x56\x4e\x00\x00\x00")

func initRailsDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerignore,
		"init/rails/.dockerignore",
	)
}

func initRailsDockerignore() (*asset, error) {
	bytes, err := initRailsDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 78, mode: os.FileMode(420), modTime: time.Unix(1464015261, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initRailsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xbb\x4e\xc4\x30\x10\xec\xfd\x15\x2b\xd1\x5f\x7a\x5a\x24\xa8\xe0\x50\x24\x0a\xba\xf3\x39\xeb\x60\xe2\xd8\x96\x1f\x27\xf2\xf7\x6c\xd6\x09\x21\xa7\xa4\xca\x3c\xec\x99\xf1\x73\x7b\x7e\x05\xe5\xdd\xcd\xff\x34\x51\x1a\x9b\x84\x78\x20\x1c\x26\xf0\xce\x4e\x90\xbf\x10\xb4\xb1\x98\xc0\x21\x76\xd8\x81\xf6\x11\xae\xc5\x75\x16\xc1\xb8\x94\xa5\xb5\xe4\x2f\x4e\xf9\x71\x44\x97\xd9\x7f\x43\xd7\xf9\xd8\x28\xa9\x08\x58\xe3\xc8\xa9\x61\xf2\x05\x2e\xcb\xc1\x20\xd5\x20\x7b\xbc\xcc\x64\x84\x1e\xc7\x24\x9e\xce\xef\x9f\xf0\x82\xe3\x9c\x05\xfc\x35\x32\x84\x66\x61\x76\xf2\xc9\x7a\x35\xec\x64\x66\xa8\x06\xbb\x76\xe9\xec\xfa\xcf\x88\xf6\xe3\xed\xbe\xff\x3a\xf8\xbb\xa4\x7c\x3c\x58\xa6\x84\x39\x3d\x86\x88\xb4\x33\xfc\x15\x6a\xe5\x80\x4b\x61\x0e\x5a\x71\x55\xe9\x51\xb5\xe9\xb7\x2d\x15\x57\x2d\x94\xab\x35\x6a\xd3\x2a\xae\xda\x8c\x6b\x60\xd5\x36\xcc\xe5\x23\x85\x1c\x14\x5a\x47\xcc\xfd\x23\xd2\x10\xaf\xf9\x9f\x4e\xd7\x6b\x4f\x7c\x9b\xf8\x0d\x00\x00\xff\xff\x08\xae\x24\x4e\xf0\x01\x00\x00")

func initRailsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerfile,
		"init/rails/Dockerfile",
	)
}

func initRailsDockerfile() (*asset, error) {
	bytes, err := initRailsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1463247575, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initRailsDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initRailsDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerComposeYml,
		"init/rails/docker-compose.yml",
	)
}

func initRailsDockerComposeYml() (*asset, error) {
	bytes, err := initRailsDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463517114, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initRubyDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initRubyDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerignore,
		"init/ruby/.dockerignore",
	)
}

func initRubyDockerignore() (*asset, error) {
	bytes, err := initRubyDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464015261, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initRubyDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xcd\x4e\xc0\x20\x10\x84\xef\x3c\xc5\x26\xde\xcb\x43\x98\xe8\x49\x6b\x9a\x78\xf0\x56\x0a\x4b\x6d\x0a\xbb\x84\x9f\x46\xde\x5e\x8a\x35\xb1\x72\x62\x67\xbf\x61\x86\xa7\x69\x7c\x01\xcd\x74\xf0\x97\x8c\x65\xa9\x42\x3c\xb4\x31\x54\x60\x72\x15\xf2\x27\x82\xdd\x1c\x26\x20\x44\x83\x06\x2c\x47\x58\x0a\x19\x87\xb0\x51\xca\xca\xb9\xc6\x17\xd2\xec\x3d\x52\xee\xfc\x81\x64\x38\x4a\xad\x74\x1b\xdc\x46\x8d\xb4\x50\xb9\xc0\x7c\x19\x83\xd2\xbb\x5a\x71\x3e\xc5\x08\x2b\xfa\x24\x1e\xc7\xb7\x0f\x78\x46\x7f\x66\x41\x3f\x52\x85\x20\x2f\xe5\xb6\x1e\x1c\xeb\xfd\xb6\xee\x4a\xab\xd1\xa9\x5b\x7a\xa7\xfe\x2a\x62\x7a\x7f\xfd\xdf\xff\xf7\xc3\x67\xf7\x88\x29\x03\xdb\x7e\x6f\xde\x9f\xe0\xa1\xbf\x23\xbe\x03\x00\x00\xff\xff\x8b\xae\xa0\xae\x2a\x01\x00\x00")

func initRubyDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerfile,
		"init/ruby/Dockerfile",
	)
}

func initRubyDockerfile() (*asset, error) {
	bytes, err := initRubyDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1463247575, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initRubyDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xf0\x6e\xe9\x98\x87\x21\xf8\x30\x5a\x03\xca\xa2\x29\x69\x3a\xed\xdb\x2f\x85\x6d\xb7\xdd\xfa\xf5\xff\xf2\x9d\x38\x0f\x0d\xc0\x9c\x37\x5a\x06\x70\xf6\x0c\xbc\xef\xd3\x61\x80\x61\x65\x68\x71\xd9\x14\x16\x0e\x4f\x94\xce\xa6\xc8\x09\x5d\xd9\x09\xce\x4d\x57\x28\x9c\x05\x92\x4e\xa2\x39\x7e\x0f\x5b\x6b\xd0\x34\x23\xa5\x1a\x06\xe8\x6c\x38\x5e\x7c\xb9\xc8\xa2\xae\xef\xef\x2e\x0a\x2b\x07\xa6\x51\x29\xfd\x57\xae\x32\xaa\x64\x34\xa1\xfe\xfe\x62\x0f\x3f\xf4\xde\xfb\x0f\x99\x5b\xf1\xd6\xbc\x03\x00\x00\xff\xff\xae\x01\x4e\xf5\xc8\x00\x00\x00")

func initRubyDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerComposeYml,
		"init/ruby/docker-compose.yml",
	)
}

func initRubyDockerComposeYml() (*asset, error) {
	bytes, err := initRubyDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1463247575, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initSinatraDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initSinatraDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerignore,
		"init/sinatra/.dockerignore",
	)
}

func initSinatraDockerignore() (*asset, error) {
	bytes, err := initSinatraDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464015261, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initSinatraDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xc1\x6e\x84\x20\x10\x86\xef\x3c\xc5\x24\xbd\xcb\x43\x34\x69\x4f\xad\x8d\x49\x0f\xbd\x49\x61\x70\x89\x30\x43\x00\xcd\xfa\xf6\x8b\xac\x9b\xac\xeb\x49\xbe\xf9\x86\xff\xe7\x63\xe8\xbf\x40\x33\xad\x7c\x95\xd9\x91\x2a\x49\x09\xf1\x56\x49\xdc\x80\xc9\x6f\x50\x2e\x08\xd6\x79\xcc\x40\x88\x06\x0d\x58\x4e\xf0\xbf\x90\xf1\x08\x8e\x72\x51\xde\x57\x7f\x21\xcd\x21\x20\x95\xe6\xaf\x48\x86\x93\xd4\x4a\xd7\x83\x77\x54\x4d\x0b\x1b\x2f\x30\x1e\x8b\x51\xe9\x59\x4d\x38\xee\x30\xc1\x84\x21\x8b\xf7\xfe\xe7\x0f\x3e\x31\xec\x59\xd0\x3e\xa9\x62\x94\x07\x39\x8d\x3b\xcf\x7a\x3e\x8d\x1b\xa9\x35\x9a\x75\x4a\x6f\xd6\x33\x11\xc3\xef\xf7\x6b\xff\xc7\x83\xf7\xee\x09\x73\x01\xb6\xed\xbf\xee\xde\x83\xbb\x76\x8f\xb8\x05\x00\x00\xff\xff\x9c\x51\x49\xbe\x2d\x01\x00\x00")

func initSinatraDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerfile,
		"init/sinatra/Dockerfile",
	)
}

func initSinatraDockerfile() (*asset, error) {
	bytes, err := initSinatraDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1463247575, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initSinatraDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initSinatraDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerComposeYml,
		"init/sinatra/docker-compose.yml",
	)
}

func initSinatraDockerComposeYml() (*asset, error) {
	bytes, err := initSinatraDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463247575, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initUnknownDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\xe1\x02\x04\x00\x00\xff\xff\x9c\x10\x28\x7b\x0a\x00\x00\x00")

func initUnknownDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerignore,
		"init/unknown/.dockerignore",
	)
}

func initUnknownDockerignore() (*asset, error) {
	bytes, err := initUnknownDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1464200391, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initUnknownDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4d\x2a\xcd\x2b\x29\xb5\x32\x34\xd3\x33\x30\xe1\xe2\x72\xf6\x0f\x88\x54\xd0\x53\xd0\x4f\x2c\x28\xe0\x02\x04\x00\x00\xff\xff\xf1\xa3\x65\xfc\x1f\x00\x00\x00")

func initUnknownDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerfile,
		"init/unknown/Dockerfile",
	)
}

func initUnknownDockerfile() (*asset, error) {
	bytes, err := initUnknownDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 31, mode: os.FileMode(420), modTime: time.Unix(1464200309, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _initUnknownDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\xcc\xcc\xb3\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\xa6\xe1\xc1\x85\x11\x00\x00\x00")

func initUnknownDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerComposeYml,
		"init/unknown/docker-compose.yml",
	)
}

func initUnknownDockerComposeYml() (*asset, error) {
	bytes, err := initUnknownDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 17, mode: os.FileMode(420), modTime: time.Unix(1464200354, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x4b\x6f\x23\xc7\x11\xc7\xcf\xe2\xa7\x18\x0b\xb0\x41\x05\x32\x35\xef\x87\x00\x5f\xbc\xeb\x00\x7b\xc8\x1a\xb0\xd7\x87\x20\x0a\x8c\x79\xf4\x28\xc4\x4a\xa4\x42\x52\xb6\xb4\x8b\xfd\xee\xa9\x5f\x57\x8f\x48\x91\x43\xea\x45\x21\xce\xe3\x30\xe2\x4c\x4f\x77\x75\x55\x77\xd5\xbf\xff\x55\xa3\x93\x13\xef\xcd\xb4\x31\xde\xb9\x99\x98\x59\xb9\x30\x8d\x57\xdd\x7a\xe7\xd3\x6f\xab\xf1\xa4\x29\x17\xe5\x68\x20\x1d\xe6\xd3\xeb\x59\x6d\xe6\xa7\xdc\x2f\xcc\xe5\xd5\x85\xf4\x9b\x9f\x8c\x27\xe3\xc5\xc9\xac\x1c\x5f\xcc\x4f\x46\xcd\xb4\xfe\x68\x66\xe3\xf3\xc9\x74\x66\xb6\xf6\x7a\x6b\x3b\xb5\xe3\x8b\xed\x5d\x54\xce\xb7\xf5\xf4\xf2\x6a\x3a\x37\xa3\xdb\xcb\x8b\xbe\xae\xd7\xd5\xed\xc3\x53\xd2\x69\xf7\x8c\xf4\x78\xd4\x84\xf3\xf1\xa4\x5c\xcc\xca\x07\xe7\xec\xfa\xed\x9c\xb6\xeb\xf4\xa8\x99\xaf\x27\x1f\x27\xd3\xdf\x27\x0f\xce\xdc\xf5\xdb\x39\x73\xd7\xe9\xa1\x99\xef\xee\x46\xe7\x53\xde\xbc\xfd\xd1\x7b\xff\xe3\x07\xef\x87\xb7\xef\x3e\x7c\x35\x18\x5c\x95\xf5\xc7\xf2\xdc\x2c\xfb\x0f\x06\x63\x11\x34\x5b\x78\xc3\xc1\xc1\x61\x75\x2b\x2d\x87\x72\x83\xf4\x99\x99\xcf\x4f\xce\x3f\x8d\xaf\x68\x68\x2f\x17\xfc\x8c\xa7\xfc\x9d\x2f\x66\xe3\xc9\xb9\xed\x38\xb5\x7f\x17\xe3\x4b\xa3\xaf\x4f\xc6\xd3\xeb\xc5\xf8\x82\x87\xab\x72\xf1\x8f\x13\x8c\xe1\xe6\x70\x70\x34\x18\xb4\xd7\x93\xda\x73\xae\xf9\x93\x29\x9b\x21\x37\xde\xdf\xfe\xce\xb4\xc7\xde\xa4\xbc\x34\x9e\x8a\x3e\xf2\x86\x5d\xab\x99\xcd\xa6\xb3\x23\xef\xf3\xe0\xe0\xfc\x93\x7d\xf2\x4e\xbf\xf3\xd0\x6a\xf4\xde\xfc\x8e\x10\x33\x1b\x5a\xb5\x79\xfe\xfe\xba\x6d\xe5\x19\xb1\x47\x47\x83\x83\x71\x6b\x07\x7c\xf5\x9d\x37\x19\x5f\x20\xe2\x60\x66\x16\xd7\xb3\x09\x8f\xc7\x9e\x98\x34\xfa\x01\xe9\xed\xf0\x10\x41\xde\xd7\xff\x3c\xf5\xbe\xfe\xed\x50\x35\xb1\x73\x89\x8c\x2f\x83\xc1\xc1\x6f\xe5\xcc\xab\xae\x5b\x4f\xe7\xd1\x49\x06\x07\xbf\xaa\x3a\xdf\x79\xe3\xe9\xe8\xcd\xf4\xea\x76\xf8\x8d\xf4\x39\x16\xdd\x64\x54\x7d\xf1\x43\xa7\xe9\xe8\xcd\x85\xec\xd3\x50\xcc\xdf\x93\x3e\x88\x51\xf9\x5b\x04\x49\x47\xd5\xdb\x35\x8a\x5a\xa3\xef\x51\x7d\x78\x74\x4c\x8f\x81\xbc\x5b\xdc\x5e\x19\xaf\x9c\xcf\xcd\x82\x25\xbf\xae\x17\x48\xb1\xf6\xb9\xfd\x90\x69\x26\xed\xd4\xf3\xa6\xf3\xd1\x9f\x65\x0f\xdf\xc9\xc3\xdd\x38\xb7\x85\x5d\xfb\x8a\x84\x95\x3d\x1c\x1c\xcc\xc7\x9f\x8c\x37\x9e\x2c\xd2\x78\x70\x70\x09\x4a\x39\x59\x7f\x91\x7b\xdb\xf2\x41\xdc\xc6\xc3\x77\x46\xdc\x21\xde\x7a\xc8\xb0\x1d\xaf\x4f\x71\xe4\xbd\x17\xc9\xc3\x23\x27\x9b\xa9\x9c\x71\xed\x78\xc4\xa4\x32\x78\xfb\xd8\x9f\x45\x11\x19\x6b\x55\xb9\x3f\x14\x15\x77\x0e\x45\x57\x19\xba\xa2\xf9\x7d\x01\xd8\xf5\x90\x00\x8c\x13\x19\x77\x86\x6e\x48\x70\xd6\x6f\x17\xf2\x6e\xfe\x76\x3c\x13\x11\xd5\x74\x7a\xb1\x3a\xba\xbc\x98\x3f\x60\xf9\xed\x5c\x0d\x17\x58\x29\x6b\xf3\xf9\xcb\xca\x68\xe7\x09\x38\xf7\xaf\x20\xcc\x4f\x60\xf8\xdb\x15\xa4\x12\xd7\x56\x5f\x18\x1e\x9e\xdd\x04\xed\xd9\x4d\x5e\x9d\xdd\xf8\xb9\x5c\xbe\xbb\x8a\xb3\x9b\xd4\x48\xbb\x6b\x6b\xa5\x4f\x13\xca\x95\x9d\xdd\xc4\xd2\x37\x2c\xcf\x6e\xea\x46\xef\x6b\xe9\x1b\xcb\x65\x92\xfb\x7d\x6a\x19\x5f\xcb\xb8\x90\x7b\xb9\xca\x56\x65\x45\xd2\x27\x91\xab\x8d\xa4\x5d\xe4\xe4\xd2\x96\xc6\x67\x37\x99\xdc\xa7\xa9\xce\x5d\x88\x8c\x4c\xc6\xc7\xd2\x56\x48\xdf\x4a\xee\x0b\x79\x97\xc8\x6f\x16\x48\x3f\xb9\x62\xa3\xfd\x99\xbb\x94\x7e\x51\xa0\x7a\xc5\x32\x4f\x94\xe9\xbc\x95\xdc\x57\x22\x3b\x14\x3b\xc2\x56\xfb\xe4\x4e\xbf\x08\xdd\x32\xfd\x4d\xc4\x96\xc4\xad\x43\xec\xc6\x85\x32\xae\xca\x54\x3f\x5f\xda\x02\x7f\xb9\x3e\xac\x07\x57\xc9\xb3\xf4\x2b\xc4\xf6\x24\x55\x9d\xee\xd6\xd0\x3f\xec\x90\xb1\x77\x13\x5c\xe0\xf6\x01\x62\x17\xde\x2b\x80\x2a\x48\xd0\xbf\x97\xc7\xf2\xe6\x70\xdb\x99\x7f\x28\x6f\x8f\xee\xc2\xaf\x77\x3c\x1a\xfc\xc9\xe2\xc5\xaa\x06\x16\x30\xee\x50\x79\x97\xfe\x0f\x81\xdf\x1d\x66\x59\xd4\x11\x61\x6b\xae\xfc\x99\x20\x3f\xf5\x76\x98\xe0\x11\xcb\xa7\x5e\x96\x1f\x7b\x04\xe5\xe9\x6a\xcc\x0e\xe3\xd0\x3f\xb2\xed\x84\xda\xa9\x86\xe2\x2f\x93\xf1\xcd\x30\x88\xd3\xd8\x0f\x92\x30\x0d\x8e\x3d\xff\x48\xc0\xb5\x64\xf2\x6f\xac\xa5\x9f\xad\x79\xa7\x9e\xb3\x12\xcd\x4e\x3d\xfb\xf3\xe5\x6e\xf1\xcb\xe3\x5d\x71\xc4\xe9\xf7\xac\x28\x4a\x6b\xf1\x14\xb9\xaf\x2a\xf5\x96\x5a\xbc\x27\xf2\xd5\xbb\x8c\xbc\x6b\xc5\x13\x83\x44\xbd\xb7\x09\xd4\x2b\x89\x8c\xa4\x54\x8f\x2c\x45\x96\xf1\x55\x06\xcf\xbe\xb4\x57\xa5\x46\x53\x44\x14\xca\xb8\x14\x59\x44\x62\xae\x51\x13\x38\x4f\x6f\x89\xba\x4c\x75\x68\x5c\xa4\x85\x32\x47\x29\x6d\x65\xac\xd1\x18\xd5\xaa\x47\x21\xef\x5a\x79\x97\x89\xdc\xac\xd2\x68\xf4\x13\x17\xe5\x8d\x46\x3f\xf6\x44\x32\x2e\x91\x7e\x01\x91\x2a\xfd\x72\x22\x99\x68\xc3\x26\x91\x13\xca\x3c\xad\xaf\x68\x80\xbd\x85\xaf\xd1\x85\xbd\x44\x66\x25\x63\x0b\xe9\x5f\xfb\xaa\x4b\x96\xaa\xde\xb9\xdc\xb7\xe8\x8e\x7e\xac\x93\xcc\x9b\xc9\x15\x48\x5b\x2d\x6d\x15\xb6\xb1\x1e\xd2\x56\xa1\x57\xac\x11\xce\x1c\x6d\xa1\x28\x11\xc7\x8a\x24\x0d\xcf\xbc\x8b\x14\x9d\x78\xcf\x1c\xa0\x42\x59\xe8\x9e\x25\xac\x2b\x88\xe4\xda\x40\x18\xd6\x0b\x9b\x43\xa3\x08\xc6\xbc\xd8\xd5\x24\xfa\xcb\xba\x54\x62\x63\x5d\x2b\x82\x61\x7b\x40\xdf\x4c\xf7\xa6\x30\x6a\xb7\x91\xb9\x8b\x56\xd7\x21\x0f\x75\x9e\xa2\x56\xd9\xad\xfc\x46\x89\xa2\x22\xe3\x41\xb4\xd4\xad\x03\xf3\x83\xae\xec\x3f\x7d\x8c\x9b\x87\x3e\xf8\x81\x9f\xaa\x1f\xd1\xd7\x14\xce\x7f\x4a\x45\x5d\x7c\x8e\xf5\x63\x2e\xd3\x28\x82\xda\xfd\xc2\x57\x72\x1d\xc7\x9e\x67\xb5\xee\x39\xf6\x97\xa9\xea\x80\x1f\x45\x32\x26\x4f\x55\x0e\xfb\x14\x45\xaa\x2b\xbe\x99\x96\xea\x07\x20\x23\x68\x89\xbe\xbe\x51\x1f\x65\xdd\x93\x44\xf5\xc1\x2f\x2a\x77\x0f\x92\x86\xce\xb7\x13\x77\x6f\xf7\xaf\x52\xdb\x90\xd9\x44\x8a\xe0\xec\x77\xc3\xc9\x80\xbd\xce\x9f\x59\xf3\x28\xd5\xb5\x66\xee\x50\xfa\xe6\x89\xfa\x58\x10\xeb\xdc\x71\xa6\xef\x89\x21\xde\x73\x1a\xe1\x67\x9c\x3c\xec\x79\xe1\xfc\x03\x1f\x40\x2e\x27\x05\xf6\xe3\xb7\xac\x89\xdf\x6c\x22\x3c\xbe\x81\x3e\xec\xa7\xf5\x2d\xde\x07\x0f\x21\x3c\xf0\xf0\x62\x7c\x47\xc8\x3a\xba\x2f\xdf\xec\x84\x76\x3a\x3c\x03\xd8\x57\xd4\x7e\x05\x58\x5f\xd5\xdd\x61\x7a\x5c\xa4\x4f\x05\xf5\x28\x8c\xb3\x24\x4b\x5e\x03\xd4\xdf\x68\xe6\xf5\xd7\xcb\x8b\x67\x41\x7b\x47\x3e\x62\xc2\x2d\x54\x08\x86\xe8\x00\xd5\x71\xbe\x24\x4d\xb8\x28\xa4\x05\xe8\xe3\x7d\x83\xbc\x48\x21\x00\xb8\x0d\xe4\xb7\x88\xf4\x1d\xcf\x39\x72\x03\x85\x0e\xe0\xbe\x83\x7c\x7e\x81\x07\x0b\x77\xe8\x93\xe8\x3d\x21\x06\x64\xf8\xee\x39\x26\x5c\x5c\x1b\x70\xcc\x55\xa7\xcb\x3e\xb1\xeb\x17\xba\xdf\x4e\x26\x30\x50\xa5\xda\xce\x3d\x10\x6c\x61\x97\x10\x8e\xf5\x22\x4c\xed\x11\xe3\x48\x10\x10\xe2\x87\xaa\x83\x85\x07\x69\x37\x81\x42\x49\x82\x6e\x0e\x6a\x0b\x77\xdf\x5d\xb1\xaf\x36\xf0\x5b\x39\x78\xb5\x61\xc7\xd1\x92\xe9\xf3\x7a\x68\x72\x64\x20\x3b\x0c\x15\x62\x81\xae\x87\xc9\xd7\x72\x93\x5f\x1c\xa0\x4b\x51\xeb\x61\xba\x99\xca\xef\x0c\xd7\xa5\xa0\x67\x04\xed\x86\x41\xaf\x10\xba\x7d\xf6\xb8\x10\x0e\xa2\xf0\xc9\x21\x9c\x04\x59\x10\xc4\xfb\x0b\xe1\xeb\xea\xf6\x3f\x2a\xbd\xd9\xea\xd1\xe8\x95\x6b\xca\x93\xf9\x4a\x26\xb6\x79\xf4\x9a\xcd\xcf\x74\xe6\x35\x29\x2b\x7e\xbc\x51\xcc\xeb\xf1\xe0\xb5\xd1\x8f\x76\xde\x7e\xdd\xf7\xeb\xb7\x3d\xfa\x3b\x8f\x8d\xfc\x7f\x7b\x22\x71\xb7\x00\xcf\xce\x23\x12\x39\x44\xf2\xd6\xb9\xa9\x71\xfc\xd9\xe5\x11\x00\x21\x80\x08\xaf\x03\xb4\x01\xfc\xc6\x71\x58\x38\x30\x7c\xdb\xe4\xea\xd6\xf0\xa0\xa2\x54\xce\x0e\xb7\x21\x8b\xe5\x60\xb1\xbc\xae\x71\x87\x8d\xe3\xf5\x96\x3f\x39\x59\x64\xe2\xb9\xe3\xa5\x1c\x64\x99\xb8\x6c\xca\x95\x29\xbf\x4f\x1d\xe8\x93\x5b\xa4\x85\xf2\x2c\x78\x13\xfa\x66\xb1\xce\x05\x37\xb7\x1c\x2f\x51\xce\x1d\x87\xaa\xb3\xe5\x9b\x81\x1e\x30\x70\xbf\xcc\x1d\x28\xf0\x61\xf2\x18\x0e\x42\x72\x90\x2e\xbf\x08\xdd\x41\x93\x47\xca\x89\xc3\x5a\xb9\x20\x07\x0d\xfa\x07\x85\x66\xfd\x95\xaf\x9c\x94\xb0\xe4\xaa\x0a\x3d\x08\x09\x47\x42\x1b\xfe\x0c\xef\x25\x7c\x8d\x51\xbe\x4b\x1e\x42\xf8\x46\xb9\x56\x23\x4a\xa7\x7f\x40\x7e\x56\xeb\x7a\xe4\x8d\xea\x6c\xf3\xba\x42\xf7\x02\x7b\xed\xbc\x8e\x87\xb2\x86\x1d\x3f\x8e\x5c\x3e\xc3\x1e\x03\x1f\xad\xcb\xed\xba\xdc\x01\xae\xec\xb7\xba\x27\xc8\x27\x77\xb1\x95\x8e\xc0\x55\x2b\x42\xe5\xbe\xd8\x4d\xce\x03\x34\x71\x50\x33\x9e\x35\x21\x07\x84\x03\xd3\x66\x5c\x15\xa5\x74\xf9\x24\xeb\x9c\x54\x0a\x47\xec\x4d\x6b\x14\xee\xd8\x3b\xda\x98\x8b\xb1\x8d\x83\x24\xf2\xc5\x3a\xd2\x7d\xe5\x1e\x1f\x0c\x0b\x25\x2c\x8d\xcb\x0b\xd3\x56\x7d\x02\xff\x20\x37\x2d\x03\x57\x41\x89\x34\xef\xa0\xef\x3a\xd4\xe1\xe3\x1c\xf0\x65\x77\xd0\x97\xdb\x79\xf5\xbd\x68\x79\x29\xd0\xad\xb1\xea\xfb\x9f\x23\x76\x61\xdc\x93\x38\x75\x9f\xca\xfb\xc7\xb7\x1e\x46\x1d\x16\x4f\x2e\x93\xec\x9d\x51\xdf\x19\xff\x42\x42\x4d\x9c\x81\x0f\xe4\xdd\xc4\x31\xb1\x44\xad\xc4\xaf\x15\x5b\x2c\xc6\x19\x47\x60\x7d\x1d\x4f\xae\x09\xbe\x11\xef\x90\x47\xf2\x3a\xc6\x10\x7b\xbe\x3b\x92\xcb\x50\x7d\x18\x5c\x22\x1f\x26\x3f\xc6\x97\x89\x25\xe2\x0a\xdc\x23\x06\xc1\x18\x62\xb3\x4d\x9c\xdf\x42\xcc\x1b\xc5\x02\xe2\x93\x7c\x91\x58\x27\x97\xa7\xe2\x68\xe7\x48\xf5\xe8\x26\xf7\x46\x4f\xea\x03\xe8\x49\x6e\x6b\x31\xd1\xd7\x5f\x30\x8f\xa3\x1e\x3c\xa0\x3e\x01\x81\xa6\xb6\x41\x1e\x4d\x12\x41\x6e\x0c\xc5\x80\x50\x53\x93\x60\x7d\xa8\x52\xd2\x97\x7b\x5b\x4f\x69\x5c\x2e\x9d\x2f\xeb\x09\x50\x07\xce\x03\x6c\x84\xb4\xdb\xba\x4d\xab\x31\x0f\xd6\x93\x6c\x50\xb3\x21\x36\xa9\x9d\x04\x4e\x67\xce\x03\xea\x42\xd8\x66\xab\xac\xb5\xc3\xb2\x5c\x65\x95\xae\x8e\xc1\x45\x3d\x84\xfd\x00\xbb\xfc\x4c\x75\x04\x57\x59\xd3\xd6\xe9\x44\x7f\xa8\x10\x95\x59\x8b\x9f\xb1\xe2\x02\x78\xc3\x05\xde\xa0\x17\x78\xd6\xc6\x8a\x1f\x6d\xa5\x6d\x85\xc3\x28\xea\x4f\xd4\xa9\x58\xc3\x3e\x0c\x61\x0e\xdf\x55\x82\xd9\xa3\x3a\x7f\x04\x5d\x7a\x31\xff\xef\x91\xb4\x86\x27\x8f\x62\xff\x3d\x62\x9e\x8e\x2e\xaf\xcc\xfd\xb7\x19\xd3\x61\x8d\xff\x64\x26\xb5\x6f\xac\xf9\x59\xbf\xda\xfe\xaf\xb1\xff\x1e\xb3\x9f\xe7\xcd\x3d\x82\x96\xce\xdc\xfb\x7d\x7d\xd3\x95\x7b\x64\x3c\xd6\x93\xb7\xdb\xb1\x57\x47\xde\x62\xc8\x1f\x26\x1f\xb8\xb7\x0c\x2f\x4f\x09\x02\xf7\x2e\x5e\x49\x09\xd2\xb5\x94\x80\xd2\x7b\xb3\x4c\x09\x80\x67\x8e\x43\x8e\x1e\x28\x2e\x94\x16\xbf\x07\x8e\x2b\x77\x6f\xcb\xd2\xc8\x2a\xf4\x48\x8d\x1d\x24\xd7\xdd\x31\x99\x3a\x5a\xe7\x4a\xaa\xa5\xa3\xa2\xb6\xd6\xe5\xca\xf2\xe8\x81\xae\x86\xb4\x20\xd2\x63\x07\xaa\x48\xdd\x87\x12\xb3\xa5\xd7\xb5\xd6\x92\x38\xea\x8a\xee\xa3\x9e\x3b\xf6\x38\x12\x12\xf7\xd1\x0f\xfa\x4d\x7c\xd9\xb8\x69\xdd\x27\x11\xa3\xeb\xc4\xb1\x41\x3d\xc9\xd6\xda\x7c\xb7\x4e\xa9\x1e\xd9\xa4\x05\x5d\x29\x97\xb9\xf9\x24\x81\x5c\xc6\xe5\x46\xd3\x0d\x8e\x33\x9b\xe6\x64\x4a\x89\x43\x97\xba\x84\xee\x13\x00\xfa\x99\x54\xe3\x9e\xda\x1e\x76\x51\xb3\x6b\x62\xa5\xdb\xe0\x04\x74\x25\xcb\x34\x8d\xe1\xa8\xb6\x1f\x11\x1d\xcd\x88\x5c\x7c\xb3\x5e\xcc\x65\xcb\xe5\x85\xf6\xe3\x08\x86\xb6\xd3\x87\xbd\xb4\xfb\x18\x68\xdf\xd6\x51\x76\xe3\x52\x04\xe8\x0f\xa9\x12\xa9\x0b\xe3\xed\xf1\x6e\xdc\x27\x8c\x50\x69\x04\x47\x34\x54\x80\x3d\x8d\x2a\xc5\x2d\xfb\x49\xc1\x68\x1b\x3a\xd1\x9f\x54\xce\xa6\x82\x89\xda\x60\x69\x7d\xaa\xfb\xc2\x67\x18\xd2\x00\xe3\x52\x02\xa8\x84\xc5\x2c\x67\x27\x6d\xe8\xcc\x1e\x51\xb3\xe4\x58\xb6\xe9\x47\xa3\x6b\xcc\x3b\xd6\x93\xbe\xdd\x27\x0b\x7c\x95\x7a\x22\xb4\x65\x1d\x0b\xa1\x33\xac\x4b\xe7\x3b\x7c\x2e\xd8\x92\x1e\x6c\x04\xcf\x1e\x90\xf0\x7e\x92\xb0\xf9\xff\x43\x0f\x80\xe0\x53\x52\x85\x6d\xea\xbf\x0a\x00\xf6\x24\x0c\x91\x1f\xfc\xa1\x0e\xf1\xff\x17\xe1\xff\x5b\x8b\xf0\x5b\xb6\x79\x0f\xe1\xda\xc7\xc4\xb7\xff\x3f\xdf\x03\xc1\xfb\x74\x3e\xbe\xdb\xb0\x57\x09\xe4\x3d\x17\xe4\xf7\x1c\xd0\xbf\xe8\x7f\x34\xee\x85\x95\x5b\xb6\xdd\x28\xd3\x36\xd1\x4a\x5b\xab\x85\x3e\x9c\x0f\x67\xf7\xe3\xfe\x53\x04\xf6\x43\x82\xcc\xc7\x7a\x1b\xb8\xfd\xce\xd9\xa3\xf2\xf3\x1c\xb3\x47\xd0\xd2\x29\x7b\xff\x6f\x74\xd3\x1f\x7b\x64\x3c\xd6\x17\xb7\xdb\xb1\x57\x3f\xdc\x62\x48\xe7\x82\x4f\x67\xd4\x92\x4b\x46\xc5\xfe\x18\xf5\xbd\x65\x78\x36\xa3\x86\x09\xc2\xda\xba\x7f\x9c\xe1\x1f\x1b\xf0\x25\x0e\x97\xee\x20\xc1\x2f\x29\x7c\x54\xee\x9f\x38\x60\xd2\x14\x7b\x22\x57\xac\xc2\x3f\x39\x84\x90\x65\x8b\x42\xee\x9f\x33\x2c\x70\x53\x98\x8a\xf4\x37\x76\xfe\x1c\xba\x7f\xf0\xd9\xe6\xd3\x30\xd5\x32\x72\x07\x4d\xad\xb6\x3c\xc6\xa7\x9f\xcf\x8c\x36\xc4\x6c\xfa\xf3\x2e\x66\xb4\x31\xfc\x59\xae\xfc\x5a\xcc\xa8\xcf\x82\x8e\x19\x3d\x99\x18\x59\x2f\xf6\x8b\x7d\x78\xf1\xbf\x02\x00\x00\xff\xff\xa7\x0b\x5a\xe2\x00\x30\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 24576, mode: os.FileMode(420), modTime: time.Unix(1464200428, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"init/rails/.dockerignore": initRailsDockerignore,
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/.dockerignore": initRubyDockerignore,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/.dockerignore": initSinatraDockerignore,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"init/unknown/.dockerignore": initUnknownDockerignore,
	"init/unknown/Dockerfile": initUnknownDockerfile,
	"init/unknown/docker-compose.yml": initUnknownDockerComposeYml,
	"templates.go": templatesGo,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"init": &bintree{nil, map[string]*bintree{
		"rails": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRailsDockerignore, map[string]*bintree{
			}},
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{
			}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{
			}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRubyDockerignore, map[string]*bintree{
			}},
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{
			}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{
			}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initSinatraDockerignore, map[string]*bintree{
			}},
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{
			}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{
			}},
		}},
		"unknown": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initUnknownDockerignore, map[string]*bintree{
			}},
			"Dockerfile": &bintree{initUnknownDockerfile, map[string]*bintree{
			}},
			"docker-compose.yml": &bintree{initUnknownDockerComposeYml, map[string]*bintree{
			}},
		}},
	}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{
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

