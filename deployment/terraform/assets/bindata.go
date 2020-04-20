// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/cluster.tf (9.072kB)
// assets/dashboard.yaml (228B)
// assets/dashboard_data.json (26.459kB)
// assets/datasource.yaml (296B)
// assets/outputs.tf (327B)
// assets/variables.tf (666B)

package assets

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

var _clusterTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x5f\x8f\xe3\xb6\x11\x7f\xf7\xa7\x98\xea\x0e\x48\x5b\x94\x92\xbd\xbb\xd9\xdb\x4d\xb0\x09\xd2\x7f\x40\x81\xa4\x09\xd0\x00\x7d\x38\x2c\x04\x4a\x1c\xcb\x84\x29\x92\x25\x29\xef\xfa\xae\xdb\xcf\x5e\x90\x92\x6c\x59\x96\x6c\x9f\xef\xf6\x2e\x39\x64\xef\xc5\x18\xce\x0c\x87\xbf\x99\xf9\x69\xc8\xd3\x46\xad\x38\x43\x03\x11\x7d\xb0\x11\xbc\x9d\x00\x68\xa3\xe6\x5c\x20\xdc\x41\x54\x96\x44\x28\xca\x1c\x5a\x17\x4d\x00\x0c\x16\x5c\x49\xf0\x2b\x95\x25\x48\xad\x23\x17\x5e\xbe\x42\x63\xfd\xc2\x1d\x44\xff\xfb\x06\x2e\xe2\xab\x57\xd1\xe4\x69\x32\x31\x68\x55\x65\x72\x0c\xbe\xd3\x25\xae\x53\x4d\xb9\x89\x20\x5a\xe2\xba\xde\xca\xcb\x24\x2d\x11\x82\xcf\x97\x6f\x57\xd4\xc4\xb9\xa8\xac\x43\x13\xe4\x4f\x64\x89\xeb\x60\xe4\xe3\xaa\x32\xc1\x73\xef\x07\xee\xc0\x47\xf8\x7b\xaf\x6e\xed\x22\xdd\xae\xfc\x61\x7f\x5f\x2e\xad\xa3\x32\xc7\x08\x22\xaa\x75\x6a\xd1\xac\xd0\xd4\xdb\x3b\x5a\x58\xb8\x0b\x3f\x01\xfe\xe9\xe3\x18\x89\x82\x6a\x4d\x5e\xbe\xcd\x55\x25\x5d\xcc\x25\xc3\xc7\x27\x1f\xd0\xd3\x64\x02\x90\x2b\x29\x31\x77\xfe\xf8\xb5\x9f\x17\xf0\xf3\x02\x81\xe1\x9c\x56\xc2\x41\x65\xd1\x84\x13\xce\x95\x01\x55\x19\xf8\xee\x87\x7f\x04\x35\xb7\xd6\x61\x3b\x6b\x17\x51\x10\x78\xcd\x80\x6c\x56\x49\x57\xd5\xb2\x85\xb2\x0e\xee\xc0\xa2\x98\xc7\xcd\x21\xb9\x6e\x77\xa6\x25\x87\xed\xdf\x1d\x44\xb4\xe4\x64\x3a\xcf\x2f\xa6\x8c\xcd\x18\xbd\x9a\x5e\xbf\xba\x99\x66\x11\xbc\x80\xd9\x4d\x3c\xbd\x82\xef\x7f\xfe\xd7\x04\xa0\x85\x23\x6d\x02\xf0\x87\xf5\xb8\xec\xc8\x77\x53\x13\xbc\x77\x53\x18\x2f\x71\x1d\x73\x16\x4e\x5f\x49\xd7\x89\x61\xcf\x5b\x50\xf0\x25\xa2\xf3\xd4\x62\x5e\x19\xee\xd6\x69\x61\x54\xa5\x53\xce\x3c\xf8\xaf\xc3\x41\xa3\x97\x6f\xfd\x06\xbb\x1a\xde\x53\xcc\xd9\x53\xf4\xa7\xc3\x3a\x69\xa1\xac\xe5\xb5\xaa\x07\xe6\x7e\x52\x57\xf1\x8a\xfb\xaa\xf4\xb5\xed\xab\x25\x6a\xf2\xd3\xd4\x46\x27\xe0\x92\x3a\x87\xa6\x54\xd6\xa5\x82\xe7\x28\x2d\xa6\xde\x20\x68\x33\xb4\x8e\x4b\xea\x9a\xf2\x4e\x16\xaa\xc4\xa4\x4e\x51\xb2\xb5\xeb\xb8\x20\x8d\x8b\x4d\x81\xec\x04\x62\xb0\x54\x0e\x09\x3e\x62\xde\xc6\xc3\xa5\xe0\x12\x37\x48\x00\x44\x0f\x0b\xdf\x7e\xaf\xe1\x77\x40\xe6\x90\xac\xa8\x49\x04\xcf\x92\x5c\xa8\x8a\x25\x2d\xb0\x49\xa6\x94\x23\x73\x2e\xb9\x5d\x20\x83\xfb\xaf\x81\x29\xc0\x7c\xa1\xe0\x8b\x7f\x53\xee\xb8\x2c\x42\xc5\x05\x23\xc2\x25\x77\x71\x1c\x7f\xf1\x35\x58\x81\xa8\x61\xe6\xb5\x25\x36\xb8\xfa\x1d\x0b\x74\x40\x88\x54\x24\x5f\x60\xbe\x24\x39\x1a\xc7\xe7\x3c\xa7\x0e\x81\xfc\xe7\x47\x20\xb0\x70\x4e\xdb\xaf\x92\xc4\x5e\x12\xac\xc8\x03\x5a\x47\x66\x31\x2d\xe9\x1b\x25\xe9\x83\x8d\x73\x55\x26\x0c\xb3\xd8\xa8\xac\xb2\x4e\xa3\xc9\x51\x7b\xcc\x62\xae\x92\xab\xd9\xdf\xfe\xfe\x97\xdb\xdb\xbf\xc6\x85\x2e\xe0\xbf\x60\x2b\xa6\x80\x6a\xe7\x5b\x1b\x28\x63\x40\xb6\x71\x6c\xd6\x42\x3c\x6b\xa8\x34\xa3\x0e\x47\xd6\x03\x14\x42\x78\x3d\x6d\x54\x89\x6e\x81\x95\x25\x52\x31\x8f\xaf\x56\xc6\xa1\xe9\x9f\xf0\x47\xe8\x24\x8a\x71\xeb\x62\x47\x4d\x5c\xbc\x81\xba\xe9\x3b\x85\xc0\xd4\x83\xf4\xd4\x97\x56\x46\x3c\x6d\xdd\x38\x6a\xe0\xf1\xcd\x7c\xc4\x4d\x2f\xd0\x72\xd5\xd1\x83\x44\x69\x97\xd4\x6d\x7d\x1f\x6a\xe3\x10\x53\x95\xe8\x0c\xcf\xed\x79\x6c\xd5\x18\x7f\x36\x14\x15\xb9\x8b\x58\x50\x53\x84\x9e\x3a\x85\x97\xce\xe5\x9b\x06\xb8\x0d\xe7\xec\x13\xc9\x6f\xfd\xfb\xcc\xfd\xdb\xd3\xb5\x6b\xeb\xb0\xcc\x9d\x00\x94\x34\x13\x38\xae\x39\xe0\x95\x32\x16\x2a\x56\xf0\x6c\xae\xa4\xcb\x95\x9c\xf3\x62\xd6\x43\xad\x05\x86\x89\xb8\x30\x74\x4e\x25\x0d\x60\x28\x6b\x13\x83\x02\xa9\xc5\xa4\x91\xa7\xd7\xf1\x75\x7c\x91\xd2\x92\x5d\x5f\xc5\x0c\xb3\x5e\x00\x4c\x2f\x0b\x20\x1c\x4e\xd3\xde\x1e\x8c\x51\x2c\x95\x24\x06\x3d\xdd\x1c\x3b\x7e\xe3\x9c\x34\xbc\xd0\xd3\x46\xb3\xe2\x79\x5f\x09\xac\xa3\xc6\x9d\x4a\x3c\xda\xa8\xc7\xf5\x79\xb4\x13\x4c\x6b\xd2\xe9\x77\xfe\xee\xdf\x28\x0f\xec\xf5\x7e\xdf\xce\x6f\x5a\x47\xd8\x1f\x57\x76\x07\x91\x21\xbb\xfd\xb1\x04\xbe\x81\x19\x7c\x0b\x33\xf8\x0a\xa6\x3e\x64\x6b\x55\xce\xa9\xc3\x74\xc3\x63\x29\x65\xcc\xa0\xf5\xe7\x77\xa6\xc2\x73\x79\x25\x84\x5c\x8f\x27\x75\x0e\x36\x24\x36\xca\x5f\x9f\x8e\xb3\x7f\x0d\x7c\x77\x2e\xcf\xc8\x82\xcb\xc7\x0f\xd3\x89\x43\xae\x4c\x19\x10\x40\x97\x27\x61\x39\xb1\xdc\xa1\x25\xb5\x05\x4b\x9a\xe4\xf5\x8c\x84\x04\x32\xb7\xfb\x56\x74\x45\xb9\xf0\x86\x49\x77\x88\x18\xf3\xbd\xd5\xe9\x75\x7a\xaf\xd5\x0d\xb3\x69\xdb\xb6\x9d\xb6\xef\x8b\x9a\xab\xe0\x81\xae\xaa\x9b\x8a\x65\xfb\xa3\x3e\x67\x28\xfd\xd7\x07\xcd\x50\xe7\x0f\x31\x07\xcb\x06\x6e\x57\x9b\x90\xf6\xdd\xd5\x5d\xd3\x39\x8a\x8f\xa3\xfd\x19\xee\x25\xdb\xa0\x04\xb5\xf6\x58\xe0\x5e\x67\x02\x80\x1e\xd8\x3e\xe9\x8c\x19\xd5\xca\x9e\x36\xb4\x16\xeb\x94\x97\x25\x32\x4f\x1e\x62\xdd\xb5\x6a\x58\x83\x56\x4e\xa5\x25\x97\xca\xa4\xcd\x5d\x39\xad\x74\x61\x28\xf3\x7d\x34\xa7\xc2\xe2\x3e\x29\x77\x8e\x17\x41\xb4\x3d\x60\x93\x99\x01\x74\x46\xd1\xf5\x70\x32\xea\x68\x46\x2d\x76\x86\xa7\x11\x83\x5a\xbf\xa4\x41\xb2\x21\x9a\x1d\x14\x5a\xe9\x56\x4f\x53\x6b\x1f\x94\x61\xbb\x7a\xad\x74\x02\x60\x97\x5c\xa7\x73\x2e\xa9\x48\xad\xa4\xda\x2e\x94\xeb\xe0\x33\x80\xe1\x66\x71\x28\x2d\x07\xf2\x51\xff\x68\x71\xee\xa9\xef\x2e\xbe\x1e\x76\x72\x7f\x78\x80\x1c\xe6\x78\x96\x05\x82\xbf\x3f\xf8\x75\x6d\x1f\x53\x52\x5a\xa0\x74\xef\xfa\x08\xe1\x6d\x4e\x7e\x86\xf8\x68\xc3\xfb\xd8\x9b\x82\x8f\xf6\x43\xbe\x2a\xec\xfa\xab\xc9\xe6\x8c\x3c\x05\x3f\x4d\xaa\x3e\x93\xaf\x5d\x7b\xb7\x75\xa5\xde\xbd\xcf\x86\x0b\x6c\xa8\xb7\x23\xd7\xd9\xad\xe5\x76\x6d\xe7\xf6\x1a\x9e\x01\x89\x77\x45\x64\xf1\xc7\xb1\x85\xad\xb1\x29\xbb\x3e\x0f\xcd\x9e\xbb\xe9\xa9\x1f\xe9\x6a\xf0\x3b\x44\x75\xe8\x5d\xae\x75\x40\x6a\x07\x93\xf0\x66\x93\x1b\xae\xdb\x37\x9b\xef\xb4\x86\x56\x09\x82\x52\x00\xbf\xed\xc5\x96\x4a\x61\x60\x83\xf0\x9c\xc4\x65\x11\x66\xc0\xba\x20\xe6\x46\x95\xa9\x56\xc6\x85\xa8\x2e\x2e\xea\x66\x53\xad\xa8\x23\xd4\x46\x39\x95\x2b\xd1\xc4\xef\x72\x5d\x03\x91\x73\x66\xd2\x4c\xa8\x7c\x59\x97\xe9\x34\x0e\xff\x92\x69\x74\xdf\x0c\xcf\x87\x76\xbc\x99\x5e\x7f\x39\xb0\xe7\x46\xfc\xe1\x77\x0d\xce\x5f\xf5\xf6\xec\x08\xb7\x3b\x76\xf7\x7b\x01\x3f\xd0\x75\x86\x60\xd0\xfa\x4b\xb5\x03\x25\xc5\x3a\x78\x85\x9f\x36\x77\x38\x68\xae\x28\xdf\x36\x26\x7f\xae\x1c\x2c\xa8\x64\x6b\xa8\xbb\xcc\xd1\xa5\xef\x95\xe6\x45\xda\xc2\x03\x77\x0b\x55\x39\x28\xa9\xac\xa8\x10\x6b\xb0\x76\x41\xbc\x06\x97\x4e\x81\x5b\x60\xe3\x30\x7e\x6f\xa0\x6b\xf8\x6e\x67\xd3\xe9\x1e\xd8\xbd\xa5\x2e\xe0\x7d\xd0\x77\x8b\xfb\x00\x2b\x75\x5f\x1e\xda\xe0\x70\xbc\x08\xf6\x83\x6a\x65\xbd\xf4\x93\xd9\x69\xd9\x3f\xa9\x2d\x9b\x77\xd6\xf7\xe8\x4e\xd2\x78\x38\xb9\x49\x6b\xfd\x13\x7b\xf5\x78\x3e\x6f\xa6\xaf\xae\x46\xf2\xb9\x59\x1a\xc8\x67\xc5\xde\x35\x9f\xed\xcb\xf5\xa9\x85\x76\x66\x60\xef\x5e\x68\x67\x04\xb6\x4f\x37\xbd\xa5\x4f\x85\xd8\x39\x81\xbd\x3f\x62\x1f\xb9\x2d\x8f\xf5\x25\xcb\x3a\xfd\x38\x7e\xbf\xea\x7f\x25\x4f\x40\xf8\xf2\x72\x7a\x3d\x82\xf0\x66\xe9\x19\x10\x3e\x21\xb2\x2f\xaf\x2e\xf7\xbf\xbb\xbd\xa5\x67\x88\xec\x04\x92\xdc\x8e\xf5\xa7\xf0\x63\x18\xe8\x8f\xce\x2f\xdf\xb7\xe4\x17\xf4\x7f\x55\xa3\xcc\xc1\xa9\xe2\x6a\xba\xf7\x71\xed\x08\x87\xa7\x0a\x7f\x53\x69\x31\x6d\xae\x87\x61\x97\x5f\xd6\xb7\xb2\xfd\x0f\x98\xa3\x8d\xd9\x28\xbe\x5b\x77\x7e\x8a\xa9\xf3\x76\x7a\x3b\x84\xe3\x46\xfc\x3c\xbb\x5e\xee\x57\xc8\x8e\xf8\xfc\x5d\x7f\x59\xf5\x52\xbf\x9c\x9f\x4a\x1b\x41\xfb\x38\x6d\xfc\xe4\xd5\x9e\x8d\x2d\x6e\x86\x60\xba\x79\xbf\xb4\x7c\x1a\x82\xfa\xb8\xa5\xf0\xff\x00\x00\x00\xff\xff\x7d\x1d\x9a\xe0\x70\x23\x00\x00")

func clusterTfBytes() ([]byte, error) {
	return bindataRead(
		_clusterTf,
		"cluster.tf",
	)
}

func clusterTf() (*asset, error) {
	bytes, err := clusterTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x47, 0xeb, 0x4, 0x3c, 0x1a, 0xa7, 0xd7, 0x12, 0x29, 0xc, 0xc5, 0xe0, 0x53, 0x33, 0xd2, 0x66, 0xb, 0x8c, 0x29, 0xa8, 0x2d, 0xc1, 0x69, 0x21, 0x81, 0xd7, 0x1, 0x83, 0x32, 0x67, 0x49, 0xab}}
	return a, nil
}

var _dashboardYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8c\xb1\x4e\x03\x41\x0c\x44\xfb\xfd\x8a\xe9\x52\xa1\x0b\xa2\xdb\x9a\x26\x35\x0a\xbd\x0f\xfb\x12\x4b\xce\x7a\x65\x6f\x0e\xf1\xf7\x68\xaf\x40\xa2\x9d\xf7\xe6\x51\xd7\x4f\x89\x54\x6f\x15\xaf\xa5\xf4\xf0\x5d\x59\x22\x6b\x79\x41\xa3\x87\x54\x9c\x92\x1e\xdd\xe4\x54\x00\x8f\xdb\x85\xa7\x07\x6c\x6e\x2c\xf1\x8f\x8e\x9f\x2e\x15\x9b\x9a\x14\x80\x35\x69\x35\x79\x17\x93\x71\xc4\x37\xb2\x9c\x40\x58\xc7\x24\x15\x23\x9e\x73\x78\x76\xa6\x21\x97\x36\x24\x76\xb2\x0f\xf9\xf2\xc6\x59\xf1\x76\x3e\x17\x80\xcc\xfc\xfb\xaa\xd7\xc3\xc9\xbf\x8f\xf7\x19\xcd\x5a\x00\xa0\xd3\xb8\x57\x2c\x3b\xc5\x62\xba\x2e\xb7\xa0\x8d\x1a\x2d\x4c\x79\x5f\x9d\x82\xf3\x37\x00\x00\xff\xff\x94\x42\xcd\x38\xe4\x00\x00\x00")

func dashboardYamlBytes() ([]byte, error) {
	return bindataRead(
		_dashboardYaml,
		"dashboard.yaml",
	)
}

func dashboardYaml() (*asset, error) {
	bytes, err := dashboardYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dashboard.yaml", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x3e, 0x13, 0x14, 0xd7, 0xd7, 0xbd, 0xed, 0xe7, 0xca, 0xfc, 0xe7, 0x3a, 0x3a, 0x0, 0xc1, 0x74, 0x2d, 0xef, 0xf0, 0x9d, 0x47, 0xe0, 0x99, 0xba, 0x23, 0x47, 0x12, 0x37, 0x12, 0x52, 0xd9}}
	return a, nil
}

var _dashboard_dataJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x5d\x4f\x1b\x3d\xf6\xbf\xef\xa7\xb0\xac\xfe\xff\x0a\x15\xf0\x64\x02\x69\x29\xd2\x6a\x45\x5b\xb5\x5b\xa9\xec\x76\x4b\xfb\xec\x45\x8b\x46\xce\xcc\xc9\xc4\xc2\x63\x4f\x6d\x0f\x90\x07\xb1\x9f\x7d\x35\x9e\x37\xcf\x0b\x4d\x88\x48\x13\x5a\x73\x01\x99\x63\x8f\x7d\x7c\x5e\x7f\x73\xc6\x31\x37\x4f\x10\xc2\x84\x73\xa1\x89\xa6\x82\x2b\x7c\x8c\x32\x12\x42\x98\x51\xa5\xf1\x31\xfa\x6a\xae\x50\x41\x35\x2d\x93\x94\x32\xfd\x9e\xe3\x63\xe4\xed\xd6\xd4\x90\x68\xa2\x44\x2a\x03\xc0\xc7\x08\xef\xed\xa1\x77\x92\x4c\x09\x27\x68\x6f\x0f\x5b\xdd\x80\x93\x09\xcb\xba\x68\x99\x82\x45\x9f\xd1\xb0\x87\x4a\x03\xc1\x5f\x0b\x26\x64\x36\xa6\x8c\x26\x64\x30\xdc\x45\x23\xcf\xdb\x45\xa3\xf1\x78\x17\x79\x3b\xf6\xd0\x9c\xc4\x66\xee\x93\x7a\x39\xe8\xff\xd1\x09\x03\xa9\x95\xdd\x4f\xcf\x13\xd3\x2f\x24\x6a\x36\x11\x44\x86\xb8\x68\xbb\x35\x7f\xcf\x9f\x20\x74\x9b\x75\xc7\x10\x52\xdd\xe2\x16\x47\x1c\xf4\xfb\x10\x1f\xa3\xd1\xf8\x70\x94\x53\x24\x49\x66\x9f\x85\x60\x9a\x26\xa5\x4c\x30\x35\x5d\xf2\x8f\x1a\xa4\xe1\x26\x6b\x1c\x1f\x3d\x3f\x1c\x0d\x87\xa3\xe1\xd1\xd1\x81\x69\x65\x94\x5f\x64\x52\xff\x7a\x6e\x2e\x13\xc2\x81\xa9\x4a\xee\xa5\xd4\x31\x61\x94\x28\x23\x09\xa3\xa2\xdb\x72\x39\x78\x42\x0c\x65\x4a\x98\xaa\x04\x67\x56\xf6\x01\x78\xa4\x67\xd9\x9c\xc3\x06\x1d\xfa\xba\x5b\xaa\xe3\x29\x63\x55\xcb\x94\x32\x66\xeb\xd9\x10\xde\x49\x12\x52\xe0\x99\x75\xd4\x43\x47\x92\x86\x1f\x45\x6d\x3f\xb9\x52\xf1\x31\x3a\xb2\x24\x7f\x95\x8d\x35\xb2\x08\xd7\xf6\x18\x08\xe1\x79\x76\x5d\x6a\xa3\x1a\x7b\x46\xc3\x10\xf8\x19\x48\xda\xc3\xbc\x91\xb4\xf7\xb2\xba\x66\x10\x01\x0f\x9b\x7c\x90\xcb\xa8\x7d\x1f\x42\x38\x48\xa5\xcc\x97\xd1\x6e\x89\xc9\x75\x1f\x95\xf2\x1e\xaa\x9a\x89\xab\xae\xe1\x6a\xa1\x09\xeb\xe9\x7d\x49\x58\x5a\x2f\xa2\xb3\x52\x46\xb9\x69\xb5\x47\x33\xc4\x2b\x1a\xe6\xca\xac\xa8\x99\x9e\x3e\x0a\xca\xf5\xa9\x30\x9e\x63\x08\x95\x99\x63\x91\x34\xfd\xb9\x52\xf4\x87\xca\xe0\x3a\xb3\x27\x20\x03\xe0\x9a\x44\xd0\x91\x72\x92\xcd\x94\xe9\x3d\x55\xa5\x61\xd7\xf4\xae\x52\x24\xf0\x10\x24\x18\xb7\x9d\x32\xa1\x6b\xbe\x94\xd1\xe2\xbf\x2e\x41\x4a\x1a\x42\x6d\xf9\x79\x63\x42\x02\xe8\x33\x5c\xa5\x49\x70\xd1\x99\x45\x69\x48\x12\x08\x3f\x50\xde\x65\x58\x13\x19\x81\x56\x56\x04\xb3\x63\x58\xe6\xdc\xd7\x89\x61\x4f\xa5\xf1\x80\xf2\x40\x02\x51\x30\x88\x89\xd6\x20\x63\xa1\xb4\x1f\x4e\x7c\xa5\x85\x04\x5f\xd3\x18\xfc\x40\xa4\x5c\xdf\xc4\xa0\x67\x22\xfc\xdb\x7f\xbf\xe1\xa7\x5a\x24\x59\x17\x43\xff\x86\x6f\xbf\x8e\xe3\xf3\x9d\x1d\x34\x99\xa3\x41\xde\xc9\x8e\x4c\x95\x55\xbe\x15\x32\x26\x99\xc1\xe1\x9b\x62\xac\xdb\xdb\x66\x3f\x09\x53\x13\x5d\xf0\x09\xae\xc8\xb7\xc5\xa7\x5a\x4e\x7a\x26\x41\xcd\x04\x0b\x5b\xf2\xcb\x58\x7d\x2b\x45\xdc\x76\xe3\x8c\xfe\x09\xa2\xc2\x20\x5a\x37\x9c\xcd\xe8\x54\x77\xef\xd0\x26\xee\xe1\xcf\x22\x41\xde\x10\xbd\x79\x85\xbe\xa7\x46\x75\xd9\x1a\xcd\xaa\x6b\x9d\xea\x2a\xf8\xdd\xd8\x7e\x41\x24\x84\x5d\xcf\x50\x42\xea\x96\xd7\x1b\xa7\xf0\xcb\xb0\x4c\x79\x48\x2f\x69\x98\x12\x86\x3b\x16\x5a\xf6\x31\x31\xb7\x66\xe0\x9a\x5c\xd3\x96\xa9\x4f\xd2\xe0\x22\xd7\xbf\xbd\xae\xcc\x8b\x0b\x77\xc9\x96\xde\x93\x3d\x5a\xbd\xfb\xbd\xbb\xf2\xe2\x1e\x27\x9a\x93\x6b\xf8\x81\xd9\x4d\x2b\x13\x50\xb3\x4c\x12\x4d\x2b\x21\x13\x60\x1d\x26\xb2\x06\x11\xbd\x22\x0a\x9a\x29\xb7\x0a\x54\x9d\xee\x79\xa4\xea\x90\xad\xc5\xd4\xc6\xb5\xfb\x48\xf8\xec\x38\xc1\xbc\xab\x74\xc2\x68\xd4\x17\xa3\x0d\xfd\x03\x5c\x56\x4c\x37\x12\x7e\x21\x82\x5f\x20\xd7\xbe\x5c\x94\x6b\x1b\x84\x95\x93\xed\x0b\x97\x6c\x5d\xb2\x7d\xd0\x64\x4b\x12\x6a\xe7\xd9\x19\xe1\x21\x03\x59\x25\xda\xac\xb9\x27\xd3\x16\xdd\x5a\xa9\x96\x72\x0d\xf2\x92\xb0\xb7\x24\xd0\xe6\xc1\xe1\x60\x41\x26\x2e\x86\xd9\xfa\x54\xac\xf3\x54\x7c\xf2\xf1\xfd\x03\xe7\xe2\xd1\x6f\x9f\x8b\x63\xb5\xf9\x04\xe7\x12\xf1\xe6\x12\x31\x66\x82\x84\x1a\x94\xde\x2b\x48\x55\xb7\x9e\x1a\x40\x4e\x97\xd2\x04\x97\xe6\xa0\x9b\x78\x5e\x3e\x5a\x21\x85\xff\x2e\x19\x7c\x64\x53\x2f\x5a\x11\xb7\x93\xd4\x03\xc1\x39\x04\x1a\xc2\x9f\x94\xd9\xc7\xbf\x50\x66\xef\xcb\xea\x49\xf6\xcb\xa8\xf6\x86\x72\xa5\x09\x0f\xc0\xa4\x74\x05\xf2\x12\x64\x96\xcb\xbd\xf8\xbc\x95\xbd\xeb\x70\x67\xf0\x40\xbe\xce\x05\x09\xde\x5b\x90\xe0\xcb\xc9\xdb\x19\x3e\x06\x2d\x69\x60\xe2\x7f\x1f\xd3\x77\xc2\x81\x66\x8c\xd4\x90\x65\xdb\xf1\xc2\x40\xbe\x08\x03\x2d\x27\xad\xb5\x8b\xeb\x73\x77\xe9\xab\xc9\xe9\xd5\x22\x39\x6d\x12\x4b\xfd\x33\x8d\x27\x20\x91\x98\xa2\x53\x50\x8a\x44\xa0\x50\x02\x12\x9d\x41\x20\x78\xb8\x00\x4d\xc5\xea\x13\x28\xc1\xd2\xa2\xa6\xdb\x8d\x70\xae\xf0\x91\xff\xdc\x1f\xc7\xe0\xd7\x0d\x34\x8b\x1c\xe4\x42\x0e\x72\x6d\x1e\x72\x75\xca\x26\x2f\x57\xc0\x5c\x96\x7c\x1c\xe8\x72\xa0\x6b\x5d\xa0\x6b\xa6\x75\xe2\x4b\xf8\x9e\x82\xd2\xea\xd1\xa1\x2f\xc3\xbd\x71\x74\xb5\x11\x10\x76\x4f\xe9\x6d\x0d\x1a\x5b\x56\x6e\x5b\x0d\xca\x3e\x15\x72\x77\x50\xcc\x41\xb1\xfb\xb0\xea\xa0\xd8\xef\x03\xc5\xda\xd5\x2f\xef\xf9\x0a\x50\xec\xb9\x43\x62\x0e\x89\x3d\x28\x12\x6b\xa7\xe2\x2b\x98\x28\x61\xa2\xff\x0f\x41\xc4\x36\xa2\xaf\x36\xeb\x6b\x44\x60\x2b\x4a\x6d\x13\xa8\xeb\x71\x20\xa8\xba\xac\xf5\xba\x74\x66\xf4\x06\x2e\x69\x00\x0a\x0d\xfe\x03\x93\x33\x23\xdd\xb2\x31\x9b\x61\xc7\x21\x2c\x87\xb0\x1c\xc2\x72\x08\x0b\xa1\xbb\x8a\x5d\xde\x8b\x15\x20\x96\xe7\x20\x96\x83\x58\xeb\x82\x58\xe1\xc4\x8f\x89\xd2\x20\xfd\xa0\xce\x64\x8f\x09\x68\xfd\x60\x01\x3f\x07\x6e\xdd\x5f\x82\x5b\x53\xea\xba\xbf\xec\x1e\x17\x64\x33\xdf\x5b\xd1\x02\x9d\x9a\x45\xa2\x37\x44\x93\x49\x96\x6a\x1d\x4e\x73\x38\xcd\xe1\xb4\xdf\x06\xa7\x2d\xdc\xcb\xdd\xae\x84\x8d\x0e\x57\x81\x69\xbf\x0b\x4e\xf3\x1c\x4e\xdb\x92\x97\x92\x7e\x98\xe6\xdf\x87\xf4\x95\x79\xc5\xa3\x7c\x95\xc6\x77\xbd\x64\x1b\xc7\xe7\x3b\xe8\x0f\x84\xd0\xfd\x47\xcd\xb7\x92\xff\x60\xdc\x2d\xc3\x84\x0b\x65\xb4\x2a\x32\xdc\x24\xba\x39\x05\xc2\xcd\x8e\xf5\xe2\xdd\x1e\xfa\x6c\x27\x7a\x07\x65\xd6\x02\x65\xee\x80\x31\x0e\xc1\x38\x04\xb3\xdd\x95\xa6\xd1\x78\x05\x08\x33\x76\x08\xc6\x21\x98\x07\x45\x30\x91\xf0\x63\x88\x95\x26\x5a\xf9\x84\x31\x11\xf8\x93\xb9\x06\x75\x07\x96\xf8\xc3\x1b\x8e\x0e\xcd\xaf\x75\x03\x8a\x93\x8c\x97\xbd\x65\x60\xc5\x1d\x2b\x58\x6b\x85\xc9\x9e\xd3\x28\xc7\xa7\x3c\x55\xb0\x2e\xd9\x15\x07\x86\xb4\xbc\xec\xfe\x42\x3d\xcb\x58\x5d\xaf\x64\x17\xd6\x9f\xee\x21\xd9\x19\x90\xe4\x71\x08\xf6\x1f\x40\x92\xf5\xca\xf5\xf5\x76\x23\xdf\x58\xc8\x39\xfa\xa2\x48\x04\x68\x40\x39\x3a\x7d\xe5\xde\xb5\x6e\xae\x86\x77\xfa\xca\xc1\x5f\x07\x7f\xb7\x0a\xfe\xb6\x0b\x78\x07\x07\x2b\xa0\x5f\xef\xc0\xc1\x5f\x07\x7f\x1f\x14\xfe\xf6\x7e\x39\x51\x8a\x00\x94\xf2\x83\x24\xad\x4a\x52\x0b\x77\xc7\x3f\x43\xde\x70\xb8\xb9\x2a\xdb\xe3\xa8\x90\xbd\xfe\xf8\x05\x7d\xd1\x94\xd1\xbf\x4c\xc1\x0f\x7d\x22\x1a\xd0\xe0\xff\x1c\x56\xd8\x1c\x56\xf8\xa1\x3a\x90\x43\x0f\xc8\xa1\x87\xcd\xa3\x87\x4e\xf1\x6c\x25\xf8\x70\xe8\xd0\x83\x43\x0f\x0f\x8a\x1e\x22\xe1\x47\x42\x8a\x54\x67\x2a\xd9\xe2\xdd\x58\x0d\x3e\xd7\xb8\xe9\x6a\x19\x79\x6c\x76\x6f\xd5\x32\x92\x78\x24\x5b\xa8\xde\x09\xf4\xa9\xbd\x10\x87\x9f\xdc\x7e\xa9\x5f\x1c\x30\x3d\x29\x86\xcd\xfc\x35\xf3\xba\x6c\xd5\xde\x30\x77\x01\xac\x82\x19\xc4\xe4\x4f\x90\x2a\x37\xf5\x51\x7e\x2c\xb5\xd2\x73\x56\x1c\x83\x2d\x2f\xf2\x9e\x9a\x44\xb5\xee\xad\xad\x09\x85\xec\x2c\xca\x5e\x02\x72\x8a\xab\x69\x35\xc4\x09\x23\x9a\xf2\x68\x99\xf3\xc3\x09\x63\x7f\x66\x26\xd8\xb5\xd7\x1a\x59\x34\x54\xa9\xe1\xda\x28\xd2\x7b\x31\xda\x3f\xf0\xf6\xbd\x83\x7d\x6f\x38\x3e\x3e\x1a\x3e\x7f\xd1\xd4\xea\x65\x31\xea\x57\x8b\x88\x7a\x6f\xb3\x3a\x9c\xf7\x59\xd2\x92\xb0\x30\xeb\x08\x53\xca\x69\x11\x45\x72\xbb\xf2\x73\x07\x1b\x94\xb1\xbe\x71\x48\x79\x51\x5d\xb7\x03\x08\xe5\x01\x4b\x43\x38\x61\x7d\xc0\xa7\x72\xaa\x3c\x59\xd8\x43\xc5\x29\xd3\xb4\xeb\xde\xe5\x31\xe8\xdd\x1b\x6a\x0c\x53\x87\x57\x84\xf0\xf7\x14\xe4\x7c\x39\xee\x6b\xfb\xf2\x1a\xd4\x08\xae\x5b\x3b\x0c\xb0\xba\xa0\xc9\x17\xc9\xce\xe6\x3c\xe8\x8b\xaf\xdd\x38\xaa\x49\x64\xcc\x42\xfd\xbb\xe0\xa7\x65\x1c\xa5\x75\x9e\xb7\x68\x77\x75\x2f\x82\x6d\xbe\x3a\xab\x21\x55\xf0\x39\x1f\xaa\x17\x4b\x3e\x84\xa1\x9e\x30\xb6\x94\x65\x3e\xf5\x7d\xc2\xd8\xda\xac\xd1\x2c\xdd\x97\xa0\x52\xa6\x07\x5a\x24\x17\x03\x6f\xb8\x8b\x96\x3b\x83\xf3\xeb\xd3\x1b\xdf\x97\x84\x47\xe0\xab\x5b\xd5\x3e\x71\x73\xe7\x5e\x36\xdd\x32\xd0\xfe\xe0\xbb\xc8\x9c\x1b\xe7\x80\xde\x61\xd5\x15\xb1\xd7\xbc\x7f\x96\x38\x16\x38\xc9\xfe\xb3\xf2\x78\xd3\x6f\x78\xb0\xff\xec\xef\x3b\xdf\xf0\xfe\xb3\x87\xf5\x9c\x86\x23\x2e\xf2\x9b\x66\xe7\x0d\x78\x8d\x02\x96\x3f\x6c\xf5\xbd\x74\x5c\xe8\x52\x1d\x27\xfa\xe9\x8e\xd3\x73\x52\x7c\xf7\x38\xf8\x0d\x79\x4c\x79\x44\xfd\x4f\x70\x98\x15\xc4\xb0\xd0\x53\x8a\x03\xf7\x7f\x01\x47\x31\x7f\xeb\x7f\x2d\x62\xf0\x7b\x85\x95\xa6\xf9\xf3\x10\xe6\xe2\x6a\xcf\x1b\x97\x1b\x3d\xb1\x16\x05\x11\x37\xee\x4b\x68\x70\x61\x4a\x01\xc5\xdd\x85\x14\xfd\xf2\x99\xd0\xd6\x2c\x1e\x5b\x0f\x42\x25\x24\x34\x17\x07\xf6\x85\x57\x6f\x2e\xc5\x63\xeb\xb3\x67\x5f\x1c\x0c\xed\x16\xeb\x01\x66\x64\x7d\xf6\x8a\xff\xa6\x72\x5e\xae\x21\x33\x86\xae\xc5\xdd\x3d\x8b\x3d\xf0\x73\x7b\x60\x7b\x96\xd1\xa1\x7d\x61\x95\x69\x5e\x84\x36\xbf\x25\x2f\x0d\xf1\xfd\x25\x4c\x79\x03\x4f\xa4\xb8\x52\x05\x40\xb2\x5e\xdb\x57\x06\x8d\x3e\x82\x34\x0f\x11\x3c\x00\x74\x2a\x38\xd5\x42\x66\x08\xd7\xf4\x4f\x4d\xf5\x0c\x0f\xf3\x1f\xcf\xcb\xa9\x97\x15\xc8\x3e\x7c\x72\xfb\xbf\x00\x00\x00\xff\xff\x21\xe8\x07\xc8\x5b\x67\x00\x00")

func dashboard_dataJsonBytes() ([]byte, error) {
	return bindataRead(
		_dashboard_dataJson,
		"dashboard_data.json",
	)
}

func dashboard_dataJson() (*asset, error) {
	bytes, err := dashboard_dataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dashboard_data.json", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0xa7, 0x0, 0x3e, 0x97, 0x29, 0xa0, 0x50, 0xec, 0x6e, 0x8e, 0xca, 0xc6, 0xe3, 0xeb, 0xc8, 0xb6, 0x3f, 0x2f, 0x43, 0x13, 0x79, 0x42, 0x28, 0xb7, 0x15, 0xa2, 0x88, 0x46, 0x4d, 0xf7, 0x2e}}
	return a, nil
}

var _datasourceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\xc5\x45\x8c\x87\x37\x05\x45\x87\x2e\x9a\xb3\xe4\x07\xba\x33\xd6\xb5\xcd\x42\x96\x0c\x91\x72\x9b\xbf\x2f\x9c\x00\x6d\x37\x82\x87\x38\x38\x1c\x30\xd6\x32\xe9\x8c\x49\x33\xb1\xb3\x99\xd6\x12\x64\xd3\xf7\xe7\x18\xf1\x1a\xc2\x80\xac\xe6\xa8\x13\x92\xb8\x58\xed\x6d\xa4\xc1\x2b\xb4\x18\x9b\xbf\xf4\x2d\x89\x13\x89\x1b\x4b\xd2\x32\x87\x01\x9f\x8b\xf8\x7f\x83\xec\xa2\x59\x6e\x99\xd0\x02\x5f\xf8\x10\xdc\xc4\x18\xfe\x98\x62\x38\xa3\xc8\xca\x88\x5c\x25\x39\xcd\xcf\x4f\x12\x00\xbf\x6f\x8c\xd8\x5a\x5d\xe9\x0b\xbb\x05\x40\xc6\x91\x66\x8f\xe5\xd7\x3d\x00\xb5\xcd\xd7\x74\x74\x02\xbd\xe5\x88\x7f\xc7\x91\xda\x85\x93\xf4\xec\x11\xde\xfa\x61\xda\x7f\x1f\x02\x98\xd4\x8f\xac\x1f\xfa\x61\xb5\x5c\xc4\x25\x06\x00\x70\x5d\x79\x2d\xce\xb6\x4b\x8e\x38\xbd\xd9\xe9\x3b\x00\x00\xff\xff\x50\x28\x34\xbf\x28\x01\x00\x00")

func datasourceYamlBytes() ([]byte, error) {
	return bindataRead(
		_datasourceYaml,
		"datasource.yaml",
	)
}

func datasourceYaml() (*asset, error) {
	bytes, err := datasourceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "datasource.yaml", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0x82, 0x88, 0xc0, 0xa4, 0xae, 0x3, 0xd5, 0x18, 0x5f, 0xf7, 0x9a, 0xce, 0x18, 0xcf, 0x8c, 0xeb, 0x4c, 0x24, 0x3f, 0xba, 0x8b, 0xe6, 0x78, 0x8f, 0xb9, 0x5e, 0x3f, 0xa9, 0x91, 0xc7, 0x9d}}
	return a, nil
}

var _outputsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x10\x5c\xb9\xe8\x0d\x5c\x79\x04\x0f\x10\xa6\xcd\x20\x85\x98\x84\x99\x49\x55\x4a\xef\x2e\xc4\x06\xac\x48\xdd\x05\xfe\xfb\x2f\x7f\x52\xd1\x5c\x14\xec\x18\x45\x31\x0e\x24\x16\x66\x03\x30\x61\x28\x04\x27\xb0\x87\x19\xef\xe2\x5a\xda\x61\xce\x4e\x88\x27\xe2\xee\xb8\x58\xb3\x18\xd3\x04\xbe\x3f\x87\x22\x4a\xfc\x53\xc0\x5e\xdc\xf0\xce\x3b\xdf\xb7\xe7\xd6\x80\x57\x8a\xfa\xe7\xff\x90\xd0\x2b\x89\xba\x0a\x7f\x6f\xb8\x91\xf2\x38\xc8\xa5\x0e\xdc\x17\xad\xe8\x7a\xcc\x56\x93\x39\x3d\x9e\xfb\xf5\x8a\x7c\x94\x5f\x01\x00\x00\xff\xff\x6f\xd3\xe1\x13\x47\x01\x00\x00")

func outputsTfBytes() ([]byte, error) {
	return bindataRead(
		_outputsTf,
		"outputs.tf",
	)
}

func outputsTf() (*asset, error) {
	bytes, err := outputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "outputs.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0xd7, 0xa7, 0xff, 0x49, 0x7d, 0xf7, 0xf2, 0xf0, 0xad, 0x3d, 0x63, 0x4d, 0x52, 0x40, 0x3f, 0x63, 0xb2, 0x30, 0x39, 0xfe, 0xb0, 0xbb, 0x6b, 0x82, 0xfe, 0xd3, 0xf7, 0xce, 0xb2, 0x30, 0x67}}
	return a, nil
}

var _variablesTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\xc8\x7b\x2c\xd4\x0a\x10\x8b\x9c\xc5\x9a\xc4\xd3\x60\xe1\xd8\x66\x66\xdc\x12\xa1\xde\x1d\x25\x5d\x40\x95\xa8\xf1\xd2\xff\xfd\xe7\x2f\xf9\x8c\x1c\xb0\x8b\x04\xa6\x8f\x55\x94\xd8\x25\x1c\xc9\xc0\x4f\x73\x6d\x9a\xbf\x10\x4b\x71\x21\x89\x62\xea\xc9\xf5\xb9\x26\x7d\x8c\xe8\x54\x36\x24\x03\x25\xdd\xd5\xdc\x43\x9b\xa2\xc2\xf9\x7b\xda\x61\x7c\xb7\xf7\xd2\x7f\x82\xd2\x10\xd2\x8e\x24\xa2\xc8\x16\x71\xeb\xba\x33\xb1\x84\x9c\x66\x02\x60\x9e\x04\x2d\x8c\x58\x1a\x00\x4f\x27\xac\x51\xa1\x5d\x22\x00\x83\x95\x33\xe3\xd3\x38\xc9\x57\x34\xb0\x9c\x16\xcc\x8b\x7d\xb3\xcb\x95\xbb\xe5\xf6\x60\x9f\x8f\xf6\x60\xee\x3a\x25\x8b\x0e\x4c\x4b\xb1\x05\xf3\x6e\x5f\xed\x71\x26\xae\xab\x59\x55\x88\x37\xff\xd2\x77\xae\xa0\xc8\x25\xb3\x5f\x65\x22\x1f\xae\xd4\x2e\x86\xde\x7d\xd2\xb4\x8a\x47\x54\x25\x1e\xb3\xa8\xf3\xf9\x92\x62\x46\xef\x2a\xc7\x47\x5c\x0c\x3d\x25\x21\x77\x0a\x71\x3d\x65\x11\x28\x6d\xea\x7e\x03\x00\x00\xff\xff\xf9\xe0\x1c\xf1\x9a\x02\x00\x00")

func variablesTfBytes() ([]byte, error) {
	return bindataRead(
		_variablesTf,
		"variables.tf",
	)
}

func variablesTf() (*asset, error) {
	bytes, err := variablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "variables.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb5, 0xeb, 0xf8, 0xdb, 0x1e, 0xf0, 0xf6, 0x16, 0x30, 0xa2, 0x33, 0xa0, 0x89, 0x3, 0x8, 0xea, 0x2b, 0x44, 0xf8, 0x6c, 0x9f, 0xdd, 0xd1, 0xd1, 0xf4, 0xdd, 0xa7, 0x36, 0xe8, 0x17, 0x10, 0x9b}}
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
	"cluster.tf":          clusterTf,
	"dashboard.yaml":      dashboardYaml,
	"dashboard_data.json": dashboard_dataJson,
	"datasource.yaml":     datasourceYaml,
	"outputs.tf":          outputsTf,
	"variables.tf":        variablesTf,
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
	"cluster.tf":          &bintree{clusterTf, map[string]*bintree{}},
	"dashboard.yaml":      &bintree{dashboardYaml, map[string]*bintree{}},
	"dashboard_data.json": &bintree{dashboard_dataJson, map[string]*bintree{}},
	"datasource.yaml":     &bintree{datasourceYaml, map[string]*bintree{}},
	"outputs.tf":          &bintree{outputsTf, map[string]*bintree{}},
	"variables.tf":        &bintree{variablesTf, map[string]*bintree{}},
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
