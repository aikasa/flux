// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 9, 2, 10, 53, 11, 743589615, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(2019, 9, 2, 10, 53, 11, 735589583, time.UTC),
			uncompressedSize: 6456,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5f\x6f\x1b\xb9\x11\x7f\xf7\xa7\x18\x28\x0f\x49\x00\x69\x65\xc7\xb9\x43\xb1\x57\x1f\x90\x4b\x2e\x6e\x9a\x8b\x63\xc4\x4d\x8b\x3e\x35\x14\x77\xa4\x25\xc4\x25\xb7\x1c\x52\x3a\xc1\xb8\xef\x5e\x0c\xb9\x7f\xb8\x96\x9c\x1c\xf2\xd6\x3c\xc4\x36\x77\xfe\xcf\xf0\x37\x33\x5c\x2c\x16\x67\xa2\x55\xff\x44\x47\xca\x9a\x12\x44\xdb\xd2\x72\x77\x71\xb6\x55\xa6\x2a\xe1\x0d\xb6\xda\x1e\x1a\x34\xfe\xac\x41\x2f\x2a\xe1\x45\x79\x06\x60\x44\x83\x25\xac\x75\xf8\xfd\xfe\x1e\xd4\x1a\x8a\x1b\xd1\x20\xb5\x42\x22\xfc\xf1\x47\xf7\x3d\xfe\x59\xc2\xfd\xfd\xf4\xeb\xfd\x3d\xa0\xa9\x98\x8c\x5a\x94\x2c\xcc\x61\xab\x95\x14\x54\xc2\xc5\x19\x00\xa1\x46\xe9\xad\xe3\x2f\x00\x8d\xf0\xb2\xfe\x4d\xac\x50\x53\x3a\xc8\x75\x33\xb5\x77\xc2\xe3\xe6\x90\x3e\xfa\x43\x8b\x25\x7c\x42\xe9\x50\x78\x3c\x03\xf0\xd8\xb4\x5a\x78\xec\x84\x65\x1e\xf0\x3f\x61\x8c\xf5\xc2\x2b\x6b\x06\xe1\x00\xad\xb3\x0d\xfa\x1a\x03\x15\xca\x2e\x5b\xeb\x7c\x09\xb3\xcb\xf3\xcb\x8b\x19\x3c\x01\x8f\x5a\x67\x14\xe0\x2d\x90\x74\xa2\x45\x58\x36\xe8\x9d\x92\xc4\xce\xb5\x56\x19\xff\x94\x80\x99\x8b\x4e\xb0\x9e\xf8\xf0\xc0\x0b\x80\x3e\x16\xf1\x77\x74\x3b\x25\xf1\x95\x94\x36\x18\x7f\x33\x25\x04\xd8\x59\x1d\x1a\x1c\x44\x2d\x3a\x51\x1b\xe5\x17\x5b\x3c\x0c\x0a\x88\xa3\xe0\x47\x85\xfd\xc9\x28\x6f\xc1\x2c\x55\x4c\x70\x46\x55\xe1\x5a\x04\xed\x3f\xd8\x0a\x4b\x38\x7f\x79\x7e\x0e\x4f\x60\x5f\xa3\x81\x86\xad\xc1\x0a\x1c\x8a\x6a\x61\x8d\x3e\xcc\x61\x8f\xb0\xb7\xe6\xa9\x87\x15\x82\x58\x69\xe4\x78\xc8\xba\xb1\xd5\x59\x27\xf0\x09\xfc\xa3\x56\x04\x8a\x40\x80\x6f\xda\x35\x41\x20\xac\x60\x6d\x1d\x6c\xd0\xa0\x13\x5e\x99\x0d\xdc\xdd\xfd\x0d\xb6\x78\xa0\x02\xde\x19\x78\xff\x17\x82\x9f\xaf\xe0\xa2\xb8\x38\x9f\x0f\x52\x7a\xdd\xc9\x05\x02\xe1\x30\xb7\x83\x2c\x9b\x62\x10\x2b\x10\x40\xd8\x0a\x2e\x8a\x2e\x50\xb0\xc7\x41\x8c\x14\x06\xf6\x4e\x79\x36\xb4\x38\x1d\xbf\x0d\x9a\x21\x18\xd8\xb4\xfe\xf0\x46\xb9\x3c\x88\x0d\x56\x2a\x34\x25\x7c\xc0\xc6\xba\x43\xee\x27\xc2\xda\x6a\x6d\xf7\xec\x51\xa7\x5a\x51\x74\x35\x10\x9f\x09\x90\x81\xbc\x6d\x14\x47\x60\x6b\xec\xde\xfc\xa7\xb6\xe4\x69\x10\xb1\x56\x1a\xe7\xb0\xaf\x95\xac\xe1\x60\x03\xec\x95\xd6\xc9\x29\x6f\xa1\xb2\x7c\xcf\xf8\x98\x99\xf8\x17\x07\x76\x6f\xd8\xec\x41\x80\xc3\xd6\x82\x13\xbe\x46\x07\xbe\x16\xa6\x53\xbc\x51\xbe\x0e\x2b\xb0\x7c\x88\xa0\xd5\x16\x0b\xf8\xb7\x0d\x4f\xb5\x06\xa1\xc9\xf6\x2a\xa6\xc1\x06\xe5\x41\x19\x6f\x23\x8f\xb4\xc6\x0b\x65\xd0\xcd\x61\x85\xda\xee\x0b\xb8\xc3\x31\xaa\xb5\xf7\x2d\x95\xcb\x65\x65\x25\x15\x5c\x58\xb2\xe2\xab\x83\x66\xc9\x57\x8f\xfc\x72\x13\x54\x85\xb4\x0c\x84\x8b\xd6\xa9\x9d\xf0\x18\x4b\x8f\x1d\x29\x6a\xdf\xe8\x41\x52\x9f\x0b\xa2\x7a\x21\xad\x59\xab\xcd\xf0\x09\x20\x1d\x7c\x10\x6d\x99\x1d\xe6\x17\x69\x91\xb1\x7d\x6f\x5e\x8a\x6d\x58\xe1\x32\x09\x19\xcb\xef\x9b\x39\xd9\x2b\xaa\xf9\xa4\x16\x3b\x04\x01\x95\x5a\xaf\xd1\x31\x68\xf6\x12\xba\x5b\x35\x02\x63\x4c\x41\x12\x97\x27\x81\xc1\x65\xa7\x2a\xec\xc3\xbe\x56\x9b\x46\xb4\xa3\x21\xca\xd7\x20\x0c\xa0\xf1\xee\x10\x7d\xf8\x92\x88\xbe\xcc\x41\x98\x0a\x82\x91\xb6\x61\xb4\x8e\xfc\xc9\xdb\x0f\x31\x9d\xc2\x54\x83\x14\x34\xbb\x28\x41\x21\x75\xf9\x3c\xca\x00\x87\xe1\x3b\x32\x90\xb1\x7d\x33\x03\x11\x09\xbc\x05\xd5\x30\x4e\xc2\xf5\xed\x75\x04\x01\x78\xc6\x6e\x91\xda\x18\x65\x46\xe5\xec\xdc\x0e\x9d\x5a\x2b\x19\x01\x1b\xda\xe0\x5a\x4b\x48\xcf\xff\x44\x20\x07\x29\x09\x3e\x52\x14\x39\x40\xac\xef\x4f\x04\x0e\x84\xdb\x8c\xd7\xf4\x91\x88\x6d\xda\x0d\xe3\x07\x65\xa1\x99\x42\xf0\x93\x47\x40\xf8\x98\xef\x04\x08\xf7\xe1\x1c\x6e\xe2\x11\xfe\x67\x1d\xa2\x8b\xba\xc3\x88\x93\xc6\xc2\xac\x4c\x37\x71\x06\xaa\x11\x1b\x4c\xd5\xcf\x0c\x05\xbc\x55\xa6\x8a\x3e\x37\x0c\x2b\x0e\xe5\x58\xb5\x09\x52\x34\x0a\x42\x06\x8f\xc8\xca\x49\xe0\x39\x01\x84\x1f\xee\x7d\x1d\x56\x45\x65\xe5\x16\x5d\x21\x6d\xb3\x74\xcb\x3d\x8a\x1d\xee\xad\xdb\xd2\x92\x95\x2c\xbd\x18\xc2\xd7\xe7\x92\x7b\x3e\xcf\x03\xac\xd9\x8b\x0d\xb0\xb5\xc5\x40\x13\x55\x95\xd0\x09\x55\x76\x99\x50\x25\xfe\x28\x2f\x8a\x8b\x97\xc5\x8b\x29\xed\x6d\xd0\xfa\xd6\x6a\x25\x0f\x25\xbc\x5b\xdf\x58\x7f\xeb\x90\x72\x4f\x1c\x92\x0d\x4e\x22\xe5\x58\xee\xf0\xbf\x01\xc9\x4f\xce\x00\x64\x1b\x4a\xf8\xe1\xbc\x99\x1c\x36\x11\xee\x4b\xf8\xf1\xe5\x07\x35\x8e\x0a\xd6\xe5\xcc\x8b\x31\x3b\xb7\x71\x6c\xb8\x3c\xbf\xe4\xee\xa9\xcc\xda\xba\x26\x96\xad\xd0\x03\xb5\x56\x3b\x34\x48\x74\xeb\xec\x0a\x73\x0b\x38\xac\xd7\xd3\xce\x9d\x54\x25\x81\xd3\x63\xe1\xeb\x12\x96\xa2\x55\x29\xd2\xbb\x1f\x97\xaa\x42\xe3\x95\x3f\x14\x6d\x58\x65\xb4\xca\x28\xaf\x84\x7e\x83\x5a\x1c\xee\xf8\x8e\x56\x54\xc2\x0f\x19\x81\x57\x0d\xda\xe0\x4f\x7c\xe3\x46\xab\xfe\x3f\x4c\xcd\x2e\xee\x24\x31\xa7\x47\x24\x48\xad\xee\x36\x59\x86\x5e\x46\xcb\xaa\x25\x51\xcd\xb3\x9e\x4d\xd3\x27\x68\xdb\x61\xce\x86\x53\x06\xca\xa4\x9a\x7b\x4a\x89\x87\xa8\x5e\x4e\xa0\xb2\x8f\xd9\x47\xa3\x0f\x25\x78\x17\x90\xa5\xf1\x1c\x14\x51\x6a\xd5\x81\x3b\x5f\xab\x16\xdd\xda\x3a\x89\x2c\x34\x0d\x3e\x3c\xf7\x3c\x66\x78\x3e\x9b\x4c\x6d\xdf\x09\xd7\xd9\x9e\xc8\xbe\xcf\xfc\xec\x8e\xbe\x33\x52\x87\x88\x9e\x3c\xbe\xa5\x26\xd7\x23\x6b\x9a\x0f\xbe\x31\xce\xf4\x03\xcd\x4f\xcc\xfa\x60\xd4\x18\x10\x16\x2a\x94\x5a\x38\x1e\xdb\x56\x76\x97\x01\xc0\x57\x46\x81\x04\x91\xb9\xf3\xce\x5a\xbf\x2c\x88\xea\x47\x1d\x10\x66\xa2\x75\x36\xb6\xa9\x59\xd2\x3c\xef\x49\x32\x09\x68\x76\xca\x59\x13\x9b\x42\xea\xb7\xb3\xf7\x9f\x7f\xf9\xf5\xf5\xc7\x9b\xb7\xef\xae\x67\xa9\x0d\xcc\x39\x1e\x76\x87\xce\x4d\x7b\x76\x26\x26\xb6\xb9\xd5\x21\x75\x54\xaf\x4f\xf9\x78\xd4\x6c\x8f\x7d\x1c\x8b\x93\x89\x1f\x75\x94\xfb\x1e\x2f\x1f\xbd\x36\x86\xe9\x6c\x1c\xe9\xac\x8b\x39\xc9\x44\x3c\x1c\x6a\xf2\xa4\xc7\x89\xa6\x1f\xbf\x85\x01\xa1\x3d\x3a\xc3\xe3\xf5\x91\xc5\x6b\x67\x1b\x2e\x8b\x7e\x6a\x99\x83\x20\x2e\xb7\xae\xb3\x72\x18\xb4\x95\x5b\x3a\x4e\x36\x9a\x5d\x79\x22\x2e\x63\xb8\x27\x71\xd9\x09\x1d\xf0\x28\x26\xdf\x2a\xe2\x87\x35\xd0\xf7\xdd\xaf\x54\x00\xb7\xfd\x69\xbb\xff\x4a\xc3\x7f\xa4\x2e\x99\x2a\x4d\x38\x13\xba\x29\x3e\x8c\x46\xb3\xca\x72\xe2\x43\x4a\x43\x5a\xd3\xb0\xe2\x46\x24\x85\xac\xb1\xe2\xc8\xe6\xa9\x1d\x26\x4b\x4e\x22\x87\x65\x9e\x49\xb1\xae\x1b\x1d\x33\x86\x6e\xcd\x8c\x8c\xf3\xa8\x84\xd7\x23\x0a\x6d\xab\x0f\x1c\x08\xca\x43\x31\x0e\x70\x7e\x6f\xd9\xca\xc0\x29\x8d\x05\x17\x77\xe2\x98\x07\xa8\xed\x3e\xae\x80\xd6\x18\x94\x3e\x0e\x77\x7e\x1a\xba\xc5\x62\x70\x20\xce\xff\xac\xfc\x6a\x38\x2a\xba\xb9\xa7\xa0\x9d\x2c\xa4\x0e\xe4\xd1\x15\x8c\x5f\x3a\x0f\xc9\x67\x4a\x57\x6d\x0c\xc5\xeb\x44\xfa\xee\x76\xe2\x14\xdf\x3a\x42\x1f\x57\xcc\x69\x62\x47\x1b\x7a\x7a\x5e\xe4\xbd\x63\xca\xb8\xf4\x65\x08\x9c\x5b\xdc\x51\x5f\x9d\x4d\x06\x2d\x45\xd0\x04\x8a\x4b\x70\x8c\x9e\xc2\x2a\x55\xd3\x2a\xe2\x7a\x1c\x71\xe2\xee\xfb\xac\x5f\x28\x9f\xe7\xb6\xf4\x77\x2b\x55\x21\x4f\x66\xd9\x0a\x3c\x31\x84\xb1\x30\xe1\xfb\xa2\x52\xee\xea\x08\xf5\x73\xb3\x3e\x65\x03\xd6\x98\xbc\xcf\x9f\x7e\x4b\x3b\xba\x30\x9b\xf4\xed\x5a\xf9\xb8\x37\x92\xf2\xd6\x1d\x06\xb4\x7a\xcb\xc3\xe1\x44\x39\xf7\xa0\xe0\xf4\xd5\xfd\x3d\x14\xd7\xca\xb3\xa4\xf8\xd4\x33\xa5\x58\x39\x61\x64\xdd\x13\xfd\x12\xff\x4a\x8f\x3e\x6a\x1d\x8f\xf8\x6e\xd0\x29\x4e\x9e\x0f\x98\xef\x2e\xa6\x81\xfe\x6e\x95\xc9\x18\x66\xf3\x59\xf7\x76\xa4\x09\x73\x76\x1e\xaf\x8e\x5b\xd5\x5e\x98\x58\x7e\x0e\x39\xab\x32\x4d\xf5\x8d\x30\x6a\xcd\xf3\x1e\x17\x28\xa9\x0a\x5d\xf2\xf5\xc1\xe4\x1c\x77\x5e\x4b\x08\xc1\x54\xe8\x1e\x04\xd0\xa1\x16\x5e\xed\x30\x8e\x33\xd4\xa7\x77\x33\x09\xe2\x83\x82\x1f\x9c\xa3\xb0\xaa\x94\xbb\x98\xa7\x9f\x2f\x86\x87\xb0\x31\x38\xf1\xa1\xeb\x54\x70\xe2\xeb\x51\x1f\xd5\x9e\xea\x84\x80\xcf\x84\xee\x14\x7f\x20\x74\x43\xe6\x98\x06\x4e\xf3\xff\xda\x08\x75\xd2\x00\xe4\x0f\xbd\x84\x9e\x6a\x7c\xca\x3b\x09\xba\xc8\xf7\x74\x6f\x39\xa0\x68\xe2\xf3\x10\xc7\x89\xbb\x81\xf2\x0f\x16\xbc\x3c\x56\x1d\xae\x76\xa8\x79\xf5\x15\x18\xed\x39\x3a\x59\xcc\x75\xf5\xd7\x2d\x1e\x40\x55\x3f\x0f\x64\x5f\x69\x95\x99\x55\x2c\x42\xf8\xe0\x70\xb2\x65\x9e\xd0\x15\x3f\x1f\x16\x03\x3d\x4d\xb0\xa0\x87\x42\x50\x1e\x6a\x41\x11\xe6\xad\xd1\x07\x10\x52\x22\x25\xb8\xac\x31\x3d\xd4\x3c\xeb\xdf\x04\xbe\xac\x85\x26\xfc\xf2\xfc\x84\xb6\x9e\x7f\x1a\x60\xf2\x2e\x48\x9f\x14\xed\xe3\x9e\xc7\x7d\x3f\x78\xa0\x83\x91\xb0\xb2\x76\xbb\x45\x6c\xb9\x5c\x07\x1d\xb3\x8d\xf2\xb3\x39\x34\x28\x38\x52\x7c\xcd\x41\xc4\xc5\xab\xab\xe0\xd0\x92\x77\x28\x9a\xa1\x94\x1f\x5a\xc3\xa2\x17\xe4\x85\xc7\xab\x8d\xf2\x8f\x27\xdc\xe0\xef\xbe\xcf\x7a\xd6\x07\x84\x81\x59\xaf\x63\xd6\xa3\x74\x26\xe4\x19\x16\x9b\x62\x0e\xff\xe2\xcd\x11\x5e\x6b\x1b\xaa\xe7\x45\x7c\x39\xf0\x76\xcb\x43\x2b\x41\x2b\x9c\x57\x32\x68\xe1\xfa\x28\x76\x52\x1e\x36\x98\x4e\xeb\xd5\x9e\x78\x39\x95\x2c\xab\x88\x1b\x69\x91\x56\xd2\x7e\x03\x79\xc0\x16\x15\x5d\x89\x95\xbc\x78\x71\x79\xfc\x7f\xee\xf0\x1d\xba\xdd\x89\x07\x5f\x9e\xb5\xc6\xee\xca\xa5\xfa\x53\x0e\xf3\x62\xcb\xed\x21\xe5\x8a\xd0\x67\xaf\xc8\x4f\xb3\x87\xe8\xec\x45\x99\x5d\x8c\x2f\x23\x71\xde\x99\x82\xb1\x56\xe4\xd1\x2c\x3a\x13\xae\xca\xcb\xf3\xcb\x8b\xb3\xee\x1a\xbf\xaa\x2a\x95\x76\x4d\x06\xf1\x57\x3c\xc3\x4c\xf0\x72\xfc\x3e\xf6\xf1\xfb\x7b\x70\xb1\x25\x7c\x83\x7b\x11\x9f\xf3\x27\x57\x7f\xfc\xad\x57\xf0\xb1\xed\xc4\xbf\xb9\xb9\xeb\x1b\x30\xcd\xbb\xb9\x30\xb8\xae\x1d\x83\xa9\xac\x27\xb0\x91\x18\x1a\x71\x88\x3b\xba\xde\x8d\xaf\x35\x86\xb4\xb5\xdb\xd0\x82\x22\x0a\x48\x60\x0d\x90\x6d\x10\xde\x87\x15\x3a\x83\x1e\x89\xa5\x87\x96\xc6\xc7\x98\xca\x50\xff\x0c\x30\xbb\xb1\x06\x67\xf9\x97\xd7\xd1\x80\xfc\x39\x26\x29\xa7\xe9\x0b\x4d\x3f\xdf\x45\xfb\x26\x5f\x86\xd1\x73\x76\x31\x3b\xfb\x5f\x00\x00\x00\xff\xff\x6c\x75\x4a\x7a\x38\x19\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(2019, 8, 6, 11, 23, 6, 137011753, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
