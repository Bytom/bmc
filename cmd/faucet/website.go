package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _faucet_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xeb\x8f\xdb\x36\x12\xff\xec\xfd\x2b\xa6\xba\xb4\xb6\x6f\x57\x92\x9d\xed\x0b\xb6\xe4\x43\x92\xf6\x8a\x1c\x90\xb6\xb8\xa4\xb8\x2b\xda\x7e\xa0\xc5\xb1\xcd\x2c\x45\xaa\xe4\xc8\x5e\xd7\xf0\xff\x7e\x20\xf5\xb0\xfc\xd8\xbd\x3c\x0a\x1c\x2e\x1f\xbc\x22\x39\x9c\xf9\x71\x66\x38\x0f\x29\xc9\x27\xdf\xfc\xf0\xe2\xcd\xcf\x3f\x7e\x0b\x2b\xca\xe5\xec\x2a\x71\x7f\x40\x32\xb5\x4c\x03\x54\x81\x9b\x40\xc6\x67\x57\xbd\x24\x47\x62\x90\xad\x98\xb1\x48\x69\x50\xd2\x22\xfc\x3a\x68\xe7\x57\x44\x45\x88\xbf\x97\x62\x9d\x06\xff\x0e\x7f\x7a\x16\xbe\xd0\x79\xc1\x48\xcc\x25\x06\x90\x69\x45\xa8\x28\x0d\x5e\x7e\x9b\x22\x5f\xe2\x61\x9b\x62\x39\xa6\xc1\x5a\xe0\xa6\xd0\x86\x3a\x94\x1b\xc1\x69\x95\x72\x5c\x8b\x0c\x43\x3f\xb8\x01\xa1\x04\x09\x26\x43\x9b\x31\x89\xe9\x38\x98\x5d\x5d\xf5\x12\x12\x24\x71\xb6\xdb\x45\xdf\x23\x6d\xb4\xb9\xdb\xef\x27\xf0\x77\x56\x66\x48\x49\x5c\xad\x39\x2a\x29\xd4\x1d\xac\x0c\x2e\xd2\xc0\x21\xb5\x93\x38\xce\xb8\x7a\x6b\xa3\x4c\xea\x92\x2f\x24\x33\x18\x65\x3a\x8f\xd9\x5b\x76\x1f\x4b\x31\xb7\x31\x6d\x04\x11\x9a\x70\xae\x35\x59\x32\xac\x88\x6f\xa3\xdb\xe8\xab\x38\xb3\x36\x6e\xe7\xa2\x5c\xa8\x28\xb3\x36\x00\x83\x32\x0d\x2c\x6d\x25\xda\x15\x22\x05\x10\xcf\x3e\x48\xec\x42\x2b\x0a\xd9\x06\xad\xce\x31\xfe\x3c\xfa\x2a\x1a\x79\x89\xdd\xe9\xc7\x85\x5e\xf5\x12\x9b\x19\x51\x10\x58\x93\xbd\xb3\xd8\xb7\xbf\x97\x68\xb6\xf1\x6d\x34\x8e\xc6\xf5\xc0\x8b\x79\x6b\x83\x59\x12\x57\x0c\x67\x1f\xc3\x3a\x54\x9a\xb6\xf1\xd3\xe8\xf3\x68\x1c\x17\x2c\xbb\x63\x4b\xe4\x8d\x20\xb7\x14\x35\x93\x7f\x96\xd8\x87\xec\xf7\xf6\xd4\x7c\x7f\x82\xac\x5c\xe7\xa8\x28\x7a\x6b\xe3\xa7\xd1\xf8\xeb\x68\xd4\x4c\x9c\xb3\x77\xfc\x9d\xbd\x66\x57\xbd\x5e\xb4\x46\x43\x22\x63\x32\xcc\x50\x11\x1a\xd8\x5d\xf5\x7a\xbd\x5c\xa8\x70\x85\x62\xb9\xa2\x09\x8c\x47\xa3\x4f\xa7\x17\x26\xd7\x2b\x3f\xcb\x85\x2d\x24\xdb\x4e\x60\x21\xf1\xde\xcf\x30\x29\x96\x2a\x14\x84\xb9\x9d\x40\xc5\xd5\xcd\xef\x9d\xb4\xc2\xe8\xa5\x41\x6b\x2b\x31\x85\xb6\x82\x84\x56\x13\xe7\x44\x8c\xc4\x1a\xcf\x09\x6d\xc1\xd4\x29\x35\x9b\x5b\x2d\x4b\xc2\x63\x00\x73\xa9\xb3\x3b\x3f\xe5\xaf\x6a\x07\x79\xa6\xa5\x36\x13\xd8\xac\x04\xb5\x12\x0a\x83\x35\x5b\xc6\xb9\x50\xcb\x09\x7c\x59\x54\xf8\x73\x66\x96\x42\x4d\x60\x54\x93\x26\x71\xad\xad\x24\xae\xa2\xd0\x55\x32\xd7\x7c\x3b\xbb\x4a\xb8\x58\x43\x26\x99\xb5\x69\x70\xa2\x46\x1f\x5c\x3a\xcb\x2e\xa4\x30\xa1\xaa\x85\xa3\x15\xa3\x37\x01\x78\x01\x69\x50\x49\x0e\xe7\x9a\x48\xe7\x13\x18\x3b\x44\x7e\xc3\x09\x2f\x19\xca\x65\x38\x7e\x5a\x2d\xf5\x92\xd5\xb8\x61\x40\x78\x4f\xa1\xd7\x7f\xab\xf9\x60\x96\x88\x66\xe7\x82\xc1\x82\x85\x73\x46\xab\x00\x98\x11\x2c\x5c\x09\xce\x51\xa5\x01\x99\x12\x9d\x87\x88\x19\x74\x83\x58\x1b\xc3\x56\xe3\x0a\x45\xcc\xc5\xda\x1f\xa0\x7d\x38\x39\xc9\x43\x60\xbf\x86\xfa\x41\x2f\x16\x16\x29\x6c\xb1\x77\x48\x85\x2a\x4a\x0a\x97\x46\x97\x45\xbd\xda\x4b\xfc\x1c\x08\x9e\x06\xa5\x91\x41\x1d\xa9\xfd\x23\x6d\x8b\xfa\xc0\x41\x7b\x3c\x6d\xf2\xd0\x69\xda\x68\x19\x40\x21\x59\x86\x2b\x2d\x39\x9a\x34\xf8\x59\x97\x06\x9e\x6f\x49\xe7\xf0\x5a\x70\x84\x17\x2b\x26\x14\x30\xce\x9d\x8f\x45\x51\xd4\x0a\xf4\xee\x76\x0e\x28\x9c\x93\x6a\x68\x1c\xd9\xbc\x24\xd2\x2d\xe1\x9c\x14\xcc\x49\x85\x1c\x17\xac\x94\x04\xdc\xe8\x82\xeb\x8d\x0a\x49\x2f\x97\x2e\xf7\x54\x60\xab\x4d\x01\x70\x46\xac\x5e\x4a\x83\x86\xb6\xb1\x08\xb3\x85\x2e\xca\xa2\xb6\x49\x35\x89\xf7\x05\x53\x1c\xb9\xb3\xa0\xb4\x18\xcc\xbe\x13\x6b\x84\x1c\xe1\xf9\x9b\x57\xbd\x53\xf3\x66\xcc\x20\x85\x5d\x96\x67\x46\x4e\xe2\x0a\x4a\x75\x20\xa8\xff\x25\xa5\x6c\x38\xb5\x07\xc8\x51\x95\x70\x34\x0a\x8d\x8b\x00\xc1\x6c\xb7\x33\x4c\x2d\x11\x9e\x08\x7e\x7f\x03\x4f\x58\xae\x4b\x45\x30\x49\x21\x7a\xe6\x1f\xed\x7e\x7f\xc4\x1d\x20\x91\x62\x96\xb0\xc7\x5c\x15\xb4\xca\xa4\xc8\xee\xd2\x80\x04\x9a\x74\xb7\x73\xcc\xf7\xfb\xa9\xdd\xe6\x73\x2d\xd3\xfe\xf3\x37\xaf\xfa\x53\xd8\xed\xc4\x02\x9e\x44\xff\xc4\x8c\x15\x94\xad\xd8\x7e\xbf\x34\xcd\x73\x84\xf7\x98\x95\x84\x83\xe1\x6e\x87\xd2\xe2\x7e\x6f\xcb\x79\x2e\x68\xd0\xf0\x72\xf3\x8a\xef\xf7\xee\x00\x35\xe8\xfd\x3e\x89\xd9\x2c\x89\xa5\x98\xd5\x8b\xc7\x6a\x89\x4b\xd9\x5a\x3e\x89\x9d\x83\xfc\x7f\x39\xcb\x8f\xb8\x5c\x6e\x81\xf4\x1d\x2a\xfb\x3f\x72\x16\x68\xbd\xa5\x32\xe5\x0d\x3c\x99\x63\xf1\x14\x5f\xaa\x85\xf6\x3e\xf3\xbc\x19\x35\x6e\xe3\x15\xf7\x1e\x0e\x53\xbb\xc8\x6e\x57\x4b\xd8\xef\x3f\xcc\x51\x3c\x90\x41\x87\xcd\x91\xc3\xb4\xa0\x6b\x27\x7f\x4d\x66\xbf\x87\x0e\xf5\x87\x79\x52\x15\x4e\x3d\xdc\x2e\xda\xb3\x08\xb9\x0c\x5b\xfc\xb5\x63\x58\x41\x78\x87\xdb\x34\xd8\xed\xba\x3b\xeb\xd5\x8c\x49\x39\x67\x5e\x3d\xfe\x70\xed\xa6\x3f\xd0\x39\xec\x5a\x58\x5f\x1a\xcf\x1a\xf9\x2d\xe4\xff\x1e\xe8\x4f\x52\x16\xe9\x62\x02\xb7\x4f\x1f\xcb\x57\x5f\x9e\xa4\x80\xdb\x0b\x29\xa0\x60\x0a\x25\xf8\xdf\xd0\xe6\x4c\x36\xcf\xf5\x5d\x69\xa3\xf4\xe9\x96\xd0\x25\xe4\x16\x53\x9b\xd1\x47\x53\xd0\x6b\x34\x0b\xa9\x37\x13\x60\x25\xe9\x29\xe4\xec\xbe\x2d\x62\x6e\x47\xa3\x16\xb0\xe3\x4a\x6c\x2e\xd1\x27\x1b\x83\xbf\x97\x68\xc9\xb6\xa9\xa5\x5a\xf2\xbf\x2e\xc3\x70\x54\x16\xf9\x89\x12\x9c\x3c\xa7\x4b\x4f\xd5\x20\x6d\xf4\x77\x11\xf5\x42\xeb\xba\x56\xe8\x02\xa8\x99\x76\x8a\x98\x60\x96\x90\x39\x78\x0e\xf1\xf7\xca\xf7\xc6\x55\xe9\x0f\xa5\xfb\x2a\x86\xb9\x33\x17\x88\xa6\x2a\x13\x9d\x5f\x82\x1f\x26\x31\xf1\x0f\x96\xeb\x7c\x6d\xce\x2c\xbe\x8b\x70\x5f\xbb\x1d\x84\xfb\xe1\xc7\x49\x5f\x21\x33\x34\x47\x46\xef\x22\x7e\x51\x2a\xde\x39\xfb\xf3\x37\xaf\x3e\x4e\x78\xa9\xc4\x1a\x8d\x15\xb4\x7d\x57\xe9\xc8\x0f\xe2\xab\x71\x17\x40\x12\x93\x79\xd8\xb3\x0e\x8f\x67\xf7\xb6\xfe\x5b\xff\xb9\x4a\xda\x0e\x23\x8e\xe1\x3b\xa9\xe7\x4c\xc2\xda\x01\x9c\x4b\xb4\x40\x1a\x5c\xbd\x04\xb4\x42\xc8\x4a\x63\x50\x11\x58\x62\x54\x5a\xd0\x0b\x3f\xbb\xf0\xf5\xe0\x55\x6f\xcd\x0c\x30\x22\xcc\x0b\x82\xd4\x17\xca\x6e\xc6\xa2\x59\xfb\x5a\xdf\x0d\x5c\x06\xef\xae\x55\xf1\x39\x08\xea\x71\x73\xc3\x20\x85\x5f\x7e\x9b\x5e\x79\x40\xdf\xe0\x42\x28\x04\xe6\x14\x90\xb9\x4a\x1f\x68\xc5\x08\x32\x83\x8c\xd0\x42\x26\xb5\x2d\x4d\x85\xd3\x65\x19\x70\x58\x1b\x3e\x15\x57\x37\x5d\x78\xb9\x0d\x8b\xc1\x8a\xd9\xd5\xd0\x57\xfa\x06\xa9\x34\xea\xb0\x52\xcd\xf6\x16\xda\xc0\xc0\x6d\x16\xe9\x68\x0a\x22\x69\x38\x46\x12\xd5\x92\x56\x53\x10\xd7\xd7\x35\x69\x4f\x2c\x60\xd0\xac\xff\x22\x7e\x8b\xe8\x3e\x72\xfc\x21\x4d\xe1\x20\xa7\xe7\x44\xd5\x3c\x6c\x21\x45\x86\x03\x71\x03\xe3\xe1\xb4\x5a\x9b\x1b\x64\x55\x9b\xe2\xfb\x10\xff\xb3\xbf\xea\xed\xa7\x5d\x1d\x78\x65\x1f\x69\xa1\x8a\xe0\x16\x18\x2c\x85\x25\x28\x8d\x74\x7a\x70\x74\x95\xda\x6b\x35\x7b\xaa\xee\xf9\xcf\xb2\x4a\xfd\x50\x47\xfb\x0a\x72\xc5\x22\xb2\xa8\xf8\xe0\x1f\xaf\x7f\xf8\x3e\xb2\x64\x84\x5a\x8a\xc5\x76\xb0\x2b\x8d\x9c\xc0\x93\x41\xf0\x17\x57\x76\x0f\x7f\x19\xfd\x16\xad\x99\x2c\xf1\xa6\x36\xe9\x04\x9a\x94\xee\x2c\x3e\xf1\xbf\x67\x32\x6f\xa0\x7e\x9c\xc0\xb1\xf8\xfd\x70\x38\xbd\x94\xf7\x3a\x89\xda\xa0\x45\x1a\x38\xb2\x3a\x3d\x1d\x6b\x8a\x41\x8e\xb4\xd2\xdc\x69\xc3\x60\xa6\x95\xc2\x8c\xa0\x2c\xb4\xaa\x15\x03\x52\x5b\xdb\x38\x5d\xb3\x9e\x9e\xba\x41\x4d\x9b\x82\xc2\x0d\xfc\x0b\xe7\xaf\x75\x76\x87\x34\x18\x0c\x36\x42\x71\xbd\x89\xa4\xce\x98\x23\x77\x8d\x29\xe9\x4c\x4b\x48\xd3\x14\xea\xde\x3c\x18\xc2\xdf\x20\xd8\x58\xd7\xa5\x07\x30\x71\x8f\xee\x69\x08\xd7\x70\xba\x7d\xa5\x2d\xc1\x35\x04\x71\x75\x95\x5c\x96\x33\x14\xb3\x42\x04\x43\x77\x0b\x1a\x4b\x68\x95\xa3\xb5\x6c\x89\x5d\xa4\xb8\x46\x45\xb5\x8f\xb9\xe3\xe4\x76\x09\x29\x78\x7b\x15\xcc\x58\xac\x08\x22\x17\x77\x2b\x67\x73\xee\xea\x89\xd2\x14\x54\x29\x65\xe3\x9f\xd5\x4d\x98\x56\xde\xd7\x21\x8c\x7c\x1c\x84\x4f\xd2\x14\x5c\x08\x72\xfa\xe5\xcd\x1e\xe7\x01\x55\x98\x1c\x46\x2e\x06\x1e\xe8\x87\xd3\xc6\x8d\x8f\xf8\x20\x7f\x9c\x11\xf2\x53\x4e\xc8\x2f\xb0\xf2\x79\xe8\x61\x4e\x55\xd6\xea\x30\xf2\x13\x17\xf8\xa8\x32\x9f\xa3\x79\x98\x51\x95\x81\x6a\x46\x5e\x9d\x2f\x15\x75\x76\xde\xc0\xf8\xcb\xe1\x05\xbe\x68\x8c\x7e\x80\xad\xd2\xb4\x1d\xec\x24\xdb\xea\x92\x26\xd0\x27\x5d\xbc\xf0\x09\xa3\x7f\x03\x4e\xca\x04\xda\xfd\x37\xbe\xe8\x9f\x40\xdf\x8f\xdc\xba\xc8\xd1\xef\xfa\x62\x34\x1a\xdd\x40\xf3\x36\xe4\x39\x73\x37\xcc\x94\xb8\xbf\x80\xc4\x96\x59\x86\xf6\x01\x5d\xbd\x13\x96\x9a\x43\x8b\xa6\x1e\x7f\x20\x9e\x36\xc4\x1f\x01\x82\xcf\x3e\x83\xb3\xd5\xae\x73\xc6\x31\xbc\x62\xe6\x0e\x7c\xf5\x67\x70\x2d\x74\x69\x0f\xe9\x22\x17\xd6\x0a\xb5\x04\x66\x81\x6b\x85\x7e\xc7\xfb\x44\xf0\x33\x74\x35\x11\xcc\x60\x74\x0a\xcd\xc5\xba\x4e\x84\xbf\x10\xf8\x5b\xae\xdd\xa8\xde\xdb\x1f\x24\x1d\xed\x11\x39\xc2\x27\x29\x04\xc1\x61\xdb\xd9\xba\x5b\xae\xd9\xf4\x2c\xd2\x9b\x4a\xef\x83\x3a\xb1\x5d\x4a\x3e\xc3\x1b\x57\xc8\x8e\x86\x47\xc2\xf7\x8d\x2a\x9f\x15\x05\x2a\x0e\x4c\x6d\x7d\x64\x6b\xf5\x28\x14\x69\xd0\xa5\x8b\x8e\x19\x93\xae\x34\x97\xe8\xa3\x8c\xdf\xe8\x94\x99\xe9\x3c\xd7\x0a\x52\x08\xc7\xd3\xb3\xe4\xd7\xd1\x5a\x7b\x98\x53\x33\x5c\xd0\xf2\xb1\x29\x8e\x35\x74\x42\x1a\x8e\x8f\x94\x7f\x64\x97\x4b\x06\xe8\xb5\x78\x45\xa3\xbf\x23\xa3\x34\x56\xe9\xea\xa7\x83\xb8\xda\x7d\x3d\x7e\x27\xe0\xed\x62\x51\xda\xd5\xe0\x04\xda\x70\x7a\x6c\x81\x97\x84\x86\x11\xfa\x4e\xc4\x6b\x1c\x15\x09\x83\x67\x8a\x07\xa6\x5c\x39\x13\x1a\x54\x1c\x4d\x53\x01\xb8\x46\xa6\xea\x3b\x3a\x86\xf1\x9f\x27\x3a\xae\xd2\x39\xc7\x99\x16\xa7\x20\x60\xe6\xea\x30\x10\x61\xd8\x9e\xc0\x17\x4b\x5a\xa1\x6b\x4b\x4f\x3c\xdb\x7b\x61\xc7\x0d\x1d\x29\x4a\x56\x58\xe4\x90\x42\xf5\x52\x79\x30\x8c\x4a\x25\xee\x07\xc3\xb0\x1e\x9f\x72\x68\xd6\x7d\x4e\xf3\xc6\xa9\x30\x5f\xa7\x10\x24\x64\x5c\xd9\xdb\x0f\xe0\xfa\xd2\x9d\x72\xd9\xb1\x3f\x6b\xa4\x77\x37\x02\x24\xc4\x67\xbe\x8d\xaa\x8a\xf1\x5f\x03\xd7\xdd\x2e\x8d\x2e\x15\x9f\xb8\x92\x68\x70\xc6\x94\xad\x19\x31\xe3\x79\x0e\xa7\x70\x20\xf7\x4d\xf0\x04\x32\x67\x95\x29\x54\xbd\x96\xef\x5f\xa1\xed\x0e\xfd\x68\xae\x0d\x47\x13\x1a\xc6\x45\x69\x27\xf0\x79\x71\x3f\xfd\xb5\x69\x98\x7d\x85\xfe\x08\xd0\xc2\xe0\xec\x0c\x4f\x96\xf9\x17\x65\xd7\x10\x24\xb1\x23\x78\x9c\x49\x7b\xd0\xee\x2b\x6d\xb8\xd0\x83\x40\xfb\x02\xba\x9e\xcf\x05\xe7\x12\x1d\xd8\x86\xb9\xbb\x6d\xce\xe6\x87\x3b\x73\x2c\x0e\xea\xc6\xa3\xa1\xdf\x03\x4a\x8b\x0f\x12\xb7\xfd\x4b\xdf\x19\x3c\x74\x07\x15\x5e\xcf\x75\x2b\xe4\xa7\x4d\xdf\x6b\xa0\xfe\x14\xc1\x4b\xe3\xab\xa0\x41\x58\x3b\xd4\x0d\xf4\xad\xab\xc9\xb8\xed\x0f\xa3\x55\x99\x33\x25\xfe\xc0\x81\x4b\x28\xc3\x4a\x43\xbe\x21\x0a\x8e\xe3\xea\x19\x90\x43\x5f\xdd\x6f\x12\x53\xbf\x56\x5c\xbf\xb1\xa6\x33\x1c\x1c\xfa\xf4\xfe\x7b\x68\xe5\xb2\x84\x70\xce\x0c\x74\x07\x61\x93\x2d\xc1\x68\x27\xb9\x59\x9b\x33\xd3\xaf\x9a\x40\x5f\x37\x2b\xbd\x49\xfb\xb7\xa3\x16\x60\x65\x58\x6f\xd7\x7e\xed\x57\x27\x06\x70\x08\x9b\x0b\x38\x83\xdb\xd1\xc7\x23\xe5\x4c\x2d\xf1\x14\x3d\x19\x51\x20\x07\x96\x91\x58\xe3\x9f\x7e\x88\x8f\x56\xee\x7b\xc3\x73\x7e\xd7\xa8\xcd\xbb\xe5\x11\x56\xb7\xda\x6a\xf5\xaf\xee\x5e\x41\xec\x75\x7b\x0d\xc1\x85\x43\x3c\xe0\x79\x47\x44\x27\x97\xf7\xa1\x7b\xed\x7b\xf9\xa0\x9b\x24\x5c\xf9\xd9\xbe\x71\x1a\x46\x2b\xca\xe5\x20\x48\xc8\x7f\x64\x72\x38\xdb\xdd\x7e\x73\x35\x7d\xa8\xb7\xf6\xdd\xae\xc1\xf5\xc8\x78\xd2\xdd\x40\xa7\x8e\x68\x3b\xa0\xa6\x68\x00\xd7\x4d\xed\x7d\x43\xf5\x9a\x98\x21\x60\xf0\xd3\x4b\x28\x0b\xce\xc8\x25\x21\x0d\x2e\xc5\xf9\x64\xd4\x7e\x91\x9b\x33\x63\x61\xa1\xcd\x86\x19\x0e\xa5\x22\x21\xdd\xfa\x16\x98\xc1\xba\x26\xb3\x48\x2f\x5d\x54\x5a\x33\x39\x38\xe9\xb3\x9e\x0c\xfa\x51\xd7\xb0\xfd\x61\x84\x2c\x5b\x9d\x92\xf9\xac\xd3\x4a\x4c\xe1\x7b\x5f\x83\x0f\x9e\x0c\x68\x25\xec\x30\x62\x44\x66\xd0\x3f\x32\x78\x7f\xe8\x6c\x37\x6e\xfb\x9e\x76\x73\xd2\xb9\x32\x8f\xed\x3f\x54\xb5\x75\x06\x6f\x88\x33\x6b\x07\x95\xd7\xf4\x6f\x3a\x7c\x8f\x9d\xa6\xff\x69\xbf\x36\xc9\xe1\xd2\x1e\xf0\xa7\x17\x30\x1c\xb1\xed\xbb\xbb\xd3\x3f\x11\xcc\x38\x7f\xe1\xee\xc5\x20\xb8\x70\x7b\xbb\x1e\x30\xac\x15\x5b\x45\xdc\x47\x34\x2a\x14\xc7\xfb\x87\xd4\x29\x78\x7f\x18\xd9\x72\x5e\xf5\xfe\x83\x2f\xea\x6e\xa7\x21\xf2\x6e\x79\x1a\xc8\xcf\xd2\xbf\x13\x70\x5c\x02\x84\x27\x25\xc3\x23\x31\xdf\x0b\x74\xa7\xd9\xdf\x38\xe5\x8e\x86\xf5\x8b\xa1\x6f\xad\xab\x80\x84\x5d\x01\x83\x0d\xce\xad\xef\xcd\xa1\xf6\x64\xff\x86\xa4\x7a\x13\xf2\xec\xc7\x97\xed\xdb\x90\xd6\xd3\x5d\x11\xd2\x7e\xfb\x3e\x7f\xd7\x70\xf1\x53\xfb\x66\xb3\x89\x96\x5a\x2f\x65\xf5\x91\xbd\x7d\x19\xe1\x1a\xf5\xe8\xad\x0d\x80\xd9\xad\xca\x80\xe3\x02\xcd\xac\xc3\xbc\x7a\x43\x91\xc4\xf5\xf7\xe1\xb8\xfa\xdf\x2c\xff\x09\x00\x00\xff\xff\xfc\xe0\xa4\xb5\xde\x22\x00\x00")

func faucet_html() ([]byte, error) {
	return bindata_read(
		_faucet_html,
		"faucet.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"faucet.html": faucet_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"faucet.html": {faucet_html, map[string]*_bintree_t{}},
}}
