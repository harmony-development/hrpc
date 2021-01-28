// Code generated by "esc -o data_gen.go -pkg main templates"; DO NOT EDIT.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/hrpc-client-go.htmpl": {
		name:    "hrpc-client-go.htmpl",
		local:   "templates/hrpc-client-go.htmpl",
		size:    5103,
		modtime: 1611811948,
		compressed: `
H4sIAAAAAAAC/+RYW2/bNhR+z684FZKCKhT63Ugf1nTrijVNUa/YwzAUjHQkE6ZIjaSSBYb++0BS18ix
66bFCqwPjUCe2/edG5PtFk5vmEFYvgRiNS9XdZ7zf4DkXGDF7PqVu6TvWYkxRLTSyqoohqY52W7BoP2F
C5SsRCCm0lzaHKIzszgzdK2rlKaCo7S0UBGQDDXmQK8ry5U09I36wNINKzAOAXibzugpt1i6cKg/qYIU
bLcwDWlmyIvzslLaQiTRLtbWVlF/UihVCKSFEkwWVOli4cHc1Hn4GCS5WnBVWy6Go7y0I0vcrusbmqpy
USjNhWCLO7wxKt3gSMpFUOuRjZt7iyZqMfoz42BKvDNoO/SayQKBrlDf8jRAAgAYbq7QrlXWX7SXpxrN
W+mTqNEocYvZG/X7fYUtnRToW1nV1h95qh8oX9d2r/Z1bSfqY31WVSizFdoBV4iHjjJzWP66tmOFDjfK
rCNn9BnYON0kcHrrAu8tdZod6w7fBpomemiB5xOaQzlf/wZW193BzN+uvJwa1Ld9vfaHvikco6O+uPTt
MPSC76qBTuvYHZSbBozVdWph27MXGgpeuNKmwVx/5+JA/enjO6fGZeEvWtN5LVN4j3dj66TWohWN4cXE
7+BQo621hOfj6+F2CGnpPp+Pwto2yUSsj24JtRbDXTMO8/9a5q0Gz4FIZYEoDWTN2npZWY2sBBr7s5Xn
sT+LpzD7XJOuUsaJi52TcfF5ZaKDVMDy9rUXJC8GxsJZAqi10jFM0+97zXHu7x25fprSK6bNmgmi47l4
7kWfvQTJxQ5zo7qTXCSQl5b+7HznJEpVLTJwJJVMW84EaPy7RmOXcHYX+RjmDpuTuXlT9QG3W6r98UEZ
S5zLVWhcErmqXi4WZ2bRradRAmk3AgKhTeOEuu8o6Yz35R8nELGqEjxlbn8tVGrRnhufzygBvyToe7z7
iCxDTTy18bfl0OcRKmUsl8XX8ZdhjtqzSF+p7J5eCmWQzDUzZlnPc1iq1CH7SQjSa38PdBpZFtCZSkmD
x8FTfgy4mJ8/bIPtXNqF2pX9J1m2hR+gB0vfA2HdeRJfj7P1EIJMnKeT3QpuGQqDfkIxme2aRO3omo+t
+DtNqHTNJBwzpiomeUqiWvKyEliitJhF8VMQ7wD7ZKwxEC7Bgbs4n8H2BQUX53uwP46/dgVda0E/fXy3
XaVrLHEJ0Z2JEvhVuQnwcFYl8IHZ9RKi4wffaLH1D4UEPvezoH8w09eYs1rY15wJ1NT9IDVd+WcJiX1N
Pr17/H+o9Y6W2LHMLl2AJdvgtMT60ts1Lh7R6VMTzx0VyhcHiR8BEGZs2s/WnUKlKczU9Z9/uR3yiPgh
n75ild5z6/59TqBEY1iBww71Y/0qnO5YBMdmb/rCdAQ4pPFB2RuNbLNXqtl76/m8OO/wnRxnpXksTfs5
NShw8syfMeB+4y1NkYDaOLYvzl2cy0MsP1ObL6A3NMoBzvZe27VbQeF3WTKv/H2qu3enh+qtfttKejpU
3+oX5yG4/fkyKLMhYVxe/jD5mrzKps91F/SPxnmokZT+obnFbsQMC+QVl0zfX3UDyYH7DxAcOShOvmh0
tDuMy0u/+C8PPdFkduAvJ/8GAAD//x9Rp4fvEwAA
`,
	},

	"/templates/hrpc-server-go.htmpl": {
		name:    "hrpc-server-go.htmpl",
		local:   "templates/hrpc-server-go.htmpl",
		size:    7102,
		modtime: 1611812329,
		compressed: `
H4sIAAAAAAAC/+xYXW/bNhe+9684FfK+kAqF3oZdZe1F025tsaUJmha7GIaBlo4lwjKpklTUzNB/H0hK
MmXJcdKlzb50Y5k8hzzPcz54qM0GFOofWIGcrhHCTPTvBAKF8gplEEHTzGabDRwxjWs4eQrEjpQ0WdEM
YbOBJSuwpDo/pQqBnJeaCa7IS3HRihhxti6F1BAkgmv8qIN+gKOe51qX2xEm5kxUmhXboUyIrECSiYLy
jAiZzUsptFhUS/fiSTKdVwuSiPU8E5IVBZ3XuFAiWaG+1Xr6ukQ1T1ElkpVayHIRtPCtrjIMcKwV6o4Y
SXmGQC5RXrHEoQUA2M6coc5F2k+0k0cS1WtulgslKlFcYfpSvLsusWWaAHnNy0rbIeuFHeXzSt+ofV7p
gbqvT8sSeXqJeovL2UM8px2WP6+0r9DhRp525Hivjo2jVQxHV8bwfqVOs/OOwbeCpgkGK1xRaYON8oRh
+sZEqYPaNPDY9xcxUfyiH7gwngXrtPCQWDSbLSueAONMhxFsrFkp1dTY20b6C/O329kZjlIaARtD5D1f
U6lyWoRGMd5nc+QgL63yo6fAWdHuZ56ScpaEKKUTa2aOQrYcxJlL4PMfQcuqGxgRPhWYRya5+1zuB23q
m5BSpWRcLyH4n7p0ZQDCFCUugRgQ0TaeTLp4yk0DjGuUS5pgC+cvnQdsCSEXGkIhIcypel4w5PpSS6Rr
IJEdcwz0Y9HQiHYhjx1omjDRH6EtdeS5+41BwuMeL3n9omliyJGmKBWYCkhe2T+RRVL2sibHnLCJFJRS
tEHRJVuh0OKgPJ2ytwU4BhfdIw5RaUhyyneshn0YRwA+zRLGh7vejz08nShmLg0HMT+08BXlaYESlJZV
or1cdv4YJEk/971xZ6dpSk/YOzmG2pn4FlUpuMKfJdO+qVWZSWM89Acced8O+dbagvYG630Gh2pgmg2z
u1kUWZan2djSIFFXksP/byG6Ze0EVDwY9gk78UwdSnXUnExwM9zGPG+RpqfVconykv2OJwBff/XNt/FI
zKL15cZizfZvM/JBmO/nKXJ4X717dxFOkhyDxA/wuJ35UKHSkceYqplOclM5PpD3b38iF1Tn/vw/vAgn
pvMM5l2X6vUkpDvtOsKNUPcenOwE3UKk13F3nLsmlJjoeFYUltpTkV5HA40UlyihmyPPC6EwHIrsO+O7
Jyd+UJt0i6GORmIuf4bhNtstmp/uR7eFDSzXMbk+N/Qra9+2dI/BNe57HIv+Yg9EiDlHe3fmxMUQ8QLA
OrU9TcJoaLP9158Po3VNasbDzu+sxW+mHwhx3RocRuRZmoaBxcb1sfF3EENAy7JgCTU3tLlINOpjZVMq
GG72W9y6tia2/IQd4h0iPieu8dtEr3BfiW8uF/05N8oqU6Am6tJO9jCbd2u6wnDclOxJHc6K3c2m2qG/
d11yxPypYsJ4NOJp25d9cUTCnXZjX/dt5k6i1MqrQ1130rUkYW1rTWysfiAXZcL1etHEJrdOAPOsVaaG
1Pzy6+Ja7xaOQ3vapknIPTNtiVqjUjTbVuFa2aQ4c8NhtFf3EK3+k9jMMbCiG+UWEulqr0Szd8YS9uS4
AzO7nXYzgW7n6jKuBrVy1bwjaNsb2/Jw1tHpPLZpJny23ycKCxxcez45hPq67r6fiJXx7ZNjQ9TJTS59
JFYHfDmRBzfkxGD59tvnwSJ2uKA5VGbB+wzQ3QI5LSUqHX0uhriJYgNrXwRMFOyBu9cq23pbVPrBnF37
FeUNftTuCublyynjVF63CXM/brxzitzN/zfz/8Wi5EDnvFbZf2zeIVBH3fkB7kxId73gd/9SNm95wM52
j9vZsPE6fJlk3H6HHN4jp281Q/Tjkf5jknfG/xEAAP//gQh6mr4bAAA=
`,
	},

	"/templates": {
		name:  "templates",
		local: `templates`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"templates": {
		_escData["/templates/hrpc-client-go.htmpl"],
		_escData["/templates/hrpc-server-go.htmpl"],
	},
}
