// Code generated by go-bindata.
// sources:
// include/wt_raii.h
// DO NOT EDIT!

package wt_concurrency

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

var _includeWt_raiiH = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5c\x5b\x73\xdb\x36\x16\x7e\xf7\xaf\x40\xd9\xd9\x94\x52\x15\xbb\xee\xec\x66\x3b\x76\xe4\xd9\xd4\x71\xb7\x99\xa6\xf6\x8e\xed\x34\x0f\x69\x87\x03\x93\x47\x14\xc6\x14\xa0\x02\xa0\x63\x35\x93\xfe\xf6\x1d\x5c\x78\x01\x08\x52\x92\x2f\x9b\x74\xf9\x10\xc9\x04\x70\xee\x38\x38\xf8\x00\xe5\x4b\x42\xd3\xa2\xcc\x00\x45\xef\x09\x87\x4c\x92\x1c\xf8\xee\x3c\xda\xd9\xa9\x1b\x9e\x13\x26\x24\x07\xbc\x38\x6a\xbd\x13\xd5\xab\x9d\xb4\xc0\x42\xa0\xb7\xf2\x02\x84\x20\x8c\x1e\xee\xdc\x30\x92\x21\x42\x6f\x30\x27\x98\xca\x1f\x30\x29\x20\x7b\x4b\xe4\xfc\x35\xa1\x20\x62\x42\x25\xe2\x20\x27\x28\x65\x54\x48\x94\xce\x31\x1f\xa3\x19\x29\x60\x82\x4a\x2a\x48\x4e\x21\x43\x05\xa1\x30\x5a\x47\xe8\xed\x65\x72\x71\x72\x71\xf1\xea\xec\x74\x8c\x84\xe1\x3d\x41\x0f\x47\xbd\xd2\xe7\xc9\x9d\x88\xef\x7c\x99\xc1\x8c\x50\x68\x18\xbc\xbd\x3c\xfb\xe9\x04\x8b\x55\x3c\xa3\x23\xb4\xfe\xf9\x75\x47\xfd\xfb\x61\x83\x9e\xc3\x14\xd4\xa3\xe4\x4e\x38\x48\x34\x45\x33\x7a\x78\x27\x0a\x33\x14\x2b\x0a\xa3\x2d\x25\x6a\x28\x18\x39\x7a\x8c\x9d\x68\xab\x26\xc9\x0f\xaf\x5e\x9f\x24\x89\xfa\xf6\xfa\xd5\xe9\x49\x92\x8c\x0e\x5b\x14\x3e\x6e\xc5\xd9\x97\xe1\x63\x8f\x4b\xe2\xda\xbb\x9b\xf9\x05\x3d\x88\x6f\x3c\x52\xe8\x1e\x4e\x0a\x91\xba\xa3\xb7\x02\xa4\xd0\x90\xdb\x6a\xd3\x3d\xaa\xff\x3c\xa9\xc2\x8e\x54\xa2\x6d\xef\xca\xbf\x82\x23\xbf\xb8\x8f\x27\x3f\x73\x47\x56\x4b\xc7\x71\xc9\x05\xe3\xe8\xc3\xce\xb2\xbc\x2a\x48\x7a\xa0\xdb\xab\xd7\xe1\x44\x2f\x64\x76\x70\x20\x24\x27\x34\x47\x25\x27\xee\x8b\x94\xd1\x19\xc9\x47\xe8\xc3\x4e\x47\x75\x3d\xe7\x1d\xab\x34\x24\xcd\x97\xa7\x47\x6c\x09\x34\x49\x0d\xf3\xba\xb5\xe4\x64\x37\x4d\x84\xe4\xf1\x68\x82\x68\x59\x14\x4b\xc9\x27\x96\x53\xd3\xf0\xc4\x8e\x1b\x8d\x0e\x2b\x25\x5d\x5d\xec\x97\x27\x4f\x10\x93\x73\xe0\x23\x74\x80\x2a\x56\xfa\xc5\x6e\x45\xa0\x25\xbc\xd3\x80\xa6\x15\xf7\x20\x87\x27\x88\x2d\x81\x63\xc9\xf8\x34\xc0\xac\xa1\xd9\x50\x73\xa8\x1f\x6e\xc6\x54\x3d\x1c\x64\xc9\x29\x1a\xcb\x39\x11\x8e\x28\x7f\xd6\xda\x3a\x1e\x50\x49\xa9\xab\x5b\xd7\x39\xb6\xd3\xd3\xa3\x26\x2a\xab\x37\x69\xc1\x04\xc4\x9e\x89\x0d\xe7\xb6\x29\x2e\x93\xe3\x37\xe7\x17\x67\xe7\xe3\xda\x16\x4f\x8f\xe2\x91\x5d\xb5\x3f\xf8\x0a\x38\x9a\x77\x49\xe4\x20\x1d\x3d\x06\x46\xed\x8d\xc7\xa6\xdb\x18\xfd\xcb\x76\x93\xbc\x04\xc4\x28\x5a\x30\x0e\x08\xa8\xe4\x04\xc4\xae\xed\xb4\xa7\x3f\xaf\x18\x2b\x10\x85\xdb\x20\x97\x2f\x6a\xd5\x75\x8f\x4a\x73\x87\xab\x26\xb0\xe4\x70\x33\x4c\x40\xf7\x08\x12\xd0\x85\x90\x00\xf9\x13\xac\x54\x75\xf6\xec\xef\x89\x44\xd7\xb0\x0a\x44\x8b\xf2\x89\x4c\xae\x61\x55\x11\x9a\xe8\x8e\x41\x6a\xbf\xe0\xa2\x84\x9a\xde\x8d\xfa\xab\x97\xa2\x6e\x6d\x68\x9a\xce\x41\xaa\xdf\xaf\x24\x18\xca\xa6\xfa\x7a\x10\xba\x1c\x16\xec\x06\xe2\xde\x7c\xa1\xcb\xb6\x9a\xb0\xed\x1d\x9e\xe8\x95\xbe\xb9\x31\xa7\x4b\xd2\x34\x71\x90\x87\x9b\xf0\xc9\x7d\x53\x3f\x51\x0b\xc1\xa8\x33\x01\x6b\x7a\x56\x84\x76\x22\x34\x62\x5c\xe8\x3f\x1c\x61\x9a\xda\x95\xe2\x05\xdc\x51\x9e\x6a\x78\x40\xa8\x96\x10\x71\xdd\xad\xcf\x52\xaf\xa8\x8d\x96\xfb\x9a\xcb\xf3\xf7\x86\x06\x33\xb6\xc8\xdb\xc1\xd5\x33\x0d\xec\x24\x69\x62\x1e\xb9\x89\x4d\x45\x1d\xe6\xe9\xbc\x49\x74\xd3\xa9\xca\x25\xa7\x67\x97\x3f\x9c\xbd\x39\x7d\xe9\x27\x3e\x6d\xa5\x94\x95\x12\x3d\x7f\x8e\xa2\x5f\x15\x71\x44\x99\x44\x33\x56\xd2\x6c\x62\xe5\x55\x7e\x7c\xba\xbf\x8b\x7e\x82\xd5\x01\x8a\x54\xcf\x6b\x58\xa9\x0f\x3d\x1a\x68\x56\x1c\x3a\x44\xad\x96\x9d\x7c\x6d\xb5\x6d\x34\xee\xb7\xec\x50\x0e\xbe\xbb\x9d\x55\x61\x24\xb0\x37\xd1\xdc\x74\xfa\xf4\x88\x50\x01\xbc\x27\xd3\x29\x02\xb6\xbd\xe5\xa0\x09\xea\x4d\x32\x41\x7f\xd5\xb9\xa9\x9d\x0e\xda\xa1\xab\x45\xec\x4a\xae\x3d\x7b\x72\x8b\x53\xd9\x17\x1e\xc1\x90\x0d\xc7\x0c\x95\x08\xb8\x5a\x56\xfb\x22\xc7\x8d\x2f\xdd\x77\x30\x94\xf6\xf6\x1e\x35\x9a\x9e\xee\xfb\x8b\x6d\x4b\x34\x5f\x16\x3b\x06\x38\x0f\x0e\xfa\xdf\x44\x9a\x4d\xd3\xdb\xcc\x64\x3f\x16\xbd\x4c\x5f\xb3\x58\x72\x72\x83\x25\x1c\xf8\xa5\x42\x5d\x13\x7c\x3c\xec\xc0\x22\x9d\xe2\xd6\xbe\x57\xd5\xed\xf1\xd9\xe9\xe9\xc9\xf1\xa5\x2e\x70\x53\x46\xa9\x5f\x36\xa9\x77\xb6\x2a\xb5\x56\xd2\xaf\x5a\x65\x68\x44\x04\x2b\xb0\x24\x8c\x4e\x05\xc5\x4b\x31\x67\x32\x52\xc5\xa8\xed\x3e\xf2\x3d\x24\xe7\x9c\xbd\x47\x51\x25\x9b\x22\x8d\x66\x7a\x53\x10\xf5\x15\x55\x8d\xc0\x35\x32\xd2\xae\x61\x2b\xc1\x6c\xd9\x58\x31\xee\x56\xb1\xb6\xa5\xb7\x8c\xad\x61\x97\x76\x1d\xdb\xe1\xd8\x5a\xf0\x1b\x82\x2e\x87\x4e\x29\xdb\xc3\x19\xad\xa9\x65\x2b\xbd\x3b\xc5\x6c\x40\xc7\x40\x7c\x37\x71\x5d\xef\x2f\x6c\x19\x5b\xb7\x58\x71\x86\x0a\xda\x7a\x03\xb4\x51\x45\xeb\x58\x20\x40\xc4\xd4\xb4\xdb\x8c\x56\x33\x2a\xe5\x80\x25\x5c\xe2\xab\x02\xe2\xb5\xdb\x2f\x34\x45\x51\xe4\xcc\xb8\xa6\x83\x81\x0e\x91\x10\xad\x2c\x29\x74\xd2\xba\x86\x55\x32\x63\x7c\x81\xe5\xf4\xf7\x89\x9e\xfb\xcd\x9f\x05\xcb\xa7\x31\x50\xc5\x3e\x9b\xce\x70\x21\x60\x34\xd1\x09\x4c\xb1\x6b\x28\xa9\x44\x78\x37\x62\xef\x39\x91\x90\x48\xb2\x00\x21\xf1\x62\x99\x94\x02\xe7\x30\x65\x3c\x03\x0e\xd9\x04\x0b\xb5\xf0\x4c\x63\xaf\xd7\x94\xd1\xd1\xe4\x06\xf8\x15\x13\x30\x7d\xe7\x35\xfe\x16\x12\xd0\x33\xf3\xd3\x23\x63\xd7\x56\x3c\x38\x5b\x4d\x21\x76\xf5\xb7\xea\x4d\x77\x65\x32\xe3\xeb\xe2\x25\xe8\xa0\xd0\x92\xeb\x0b\xe0\x44\x71\x8f\x34\x03\x46\x2d\x43\x46\x8d\xfa\xc4\xd5\x52\x6a\x04\x82\xe5\x9f\x40\x58\x2f\x02\xd4\x36\xcd\x97\x55\x6f\x0d\x70\x21\x81\x6b\x59\x5f\xb3\x3c\x57\xd5\x6c\x27\xf2\xf5\xfe\xcb\x10\xf2\x53\x44\xf7\xed\x66\x09\x42\x73\xed\x0b\x88\x28\x24\x79\x3b\x75\x20\x28\x04\x3c\x26\xcf\xca\xb5\x7d\xf9\xca\xb3\xdc\x31\x5b\x2c\x88\xbc\xac\x26\xc5\x0b\x3d\x91\x08\xa3\x62\x38\x8b\xdc\x00\x27\x33\x92\xea\x55\xed\x17\xbf\xb2\x1b\xc0\x73\x6a\x15\x82\x6f\x3b\x8a\x3a\xbd\xfc\xa7\x6d\x84\xc1\x8e\x71\x54\xe5\x87\x54\x6b\xdb\x4a\x10\x11\xfa\xba\xab\x0b\xfa\x1a\x45\xa3\xa8\x99\xd4\xa3\x5e\x20\x61\x89\x39\x5e\x20\x0e\xbf\x97\x84\x43\x86\x52\x4c\xd1\x15\xa0\x77\xb8\x78\x8f\x57\x62\x82\x28\xdc\x00\x9f\x20\xca\x28\xfc\xe6\x02\x0b\x6a\xae\x35\x4e\x38\x37\x04\x3c\x5f\xac\x71\x41\xc5\x75\xeb\x3c\x3e\x60\x8c\xe7\xcf\x1b\x65\x54\xd7\x51\x34\x90\x1a\x07\x83\x72\x7d\x66\xcc\x38\x5b\x6e\x9f\x0e\xd5\xa8\x3e\x9e\xd5\x52\x1d\xc4\xde\x74\x1d\x65\xa1\xaf\x0d\xf0\xc9\xce\x12\x69\x25\xa9\xe1\x33\x47\x86\x0a\x6b\x1c\x60\xfd\x3d\x4e\xaf\xcb\x65\x00\x7b\xeb\x27\x1c\x5d\xe9\x31\x07\xd1\x44\xc9\x12\x28\xa2\x71\xf6\xa2\x09\x17\xa3\x55\x59\xd5\xd4\xb5\x5b\x7d\x64\x41\x8f\xbb\xbc\x38\xd6\x02\x1b\xf4\xe1\xdd\xfe\x3f\xd0\xde\x58\x37\xb4\xc2\x01\x8d\xf7\xd0\xd7\x68\xff\x99\x6a\xda\x7f\x86\xe6\x70\x8b\x33\x48\xc9\x02\x17\x28\x23\x39\x91\x42\x77\x18\x9c\x79\xea\xd9\xff\xa7\x22\x30\xe1\x6a\xa3\x93\x48\x96\xb0\x22\x03\x21\x2d\x75\xcd\x57\xa7\x2c\xcb\x4d\xfd\x2d\x39\x26\x85\x72\x85\x72\x28\x1a\xef\xfd\xd6\x04\x21\x2e\x25\x43\x82\xfc\x01\x68\x6a\x9d\x46\x97\x9c\x50\x39\x8b\xbb\x5a\x0d\x27\x05\x3b\x25\xc8\x1f\xc0\x42\x83\xd7\xa4\x14\xf3\x44\x9e\xc5\xfe\x56\x14\xb7\x1d\x45\xb5\x76\xd1\x46\xc2\x48\x2c\x49\x9a\xa4\x58\xc8\xe7\xcd\x99\x25\xa3\xb9\xfe\xe7\x28\x6e\x5c\x1a\xd8\x17\xd5\x33\xa4\xa9\x91\x24\xc7\x54\xe0\x54\xe5\xb5\x56\x54\x05\x74\xed\x84\xd6\x15\xe4\x84\x3e\x7e\x6c\xc5\xdf\xa1\x31\xfa\x76\xa4\xda\xbf\x73\xe2\x4b\xd1\xc2\xa9\x04\xbe\x69\x8c\xfd\xe5\xc3\xe6\xf1\x02\x64\x7d\x75\xa1\xbd\xbd\x79\xb4\x74\x83\xef\x9b\x70\x00\xc5\xba\xfa\x22\x39\x65\x1c\xfe\xc3\x61\x89\xb9\xb2\xbf\x29\x50\xbc\x62\xcc\xe9\xd4\x03\x59\x6c\x26\x70\x64\x48\x25\x4b\x43\x4b\xd7\x60\xd1\xba\x0a\x6c\x2b\x0e\xee\x32\x83\xfc\x02\x4b\xa7\x66\x56\x14\x2a\x79\x07\xd1\xb4\x9a\x49\xd5\x6b\x23\x3e\xed\x22\x5d\x2f\xdc\xfa\xb2\x86\xf2\x38\x9a\xa2\x6f\x7c\x7b\xaa\xf7\x41\x40\xb3\xb7\x38\x40\x4d\x81\x10\x2e\x0c\xf4\xf0\x39\xdc\xaa\xef\xaa\xc9\x1d\x6a\xf6\xab\x81\x8d\x66\x55\x08\x04\x91\xab\xd6\xce\xc1\xf2\x0c\x5a\x42\x91\xf2\x0a\x89\xcd\x5d\x39\x48\x79\x03\x5f\x9a\xf1\x36\x34\xb3\xb8\x79\x75\x29\xcc\xad\x93\xac\xe4\xaa\x8c\xb9\x14\x8f\x68\xf0\x8a\xa3\xee\x3e\xb1\x1c\x87\x06\xd4\x42\x7d\x5a\x37\xb5\xec\x68\x27\x64\x1d\xb6\x5b\xd7\xad\x96\xc0\xe6\x71\xb9\x85\xb2\xf6\xf6\x10\x9a\xb6\x74\xad\xf9\x6d\x13\x93\x2e\xce\xab\xa2\xe6\x1c\xa4\x05\x73\x15\x83\x30\x98\x3b\x04\xca\x43\x6b\x19\xbe\xb3\xed\xb6\x9d\xd2\x5b\x98\x6e\xeb\x02\x64\x38\x4c\x24\x2f\x69\x8a\x25\xbc\x98\xa9\xdd\x45\x73\x72\x2e\x24\xe6\x72\x30\x9d\x56\x23\xbb\xd3\x7b\x62\x46\xef\x6a\x68\xad\xf5\xb6\x93\x61\x7d\xf0\xb8\x06\xe5\x1a\xdc\xad\x81\x8f\x7f\x06\x89\x33\x2c\x71\xf0\x82\x84\xdb\xe8\xa0\xa4\x35\x34\xd9\x42\x66\xb5\xa5\x35\xa0\x1d\x42\x85\xef\xb1\xa1\x6e\x5f\x97\xa8\xde\x56\x76\x88\x16\x56\x48\xb5\xc3\xf0\x8d\xd2\x7b\x5d\xc2\x53\xbb\x41\x82\x4d\xcc\xb8\xcd\x35\x26\x3c\x45\x19\x14\x20\xe1\x30\x44\x24\xf6\x06\x55\xa3\x6a\x9d\x82\x96\xf2\xd0\xec\x51\x7d\x3a\x11\x57\x7e\x45\x1f\x3e\xe9\xe1\x7f\xfb\x26\x62\xe0\xcc\xb9\xdd\x3c\x74\xdc\x17\x3e\x7c\xb9\xcb\xe9\xb3\x27\x50\xf7\x5c\xf7\x7e\x22\x6d\x7f\x1e\xe4\xce\xb7\xfa\x38\xc6\x45\xb9\xd7\x1d\xe2\x1c\x33\xda\x3d\xc1\x51\x2f\x9d\x7d\x7e\x76\xb5\xc4\x72\xde\xbb\xd5\xb7\x50\xa1\x90\x1a\x14\x9b\x43\x7a\xbd\x64\x84\xca\x33\x7a\x5c\x30\xa1\x4a\x2c\x8d\xe5\x39\x01\x19\xee\x1a\xf7\xbc\xdf\x3e\x65\x6b\x2c\x75\x92\xe2\x74\x0e\x89\xda\x70\x4c\xf7\xff\xfd\x7d\x08\x19\x75\x4f\x23\xed\xa5\xa7\x02\x68\x2e\xe7\xf1\x08\x1d\xb9\xc5\x61\x8b\x41\x83\x7f\x93\x3c\x78\x12\x9d\xcc\x31\xcd\x0a\xe0\xbb\xe6\x33\x01\xce\xc3\xb7\x8d\xfc\x8e\x0b\x10\x02\xe7\xb0\x49\xd7\x25\x67\x39\x07\x21\x36\xe9\x9b\x5a\x57\x74\x3a\x2a\xbd\x9b\x0b\xd2\x89\xca\x79\xb1\xf1\x77\xfb\xf2\x97\xa5\xd6\xc5\xa6\x74\xaa\x63\xb4\xf7\x24\xce\x23\xbd\xf6\x30\x4e\xc7\x9e\xf9\xe8\xcb\x7f\x7f\xda\x5e\xdd\x73\xab\xf5\xd1\x83\x42\x27\xca\x6f\x84\x0a\xe8\x7a\xe9\x15\x2a\xb5\x69\x7b\xed\x46\x03\x07\xc8\xe6\x7a\x84\x3e\xc1\xac\xae\x6f\xe9\xb3\xcb\xa8\x14\xed\x3a\xcb\xec\xa1\x7c\x19\xac\x1c\x5d\x23\x19\x3f\x75\xac\x84\x9c\x93\xe6\x30\x12\xbe\x99\x3c\x06\x52\x79\x68\x81\x5c\x27\x56\xb9\x28\x87\xe0\x21\x63\x8d\xda\x55\x8d\x26\x82\xba\xe5\x9b\xe7\x4f\x87\x48\xc3\xa5\x39\xfe\x6c\xf3\xeb\x5e\x7f\xa8\xeb\xf0\x86\xa0\xbf\xa8\x87\x5c\xd7\x11\xab\xa4\x9f\x5a\x30\xeb\xc3\x50\xbd\x7b\xa1\x45\xeb\x03\x9f\x84\xdb\x6a\x13\xf7\x8c\xf1\xb4\x17\x5d\xd0\x8d\x7e\xb4\x68\xbc\xca\x12\xf3\x10\x2b\x8d\x58\x46\xa6\xcd\x07\xad\x36\x80\x69\xd4\xd3\x82\xb6\xfa\xb1\x53\xb4\xaf\x81\xab\x68\xa2\x05\x34\xbe\xda\x86\xc9\x3a\xd8\x0b\x0d\x42\x5f\x21\xdd\x37\x41\xa1\x50\x03\x80\x85\x48\x6c\x04\x81\xa9\xa7\x6b\x60\x8d\x9f\xb6\x6c\xb1\xb1\x34\x6b\x70\x31\x2f\x64\x46\x3d\xbb\x5c\x93\x78\x04\xb4\xf6\x49\x55\x02\x0a\x2a\xba\x06\x89\xf8\x2c\x02\xec\xff\x30\x42\xfe\x42\x61\x11\x48\x6e\x67\xfa\x54\xa0\x2f\xb9\x31\xb7\xb5\x83\xaf\xdb\xf6\x60\x38\x99\xb6\x3b\x84\xd3\xc3\x84\x52\x7f\x18\x85\x84\xde\x02\x65\x0f\x0d\xdf\x0c\x67\xef\x58\xe4\xc1\x90\x76\xdf\x4f\xa1\x8b\x6a\xfd\x61\x13\xd4\x68\x68\x93\x9d\x83\xf4\x76\xcc\x81\x5a\xc4\xeb\xd1\x5e\xaf\x1b\xe2\xfa\xb3\x0e\xb7\xdf\x4b\xe0\xab\x17\x45\xf1\xd2\xa2\x85\xb5\xa0\x9d\xc0\xbb\x2a\x67\xef\xe2\x6f\xd1\x18\x7d\x87\xf6\xc6\x57\x2b\x09\x02\x11\xaa\x62\x65\xbc\x37\xb2\xc1\x41\xcb\x02\x49\xe0\x0b\x42\xb1\x64\xdc\x89\x8d\xd0\xd5\x60\x6d\x1e\x2d\x40\xd7\x40\x57\xe5\x6c\x82\xa2\x1c\xe4\x14\x17\x45\x62\xb1\x4c\x7d\xa7\xa1\x53\x78\xb4\x1d\x65\xd5\x52\x33\x9b\x4b\x56\xc4\x9a\x4c\x8d\x67\xec\x3f\xf3\x90\x0c\xd7\x0c\xe7\x90\xb2\x1b\xe0\xab\xcf\xd0\x08\xdc\x8a\xf6\xc8\x16\x68\x8a\xc1\xcf\xd0\x06\x05\x16\x32\x69\xaa\xcb\x07\x37\x45\xfb\xd8\xe6\x92\x99\x12\x34\x7c\x19\x5a\xcb\xdb\x9c\xdd\x30\xbb\x67\xab\x44\xee\x3d\xb6\x11\x20\xcf\xa1\x00\x2c\xdc\x1b\x0f\xdc\xbc\xeb\xdc\xa2\xe9\x02\x15\x17\x52\xed\xbe\xa3\x94\x2d\x96\x58\x92\x2b\x52\x10\xb9\x9a\xc6\x76\xbc\xbe\xca\xd2\xa6\x65\x6e\xb1\xf4\xe5\x25\x15\x52\x8a\x66\xc9\x6b\xc1\x6b\x26\x61\x70\xb6\x7b\xef\x45\xb2\x97\x70\x55\xe6\x88\x08\x84\xe9\x0a\xb1\x19\x8a\x6c\xcd\x2f\x26\x06\xa8\x11\x13\xb3\xe9\x16\x93\x82\xe5\x13\x79\x4b\xa3\x5d\x74\xb2\x9b\x1f\x54\x84\xbc\xa7\x1e\xae\xcb\xbf\x8a\x86\xa9\x05\x5d\x6c\x4e\xdf\x62\xca\x14\xf7\x57\x74\xc6\x1c\x7b\x5a\xa9\xdc\xdf\x9d\x68\x9d\x75\xff\x84\xa8\x01\x56\x65\xdb\x37\xac\xb0\xfd\xf9\x09\xce\x7e\x84\x62\x09\xee\xad\x11\xed\xf1\x37\x9c\x98\x13\xa0\xea\xa2\x7b\xb3\xff\xf6\xcf\xe4\x36\xdf\x50\x55\xa7\x77\x86\x4c\x08\xbd\xb1\xdb\x2a\xff\xfa\x47\x73\xf8\x1b\x04\x73\xea\xcb\x28\xf5\x8f\xc6\x2a\x42\xad\x8b\x31\x95\x5a\xbd\x07\x1a\xe7\x80\x33\x42\xf3\x5d\xa4\x6f\xee\xd8\x73\x8d\x6a\x94\xee\xe1\x5e\x5e\x0f\x2c\xb5\xba\x53\x2d\x76\x45\xa2\xd6\x58\x37\xeb\x00\xb6\x4d\x46\xde\xdd\xf6\x15\xff\x6b\x58\x8d\xc2\x94\x3d\x9c\x23\x74\x1f\xbc\x75\x9f\x5b\x47\xc1\x61\x0b\x1a\xd5\x5b\xc9\x3e\x0c\xa6\x46\x23\x4f\x7e\x39\x39\xbd\x4c\x7e\x7c\x71\xfa\xf2\xf5\xc9\x79\x8d\x52\x19\x50\xf2\x81\x7e\x5f\x6f\x7d\xee\x18\x5f\xf5\x53\x9f\x5f\x1d\x7c\xa5\x3e\x54\x3f\x6d\xad\x03\x74\xc2\x39\xe3\xcd\x21\xd3\x4e\xc0\xe0\x3f\x1b\x48\xce\x76\x6a\x61\x23\x42\x72\x8d\xeb\xc5\xfa\x97\xac\x5d\x13\xc2\x2d\x91\xf1\x37\xa3\xc3\x9d\x8f\xeb\x94\xbb\xdf\x6f\xfe\x1f\x59\xe3\x1a\xa0\xa8\xf5\x6d\x5d\x71\xb8\xaf\xe2\xf7\xf9\xef\x08\x3e\x95\xde\x15\x2c\xb3\x5e\xfb\xff\x06\x00\x00\xff\xff\xf5\xbd\xb4\x7f\x85\x42\x00\x00")

func includeWt_raiiHBytes() ([]byte, error) {
	return bindataRead(
		_includeWt_raiiH,
		"include/wt_raii.h",
	)
}

func includeWt_raiiH() (*asset, error) {
	bytes, err := includeWt_raiiHBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "include/wt_raii.h", size: 17029, mode: os.FileMode(436), modTime: time.Unix(1645816211, 0)}
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
	"include/wt_raii.h": includeWt_raiiH,
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
	"include": &bintree{nil, map[string]*bintree{
		"wt_raii.h": &bintree{includeWt_raiiH, map[string]*bintree{}},
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

