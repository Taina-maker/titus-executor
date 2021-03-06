// Code generated by go-bindata. DO NOT EDIT.
// sources:
// default.json (12.073kB)
// fuse-container.json (13.415kB)

package seccomp

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _defaultJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5a\x4d\x6f\x1b\x39\xd2\x3e\xcb\xbf\xc2\xd0\x39\x03\xd8\xb2\xe3\x71\x72\xd3\xab\xf8\xdd\x09\x36\x8e\xb3\xb6\x77\x67\x06\x8b\x80\xa0\xbb\xab\x5b\x5c\xf1\xcb\x2c\xb6\x6c\x61\x90\xff\xbe\x60\xb7\xd4\x2c\xb2\x35\x41\x14\x25\xf6\x2c\x72\xb0\xc1\xe7\xe1\x57\x55\xb1\x58\x2c\xb2\xf5\xc7\xc1\x68\x5c\x42\xc5\x1b\xe9\xa7\x85\x17\x46\x8f\x5f\x1f\x8e\x6f\x66\x97\x1f\xd8\x74\x76\xcb\x2e\xae\xaf\xdf\x5f\x8d\x5f\x1c\x8c\xc6\xdc\x15\xf3\x4b\x6e\xc7\xaf\x0f\xff\x7d\x30\x1a\xfd\x71\x30\x1a\x75\x9c\xf0\x50\xf8\xc6\x41\xec\x76\x3d\xfb\x85\xfd\x76\x7e\xc6\xce\x4e\x43\xc7\xd1\x68\x8c\xcd\xdd\x94\xb4\xc4\xf5\x18\xa3\x51\xda\xa1\x6b\x9d\x90\x27\x93\x71\xe0\x3e\x1e\x8c\x46\x9f\x5e\x7c\xc1\xb4\xd3\xf0\x7f\xa7\x79\xa7\xd7\x97\xbb\x4d\x71\xf9\xf6\xc3\xcd\x4e\x33\x84\x0e\x43\xd5\xba\x61\xde\xef\xaa\x60\xec\xf6\x6d\x04\xd8\x7d\xf6\x8b\x77\x3b\x6b\x7f\xf1\x6e\xfb\xf4\x61\xa8\xaf\xb2\xc0\xa6\xe3\xb7\x13\x63\x37\x19\x6e\x4e\x5e\x1d\xfd\xb6\xc3\xec\xa1\x7d\x9c\xe1\x60\xf4\x31\x6c\x28\x5c\x61\xc1\xa5\xc4\x64\x47\x69\xae\xe8\x10\xbc\x28\xc0\xfa\x8d\xd8\x1d\x3a\xa5\x10\xb1\x47\xe5\x7f\xbc\x50\xf0\xd8\x63\xc9\x9d\xda\x80\x3b\xa1\xcb\xbe\xec\x16\x9b\x62\xc1\x6d\x0d\x9e\x20\x24\x68\x5e\x0a\x17\x81\x32\x65\x04\xe6\x41\x27\x60\xb3\x14\xa3\x71\x21\x4d\xb1\x60\x35\xf8\x60\x89\x9c\x0b\xf2\xa5\xa4\xe6\xda\xa0\x04\xb0\x84\xc6\xd8\xc6\x68\x0d\x45\x94\xc8\xd8\x15\xab\x84\x04\xe6\xb8\xae\x63\x2b\x07\xbc\x6f\x53\x36\x96\x14\x27\xa4\x7c\xb2\x29\x83\x35\x52\xb2\xb6\x17\x6c\xe3\x8e\x33\xd2\xcb\x01\xc1\x8c\x2c\x53\xd2\x3e\x70\xe1\x53\x6a\x3b\x93\x74\x5d\x82\xf6\x55\x0e\x7b\xa1\xe1\x11\x8a\x25\xa4\x28\x6a\x0a\x8f\x22\x29\xb3\xda\x99\xa8\x7c\xd5\x39\x47\x6c\x5e\xf1\x72\x29\x10\xce\x4e\x07\x04\xa3\x9c\x94\xa6\x20\x76\xa9\xb8\x36\x5e\x54\x2b\xa6\x78\xf4\x9a\x2a\x71\x8d\x2a\xf1\x8d\x0e\x91\x69\x13\x67\xa9\x32\x6f\xe9\x30\x6d\xad\xa3\xb5\x5b\x40\x64\x2b\xb9\xe7\xb8\xd2\x45\x4f\xd4\xe0\x1f\xb9\xf7\x51\x12\x29\x30\x67\x4c\x11\xc5\x36\x44\x05\x07\xca\x2c\x21\x6d\x8c\xf9\x78\xe8\x89\x68\x01\x10\x69\x02\x1c\x10\x15\x66\x90\xd6\x53\xd1\xbd\x6b\x74\x62\xe7\x0d\x41\x3a\x34\x3e\x6e\xe6\xaa\x09\x7b\x87\x2c\x67\x0d\xbe\xb0\x0d\x45\x0f\x25\x41\x25\x68\x8f\x39\x8e\x63\xd7\xe0\xa1\x16\x65\x06\xe3\xba\x04\xa2\x49\xeb\x9b\xac\x3e\xed\x9e\xf7\x6e\x7d\x11\x07\x44\xd2\x46\x04\x95\x1c\x21\x2c\x80\x0b\xf1\x8f\x52\xe9\x34\xb6\x76\x96\xc2\xb4\x32\x83\x4e\x18\x27\xfc\x8a\x50\x8e\xeb\xd2\x28\x4a\x00\xa6\x13\x74\x44\x22\xa6\x03\x6c\xf2\x36\xb9\x31\x9c\x14\x4a\xd0\xb5\x61\xce\xdc\x35\xe8\x59\xf0\x48\xda\xae\x41\x5e\x53\xfd\x30\x19\x19\x4d\xb1\xc8\x0c\x10\x28\x63\x93\xa1\xfd\xdc\x01\x2f\x19\x77\xc0\x09\xed\x93\x91\x82\x69\x4d\x55\x72\xaa\x7d\xaa\x45\xae\x42\xe2\xf9\x62\xbd\xeb\x79\x59\xb2\x07\xee\x8b\x79\x5e\x21\x74\x54\x97\x72\xc7\x39\xe9\x54\x36\x80\x61\x05\xd7\x05\xc8\x48\x90\x00\x2b\x0c\x2b\x01\xbd\x33\x2b\xc2\x04\xef\x5b\x52\x7f\x16\x26\xac\x2d\x23\x27\xd7\x9a\x41\xca\x04\x14\xe3\x61\xc0\xcd\x1d\x59\x24\x61\xfb\xdd\xb8\x10\xb2\x97\x40\x26\xf1\x4a\x66\xf1\x4a\xe6\x86\x92\x42\x2f\x68\x39\x6e\xd0\xb0\xf2\xa0\x29\x4a\x3b\x0e\x18\x26\x25\x02\xc4\xd1\xb6\x44\xa8\xb4\x41\x1e\xae\x24\x0d\x57\x32\x0d\x57\xaa\x0b\xf6\x3d\x04\x55\x95\xd9\x11\xa8\x84\x2e\x8c\x8b\x70\x41\x62\x7c\x0b\xe2\xe0\x6a\xa1\x63\xc4\x6f\x01\xa9\xa3\x21\xb7\x05\x93\x04\xf1\x68\x6b\xa5\xb8\xa5\xe5\xd8\xd0\x3a\xe3\xc9\xc1\xaf\xee\xc3\x62\x23\x78\xaa\xae\xba\x67\x9d\x8f\x11\xc2\xd8\x68\x72\x75\xcf\xc2\x2e\x28\x1d\x14\x20\xe2\x31\xba\xa1\x11\x62\x46\xa4\xee\x59\xa3\xe9\x52\x2a\x07\x54\x36\xac\x89\x8f\x2a\xac\x89\xe7\x29\xac\x5d\xb1\x24\x08\xc9\xb0\x34\xe4\xab\x46\x27\x96\xe9\x20\xb5\x46\xa3\xc9\x9c\x83\xc4\x48\xc3\xc3\xfa\xcc\xe9\x1d\x46\xc3\x03\x82\x24\x76\xa2\xea\x87\x72\x6c\x6b\x79\x13\x57\xdf\x0a\x9b\x94\x7b\xbb\x87\x24\xa5\x2f\x27\xc0\x11\x03\xd8\x10\x7e\xa2\x6b\xb5\x70\x99\xa2\x38\x62\x17\x19\x49\xeb\x4e\xe2\xfe\xae\x67\x1f\x9c\xa0\xa7\x5e\x87\x97\x19\xec\xc7\x0b\xa3\xd3\x32\x9f\x67\x04\x5d\xc5\x0d\x8e\x66\x48\x44\x75\x50\x24\xe5\xca\xc5\xb3\x21\x60\xa5\xb0\x4e\x30\x85\x8a\xdb\x2e\x15\xb5\xbc\x8e\x99\xee\x96\x3d\xeb\x80\x46\xf4\x0e\x51\x81\x3a\x4c\x14\x44\xcf\x9d\x67\xeb\xcb\x41\x4f\x2b\xb2\x1b\x43\xb5\xa8\x79\x77\x59\x4f\x38\x0b\xba\x14\xba\xce\x48\x67\x0a\xc5\x71\x91\xb2\xf7\x0d\x34\x20\x74\x65\x52\xda\x81\x6f\x5c\x36\x2a\x36\x68\xc9\x6e\xe9\xc8\x76\x13\xd1\x14\xd7\x79\xe6\xeb\x6d\x03\x63\x31\x87\x32\x6c\x60\x5e\x55\xe1\x8c\x58\x0d\x2b\x88\xbd\x7a\xd2\x72\xc7\xd5\x80\x65\x9b\x63\x9d\x29\xfe\xf8\xb9\x5a\xa1\x07\xb5\x6d\xa1\x91\x90\xcd\xe5\x5c\xdb\x55\x68\x0f\x6e\xc9\x65\x5a\x89\x7f\x26\x36\x6e\x13\x1b\xb7\x8a\x8d\x7f\x3a\xf5\x4a\x40\xbc\x0d\x20\x14\x85\x51\x36\x42\xba\xb5\x11\x14\xd9\x82\x08\x8a\xc4\x20\x04\x65\x48\x37\xd5\x2e\x0d\x65\x74\x49\xcb\xc1\x6d\x73\x1c\x37\x60\x60\xa8\xe7\xb7\x38\x85\x3e\x2e\x2d\xf8\x8a\x66\x50\x1b\x1c\x4f\xcc\x96\x69\xb2\x16\x4d\xd6\x22\x1d\x21\xef\x9f\xe6\x92\x38\xcc\x25\x31\xcf\x25\x31\x4d\x1c\x71\x98\x0c\x62\x48\xe3\xd2\x26\x0e\xf2\x89\xd3\xec\x10\x87\xd9\x21\xe6\xd9\x21\x0e\xb3\xc3\x96\xca\x9b\x0c\x5a\x24\xf9\x23\x6e\xcf\x1f\x31\x49\x17\x71\x7b\x26\xd8\xd2\xa2\x0c\x99\x9b\x23\x8f\x03\x98\xa4\x7f\x98\xa5\x7f\x79\x26\x81\x73\x15\x43\x14\xce\x13\xcf\x9b\xab\x92\x56\x51\x37\x9c\x37\xbe\x24\xc9\x53\x88\x51\xd2\xa3\xe7\xf1\xd4\x43\x51\x6b\x2e\xe3\x8d\x77\x83\x4f\x09\x91\x46\xa0\x90\xff\x92\x29\x5a\x44\x03\x63\xc7\x58\x1e\xa3\x23\x5a\x29\x8a\xe8\xe0\xe4\xc0\x4c\x53\xa2\xf4\xbe\x96\x5f\xd7\x70\xa5\xe8\x51\xb2\x86\x64\x2c\x72\xb4\x87\xf2\x96\x77\x89\x40\x93\x09\x56\x98\x04\xc5\x15\x4a\xd3\xef\x2a\x0f\x7d\x2f\x5f\xd3\x74\x94\xbe\x98\xb4\x2e\x9e\x25\x6d\x1d\x57\x82\x84\x8c\x1b\xa4\x77\x1b\x36\x7b\x86\xd9\xd0\xb8\x85\x0e\x6d\xcd\x12\x9c\x6b\xf4\x80\x1f\x36\xde\x32\x42\xaf\xbc\x4f\x74\xca\xee\xbe\xc3\xab\x6f\x33\xb8\x52\x35\xf4\x00\x6b\xe8\x89\x9a\x26\x6e\x1d\x8a\xcb\xd4\x50\x91\x5a\xa0\x31\xab\xed\x85\x5c\xd2\x07\x82\xa5\x4a\xfd\x28\x9c\x74\xa7\x14\xc4\xdd\x14\x10\xb9\x7a\xb6\x19\x4b\x02\x96\xdd\xd3\x5f\x4b\x8d\xf9\xf0\x8d\x7d\xfa\xee\xdd\xd5\xaf\xe3\x75\xb5\xab\xdb\xe7\xbf\x75\xeb\xc2\x28\x05\xda\x87\xe6\xeb\x06\x42\x17\xb2\x29\xdb\x37\xc2\x3f\x3e\x75\x14\x3c\x12\x2a\x7d\xc2\xcc\x9e\x13\x31\x5e\x28\x77\x17\xa9\x1d\xa2\x1d\xb6\x15\xa3\x84\xc7\xf1\xeb\xc3\xe3\x17\x6b\x62\xc9\x65\x03\xe3\xd7\x87\x93\xb3\x93\x84\xba\x7d\x30\xe3\xd7\x87\x47\x1b\xce\xd8\x7e\x9e\xf0\xf7\xfe\xa2\x15\x64\xf4\x89\x88\x43\x74\x9e\xfd\xeb\xe2\xa7\xc9\xd1\xe4\xe8\xa7\xe3\xd3\x93\xcd\xf7\x81\xbf\x9a\x05\x26\xb9\x05\x8e\x5f\x6e\x31\xc0\x5f\xc7\x02\x16\x1c\x1a\xcd\x65\x38\x0d\xbf\x91\x09\x8e\x72\x13\x1c\xed\xe0\x02\x17\xff\xf8\xac\x01\xfe\x57\x54\x3e\xff\xf1\x54\x3e\x3e\x39\x3e\xfa\x79\xf2\x63\xea\x7d\xfe\x03\xba\xf8\xe9\xe4\xd5\xe9\xab\xb3\x9f\x27\xaf\xb6\x05\xb8\x67\xd4\x3d\xcb\xbf\x26\xdf\xeb\xb8\xed\x66\xe3\xae\x98\x93\xd9\x47\x63\x6b\x8b\xb3\x53\x09\x9d\x9a\x1f\xc3\xff\x9d\x35\xe0\x4e\xb1\x6d\xdf\x63\x02\xff\x99\xec\x32\x51\x7a\xcd\xdf\x39\xe0\x0b\x6b\x84\x26\x9f\xf5\x8a\x39\x54\xb2\xc1\x79\x72\x49\x90\xf8\xc4\x76\x8a\x9f\x23\xdb\xf2\xfa\x9b\xeb\x57\x5b\xac\x98\xb3\xee\x59\xea\x89\xb5\x50\xf1\xed\x6b\x34\x7e\x5c\x7f\xbd\xfe\x4a\x2d\x94\x29\x45\xb5\x62\xb2\xfc\x9a\x6c\xe4\xdb\x6a\xd1\x97\xcf\xcf\xf6\xd1\x08\x4f\x5e\x1d\x31\x5b\x08\xa6\x94\x30\x8c\x3e\xd7\xa5\x35\x49\x96\xdc\x56\xb9\x46\x87\x74\x9c\x09\x8d\xde\x3d\xb1\x39\x82\x00\xbd\x05\x02\x78\xdc\xc7\x06\xc6\x82\x66\x77\x2b\x36\xe7\xba\x94\xc0\xf8\x77\x5e\xdc\x82\x5b\xaa\xcb\x6c\xfa\x81\xbd\x99\xce\xd8\xf5\xc5\xf4\x0d\xbb\xb9\x98\x5e\xcf\x7e\xd9\x47\x99\x3b\x5b\xf5\x81\x44\x1a\x3d\xfc\x32\x4c\x3f\x05\x49\x63\x16\x8d\x65\x65\x61\xcc\x42\xc4\x27\x77\xd3\xc4\x68\x14\x86\x67\xde\x10\xe3\xbc\xe8\x4f\xb0\x8a\xb5\x1f\x7a\x92\x67\xfc\xfb\xc6\x78\x9e\x3c\x7d\xf9\xd2\x28\x2e\x34\xbd\x04\x22\xf8\xb9\x41\x9f\x51\x1a\xe3\xf5\x91\x8a\xd0\xa1\x49\xbc\x33\xe2\x9c\x3b\x78\xf2\x45\xba\xf9\xfd\x86\x4d\xdf\x5c\xbe\x7d\xbf\xcf\xf2\x74\x6b\xf2\xbd\x4e\xfc\xc9\xd1\xf9\xd1\xcb\xa3\x97\xe7\x2f\xcf\x76\x38\xf1\x2f\xa7\x37\x7f\xbf\x78\xf3\x6d\x0e\xfe\x2f\x36\xde\x8b\xaf\xda\xd9\x4f\x67\xda\xe1\x3d\xf9\xfb\x99\x36\xa8\x79\xd8\xbe\x43\x83\x07\x77\x68\x5c\x09\x4e\xe8\xfa\xb0\x32\xee\xb0\x55\xea\x50\xe0\x61\x29\xaa\x0a\x1c\x6c\x76\xc5\xf7\x88\x92\xbb\x79\xfe\x67\xd7\xc2\xc1\x9d\x31\x4f\x1f\x47\x83\xa0\xff\x77\x75\x75\xbb\xd7\x0e\x9d\xbb\xe7\x12\x7d\xf6\xcb\xf5\x9e\xc2\x77\x2f\x8b\x4c\x99\xb2\x89\x1f\x0d\x42\xc4\xcf\xa8\x6a\x0b\x77\xdf\x80\x5b\x6d\xb8\xe7\x50\xff\xf2\xea\xcd\x3f\xdf\x5d\xec\x95\x65\x16\xc5\xf3\xac\xdc\x87\xe9\x6c\xb6\xd7\xc2\x2d\x8a\xf8\x25\xc9\x3a\x53\x00\x22\x5b\x2a\x96\x7d\xab\xed\xf9\xec\xab\xab\x77\xbc\x78\x9e\x35\xfb\x70\x7b\x3d\x9d\xed\xb5\x66\xc2\x58\xf2\x7b\x12\x0b\x4e\x3d\x8b\x22\xd7\xd3\x5f\xdf\x5e\xed\x95\x49\x6f\xf9\x15\x0f\x0e\x7f\x41\xb9\x79\x76\x7f\x0e\x1d\x6f\xdf\x5e\xee\xb5\x54\xcb\x39\xd7\x75\x63\x9f\x47\xf6\xdb\xdf\xd9\xec\xea\xfd\xff\xbf\xfd\xdb\x17\x68\x70\x30\xfa\x78\xf0\xe9\xbf\x01\x00\x00\xff\xff\x4d\xd9\x41\x03\x29\x2f\x00\x00")

func defaultJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultJson,
		"default.json",
	)
}

func defaultJson() (*asset, error) {
	bytes, err := defaultJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "default.json", size: 12073, mode: os.FileMode(0755), modTime: time.Unix(1600885525, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xee, 0xcc, 0xab, 0xb7, 0x17, 0xd, 0x82, 0xf5, 0x94, 0x19, 0x44, 0xbe, 0xe3, 0xe4, 0x5e, 0x25, 0x9d, 0x96, 0x3b, 0x91, 0xf7, 0x10, 0x2f, 0x14, 0xd, 0xbe, 0x53, 0xcb, 0xf1, 0x22, 0x51, 0xd2}}
	return a, nil
}

var _fuseContainerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9a\x5b\x53\x1b\x3b\xf2\xc0\xdf\xf3\x29\x5c\x7e\xce\x03\x01\xc2\x81\xbc\xf9\x6f\xf8\x6f\xa8\x85\xc0\x02\x5b\xe7\x9c\xda\x4a\xa9\x84\xa6\x67\xac\xb5\x6e\xa8\x35\x06\x57\x2a\xdf\x7d\x4b\x33\xf8\xa2\x96\x7c\xce\x81\x38\x29\x6f\x6d\x1e\x42\xac\x5f\xb7\x2e\xdd\xa3\x4b\xb7\x66\xbe\xbc\x19\x0c\x86\x15\xd4\xbc\x55\x61\x24\x82\xb4\x66\xf8\x61\x30\xbc\x1d\x5f\x5e\xb3\xd1\xf8\x8e\x9d\xdd\xdc\x7c\xba\x1a\xbe\x8d\x4a\xdc\x8b\xc9\x25\x77\xc3\x0f\x83\x7f\xbd\x19\x0c\x06\x83\x2f\xdd\xdf\x67\x81\x0c\x20\x42\xeb\x61\x55\xf9\x66\xfc\x91\xfd\x76\x7c\xc4\x8e\x0e\xbb\xea\x9d\x26\xb6\xf7\xa3\x35\x65\x5c\xb6\xd5\x49\x93\x7a\xcb\x4a\xa9\xe0\x60\x7f\xf8\xcc\x3f\x77\xff\x7f\x7d\xfb\x57\xc7\x32\x8a\x7f\x5f\x33\x98\xd1\xcd\xe5\x6b\xfb\xbc\x3c\xbf\xbe\x7d\x4d\x97\xb1\x5e\xd9\x01\x7d\x8b\x9f\x5e\xef\x86\x55\x03\x5b\x1f\xd5\xb7\x0c\xe9\xec\xe2\xb5\x7e\x3a\xbb\xd8\x3c\xa6\xd8\xea\x37\xfa\x6a\xd1\xc4\x77\x19\xdb\x6b\x07\x76\x7b\x70\xb2\xf7\xdb\xcb\x87\x14\xab\x91\x2e\xdf\x0c\x06\x9f\xbb\xd5\x8d\x73\x14\x5c\x29\xcc\x97\xb7\xe1\x9a\x36\xc9\x85\x00\x17\xd6\x6d\xeb\xc9\x21\x45\x88\x09\xa9\xfe\x1d\xa4\x86\xa7\x84\x29\xee\xf5\x3a\xb8\x97\xa6\x4a\xca\x7e\xba\x5e\x14\xdc\x35\x10\x08\x41\x42\x26\x95\xf4\x29\xd0\xb6\x4a\x81\x7d\x34\x19\x58\x7b\xcc\x11\x29\x2b\xa6\xac\x81\x10\x3d\x5a\xe2\xd1\x96\x5c\x60\xb8\xb1\xa8\x00\x1c\x11\x61\xaa\x6b\x8d\x01\x91\x8e\xda\xba\x39\xab\xa5\x02\xe6\xb9\x69\x52\x6d\x0f\x3c\xd1\xad\x5a\x47\x8a\xfb\xa4\x7c\xb0\x5e\x06\x67\x95\x62\x5d\x2b\xb0\x89\xbf\x2b\x08\x82\x2a\x42\x66\x55\x95\x0b\xdc\x23\x97\x21\xc7\x9b\x69\xd6\xcc\x0c\x4c\xa8\x4b\x28\x31\x0e\x9e\x40\xcc\x20\x27\xa9\x87\xe0\x49\x66\x65\xd6\x78\x9b\x3a\xae\xee\x27\x69\x5a\xb5\xe6\xd5\x4c\x22\x1c\x1d\x16\x21\xa3\x5c\x29\x2b\x88\x5f\x6b\x6e\x6c\x90\xf5\x9c\x69\x9e\xce\xde\x3a\x9b\x9a\x75\x36\x37\x7b\x42\x86\x94\x4d\xd8\xba\x30\x63\x7b\x46\x6b\x9a\xf4\x29\x76\x80\xd8\x50\xf1\xc0\x71\x6e\x44\x02\x1b\x08\x4f\x3c\x84\x74\xb4\x4a\x62\x89\x5a\x91\x9a\x69\x89\xd9\x1e\xb4\x9d\x41\x5e\x11\x4b\x7d\x60\x20\x26\x44\x40\x46\x1c\x51\x11\xd6\x58\x40\x54\x8f\x9a\x1a\x7c\x6b\xb2\x67\xb8\x80\xa4\x72\x1b\xd2\x0d\xac\x6e\xe3\x3e\x40\xa6\x50\x03\x41\xb8\x96\x92\xc7\x8a\x90\x0a\x4c\xc0\x12\x4b\xfb\x6c\x20\x40\x23\x69\xe5\x88\xd2\xe7\x1f\x61\x9b\xeb\xb5\x05\xbd\xbc\xb9\x52\x6b\xdd\x7a\xa1\x03\xec\x61\xa6\x2b\xa3\x1b\x3c\x81\x0e\xc0\xc7\xf3\x83\xe2\xbc\x7b\xd7\x78\x47\x51\xae\x54\x40\x5e\x5a\x2f\xc3\x9c\x60\xcf\x4d\x65\x35\x85\x80\x79\xc7\x3d\xcc\xcc\xf1\x80\xb9\x2b\x7b\x98\xeb\x2a\xa9\x25\x7d\xfe\xcc\xdb\xfb\x16\x03\x8b\x2b\x86\xea\xb7\xc8\x1b\xea\x13\xcc\x7a\x43\x2b\xa6\x05\xe7\x45\x6c\x5d\xd6\x5d\x98\x78\xe0\x15\xe3\x1e\x38\x11\x85\xac\xe5\xf8\xa8\x6c\x5d\x71\xea\xb5\xdc\xe2\x92\xb9\xd9\x8a\x95\xcf\xbb\x1d\xaf\x2a\xf6\xc8\x83\x98\x94\x84\xd2\xa4\x2e\x5a\xe7\xef\x4a\x02\xaf\x0b\x8d\x59\x26\xb8\x11\xa0\x52\x48\x0e\x2a\x69\x59\x05\x18\xbc\x9d\x13\x1a\x57\xc3\x8c\xae\x39\x69\xe3\x1c\x62\x24\xaa\x78\xa6\x48\x69\x24\xe9\x19\x12\x59\x7b\x4f\x26\x80\x74\xc9\x0e\x33\x95\x2a\x19\xa1\xca\xf6\x73\x55\xd8\xcf\x55\xc9\xd9\x4a\x9a\x29\x2d\xa7\x9b\x4f\x9c\x71\x60\x28\xc9\x1b\x2a\x52\xa6\x14\x02\xa4\x3d\x6c\xd8\xbd\x73\xc5\xd2\x76\xae\xe8\x76\xae\xf2\xed\x5c\xf7\x87\x6b\x82\x40\xd7\x55\x21\x64\xd1\xd2\x08\xeb\x53\x34\x25\x67\x6a\x07\xd2\x4e\xf5\xd4\xa4\xa7\x6c\x07\x88\x0e\x3d\xc6\x3a\xb0\x9f\x11\x9e\x3e\x4b\xad\xb9\xa3\xe5\xb4\x92\xf3\x36\x90\x60\x4f\x3f\xc4\x09\x87\x10\xa8\xbb\xf4\x03\xeb\xe7\x3f\x81\xd6\xa5\x8f\x54\x3f\xb0\xb8\x8a\x2b\x0f\x02\x64\x1a\x0e\x2d\x44\x08\x69\x24\xad\x1f\x58\x6b\xe8\xf4\xd1\x1e\xe8\xf8\xb1\x21\xeb\x49\x63\x43\x56\x87\xc6\xc6\x8b\x19\x21\x48\xba\xa3\xc7\xac\x6e\x4d\xe6\xe1\x1e\x51\x8f\xb6\x86\x8c\xa9\x18\x54\x1b\x78\x7c\x8e\x03\x92\x09\x6c\xe0\x11\x41\x11\x7f\x53\xf7\xc5\x72\x5a\xcf\xf1\x36\x9d\x81\x4e\xba\xac\x9c\x3c\xd7\x18\xc8\x26\xe5\x0c\x78\xe2\x48\x17\xb7\xe8\x74\xea\x77\x68\x96\x93\xb4\xa7\xfe\x84\x21\x35\x7b\x2b\x93\x9b\x12\xf7\xe8\x25\x8d\x58\x7a\x36\x2b\xa0\xa4\x8f\xd8\x2b\x2d\xf3\x49\x01\xd2\x19\xb4\x60\xa9\x3b\x33\xb3\x3c\x88\xac\x5c\xfb\xf4\x8c\x8e\x4c\x6b\x6c\x32\x46\x91\xe6\xae\x4f\x95\x1c\x6f\xd2\xec\x6c\xc3\x5e\xe5\x81\x9e\xa4\x3d\xa1\x83\xee\x19\x71\x0c\x06\xee\x03\x7b\xce\x90\x13\x91\x26\x3b\x4f\x54\x93\x0d\xef\xef\xd1\x32\xee\xc0\x54\xd2\x34\x05\x81\xb7\x42\x73\x9c\xe6\x92\x87\x16\x5a\x90\xa6\xb6\xb9\xc8\x43\x68\x7d\xa1\x17\x6c\xd1\x91\x95\xdf\x0b\xba\x4d\x81\xa6\x63\x3e\xb0\xd0\x6c\xea\x08\xc5\x04\xaa\xb8\x51\xf1\xba\x8e\xe7\xf4\xbc\x2c\x24\xbe\x5e\x0a\x1c\xf7\x5c\x17\x25\x6c\x11\xb6\x31\xcd\x9f\xfe\x4c\x43\x9a\xa2\x46\xf7\xa3\x55\x50\xe8\xdb\xfb\xae\x09\x69\x02\xf8\x19\x57\xb9\x02\xfe\x91\x49\x85\x5d\x79\x29\xd8\x60\x12\xfe\xe1\x70\xe6\x12\xd2\x4c\x17\x41\x08\xab\x5d\x8a\xe8\x96\x85\xa0\xc9\xf6\x81\xa0\xc9\x3e\x8c\xa0\x2d\x69\x46\x77\x8f\x99\x52\x53\xd1\x72\x5c\x3e\x25\x96\x6e\x1e\x91\xd2\x15\xd9\xb1\x1c\x85\x74\xea\x40\xa8\x69\xb4\xbd\x60\x69\x94\xd3\xd1\xb6\xa0\x99\xc5\x9e\x98\xe5\x2d\x58\xc8\x5b\xb0\x94\xb7\x60\x39\x6f\xc1\x52\xde\x82\x79\x82\x82\xe5\x44\x03\x63\x4a\x90\xab\xfa\x3c\x35\xc3\x52\xf6\x81\xe5\xec\x03\x4b\xd9\x07\x96\xb3\x8f\x0e\x97\x54\x8b\x9a\x59\x9e\x82\x9b\xf3\x14\xcc\x52\x12\x2c\xe6\x1e\xb8\x39\xf7\xe8\x44\xb2\x8a\x79\x81\x27\x57\x81\x98\x25\x1b\x58\x48\x36\x4a\xf1\x24\x4e\x74\xba\x69\xe3\x24\x5b\x27\x13\x5d\x51\x15\xba\x70\x26\x6d\xa8\x48\x08\x1e\x77\x6e\x15\x30\xf0\x34\x46\x41\xd9\x18\xae\xd2\x7b\xa9\x05\x3b\x24\x30\xdf\x93\xa3\xc7\x48\xd7\x1d\xa1\x47\x49\x4f\x1d\x4f\xcf\x13\x74\x4a\x8a\x74\x99\x92\x70\x27\x0f\xa6\xf3\x5b\x90\xd2\x25\x08\xce\x35\x3d\xc8\x9f\x11\x69\x9f\x04\x71\xb1\xbc\xe1\x96\x32\x8a\x48\xc7\x73\xcc\x8e\x94\x39\x2a\x9b\xec\x1d\x01\x92\x56\x42\x43\x93\x25\x7a\xcf\xda\x2d\xd8\x42\x6a\xd0\xf3\x0a\x14\x14\x78\x31\x99\x58\x48\x0a\x97\xb9\x0b\x11\x6e\x10\xc5\x3a\x76\x06\xde\xb7\xa6\x28\x2b\x57\xda\xd0\x5a\xe2\xb4\x90\xd9\x5f\xb8\x9d\x2a\x5f\x4e\xb5\xc5\xcb\x88\x96\x86\x16\x2d\x8d\x85\xf2\xb4\xa0\x27\xe9\x54\x68\xe9\xd0\x3b\x60\xb0\xa0\x95\x18\x34\xa3\xd7\x81\x33\x9d\xcf\xeb\x18\x97\x1c\x52\x90\xee\x10\x91\x90\x0b\xa0\x2e\x8e\xcd\xc0\x6c\xf9\x92\x63\x21\x19\xf2\xfc\xd5\xe6\xe8\xe2\xe2\xea\xd7\xe1\x4a\xc3\x37\xdd\x5b\x8e\x55\x1d\x61\xb5\x06\x13\x62\xa5\x95\x9a\x34\x42\xb5\x55\xf7\x42\xe4\xcb\xd7\x25\x85\xa7\x35\xda\x41\xfa\x32\xa7\xf0\x16\x45\x28\x6b\xe0\x5b\x86\xba\x6c\xea\xcb\xf2\x57\x37\xc2\x0a\x9e\x86\x1f\x06\x7b\x6f\xd7\xe9\x8c\xab\x16\x86\x1f\x06\xfb\x7b\xc7\x7b\xef\xf7\xde\x1f\xbf\x3f\xca\xc5\x77\x8f\x36\xab\x67\xdd\x72\x24\xf1\xdf\xe5\xe8\xf6\xef\x67\xa7\xec\xec\x1f\xc3\xa5\xd2\xd7\xcc\x82\xef\xed\x38\x07\x1e\xad\xe1\x2a\x9e\xc6\x3f\xd8\x7d\x05\xfa\x17\xbc\xf6\x3f\xeb\xae\xe3\x9f\xee\x7a\x89\xbb\xde\x1d\xbc\xdb\xfb\x65\xff\xa7\xcf\x5e\xea\xb3\xe3\x9f\xcb\xf2\x45\x3e\x3b\xdc\x3f\x39\x3c\x39\xfa\x65\xff\xe4\xfd\x7f\xa1\xdf\x48\x04\xba\xff\x83\x0e\xfb\xd5\x00\xb8\x17\x13\x32\xa8\xee\xf6\x4f\x1c\x1d\x2a\x58\x79\xe6\xf3\xf3\xaf\x6f\x34\x97\x7b\xcd\x36\xbd\x7c\x8e\xb2\x3f\x09\xc8\x13\x4f\xad\xc9\xee\x3d\xf0\xa9\xb3\xd2\x90\x6f\x29\xc4\x04\x6a\xd5\xe2\x24\xcb\xe3\x14\xee\x88\xa3\xd3\xef\x45\x7a\xb0\xfa\x94\x66\x8b\x7e\x17\x13\xd6\xdf\xe0\xee\x88\xdd\x3a\xbd\x3a\x1e\x0c\x86\x4f\xab\x6f\x9b\xb6\x66\xb7\xb6\x95\xac\xe7\x4c\x55\x61\xb7\xed\x4e\xc1\xf1\xd1\xd6\x1d\x81\x07\x27\x7b\xcc\x09\xc9\xb4\x96\x96\xd1\xeb\xf1\x54\x9a\xe5\x21\x9d\xd8\xb7\x26\x26\x43\x4c\x1a\x0c\x7e\x47\xbc\x19\x07\x96\xfa\x2e\x92\xa7\xad\x7b\xcf\x3a\x30\xec\x7e\xce\x26\xdc\x54\x0a\x18\xff\xf1\xb3\x49\x70\x97\x59\x3f\x1e\x5d\xb3\xd3\xd1\x98\xdd\x9c\x8d\x4e\xd9\xed\xd9\xe8\x66\xfc\xf1\x3b\x2c\xa0\xd6\x90\xfc\xbb\x4c\x7e\xfc\xc9\xb5\xc9\x25\xb7\xbf\xdf\xb2\xd1\xe9\xe5\xf9\xa7\xad\x3b\xc3\xc3\xbd\xb5\xbb\xf3\xec\xa3\xa1\xff\x77\x75\x75\xb7\x75\x3b\xc5\xc4\xef\x9a\x9d\xe3\x8f\x37\xdf\xc3\xd2\xfe\x7e\x8d\x69\x5b\xb5\xe9\x0b\x04\x69\x64\x28\xe0\x7a\x03\x7f\x68\xc1\xcf\x17\x7c\x97\xfc\x76\x79\x75\xfa\xcf\x8b\xb3\xed\xc7\x13\x42\xec\xd6\xfc\xb8\x1e\x8d\xc7\xdb\x9f\x1e\x53\x91\xbe\xd7\x72\xde\x0a\x40\x64\x33\xcd\x0a\x6f\xba\x97\xb2\xc2\x3b\xea\xe0\xb9\xd8\xad\x99\x71\x7d\x77\x33\x1a\x6f\x7f\x66\x48\xeb\xc8\xe7\x4b\x0e\xbc\xde\x29\xcb\x6f\x46\xbf\x9e\x5f\x6d\x3f\xc2\xda\xf0\x25\x1a\x96\xbf\xac\x5e\x5c\x9e\xef\x92\x63\xee\xce\x2f\xb7\x3f\x21\x66\x13\x6e\x9a\xd6\xed\x96\xa1\x77\xbf\xb3\xf1\xd5\xa7\xff\x3f\xff\xdb\x0b\xcd\x7d\x13\xf5\xbe\xfe\x27\x00\x00\xff\xff\x03\x9e\x72\x97\x67\x34\x00\x00")

func fuseContainerJsonBytes() ([]byte, error) {
	return bindataRead(
		_fuseContainerJson,
		"fuse-container.json",
	)
}

func fuseContainerJson() (*asset, error) {
	bytes, err := fuseContainerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fuse-container.json", size: 13415, mode: os.FileMode(0644), modTime: time.Unix(1600711441, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0x3c, 0x72, 0x2e, 0x2, 0xf6, 0x24, 0xe, 0xec, 0x45, 0x6b, 0x2f, 0x80, 0xf1, 0x66, 0xd7, 0xf4, 0x28, 0x82, 0x44, 0xe, 0x35, 0x6e, 0x8, 0xab, 0xf1, 0x25, 0x2d, 0x67, 0x99, 0x38, 0x76}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"default.json":        defaultJson,
	"fuse-container.json": fuseContainerJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"default.json":        &bintree{defaultJson, map[string]*bintree{}},
	"fuse-container.json": &bintree{fuseContainerJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
