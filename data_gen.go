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
		modtime: 1611712143,
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
		size:    6890,
		modtime: 1611712143,
		compressed: `
H4sIAAAAAAAC/+xY227cNhO+91NMBOeHFKy5f4tebeOLOG2ToHVsxAl6URQFLY0kYrWkQlKW3YXevSCp
s7ReO3XingQYlsjhcL5vDhzudguHl1QhrI7B15JtLoo4ZtfgxyzDnOr0xEySt3SDAXgkl0ILL4CqOthu
QaH+gWXI6QbBV7lkXMfgPVXLp4qkMg+JQnmFkiTCAz9CiTGQs1wzwRV5Jc5puKYJBs4Aq9MoPWQaN8Yc
YkdyJwXbLQxNmiiy4myTC6nBCwXXeK29doCjXqZa590IE0smCs2ybigRIsmQJCKjPCFCJksL+LKI3UtP
kum0uCSh2CwTIVmW0WWJl0qEa9RejcOKKgOFY6lQNwgl5QkCuUB5xUJnNgBAN3OKOhVRO1FPHkpUb7h1
lEQlsiuMXon3NznWlBEgb3heaDtk6RwtPiv0ravPCj1Y3l9P8xx5dIG6w+XsIT3298ufFbq/oMGNPGrI
6b06Ng7XCzi8Moa3mpqVjTMMvjVUlTfWwOIBzS5kz34ELYtmYLLfnF8OTSC3MdkO2sA3jPZi/8KGfBfv
NnM6OrVht1tcVcC4RhnTEGH71w8DFoPPhQZfSPBTql5mDLm+0BLpBkhgxxwD7VgwNKJW1GMHqsoP9TXU
KUteuv8LkPCsxUvefFdVC0iRRigVmEwmr+1HYJHkrawJMSeMUpo/IYM+BMwUWhyUR3P21gCn4IIHxCEK
DWFK+chq2IVxAuDTLGF8uOvD2MOjmVw2X6OYH1r4mvIoQwlKyyLUdfCbx/ljkCTt3PfGnc3KuOCh3zp5
AaUz8R2qXHCFP0um+6YWeSKN8dAWavKhHupba7TCWyx3GeyrgWk2zO5nUWBZnmejo0GiLiSH/91BtGNt
BWoxGO4TtuqZOpRqqFnNcDPcxjzvkEYnRRyjvGC/4wrgq/9//c1iImbR9uWmYlX3WU184Ke7eQoc3tfv
35/7syQvQOJHeFbPfCxQ6aDHmCqZDlNTOT6SD+9+IudUp/35f3gRDk0H5S2bbqt3JJPmtGsIN0LNu7ca
Bd2liG5cnV0dg2umiImOF1lmqT0R0U0wWBFhjBKaOfIyEwr9oQiLrcYnx8BZBtPwS0k/qE26LaAMJmIu
f4bhdjAump/uR7eFDaxz0xrWbZ7fr6xQVUO7DK5jsK0k+cA3VKqUZr5jsa/skQgx52jrzpS4GCK9ALBO
rU8TPxjabL/a82Gi16Rmq9sxcFrjN9OPhLisDfYD8iKKfM9i4/rI+NtbgEfzPGMhNTeNpQg16iNlU8ob
bvbbonZtSWz58RvEIyI+J67p20yv8FCJf0Vl115NssoUqJm6NMoeZvNuQ9foT5uSHanDWTbebK4d+nvX
JUfMnyomjAcTnrq+7IsjEu60m/q6bTNHiVKqXh1qupOmJfFLW2sWxupHclEiXK8XzGxy5wQwz0YlakjN
L79e3uhx4di3p22ahNwxU5eoDSpFk64Kl8omxakb9oOda/fR2n9CmzkGVnCr3KVEut4pUe2csYQ9P2rA
HNxtdTWDbnR1mVaDUrlq3hDU9ca2PJw2dDqPbasZn+32icIMB9eeTw6htq5HVNMFiLXx7fMjQ9TqNpc+
Ees9vpzJg1tyYqC+/g1vbxHbX9AcKqPwIQN0XCDnpUShg8/FEDdRbGDtioCZgj1w90YlnbdFoR/N2WW/
orzFa+2uYL18OWGcyps6YR7GjfdOkfv5/3b+v1iU7OmcNyr5j817BOqkO9/DnQnpphf89l/K5h0P2IPx
cXswbLz2XyYZt79DDu+R87eaIfrpSPtjUu+M/yMAAP//FMSc2uoaAAA=
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