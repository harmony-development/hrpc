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
		size:    8264,
		modtime: 1612055902,
		compressed: `
H4sIAAAAAAAC/+xZW2/cuhF+31/BLNxCMmRuW/TJjR/ipE2CNrHhC/oQBAe0NCsRK5EySXnjs9B/PyAp
aqnbru3jE5/bvnhNDsmZb775OLQ3GyRB/YfmwEgBKEh5+x2juQRxB2IeorqezTYbdEAVFOj4BGEzUpJ4
RVJAmw1a0hxKorJTIgHhs1JRziR+z88bE21Oi5ILheYxZwq+qXk7wEAtMqXK7QjlC8orRfPtUMp5mgNO
eU5YirlIF6Xgit9US/vFs6Qqq25wzItFygXNc7JYw43k8QrUg/ZT9yXIRQIyFrRUXJQ3o5tnRBSc3R8l
cAc5LwtgapGJMl40sDWQmYVSo8ZgLUE5MAVhKSB8CeKOxhYhhBDaznwClfGknWgmDwTIj0xvFwiQPL+D
5D2/ui+hyQ5G+CMrK2WGTOZ6i88qtXP1WaU6y/31pCyBJZegtnFZf7CX6P32Z5XyF7i4gSUOHO+rReNg
FaGDO+14u5Nb6VKj41uhup53drgjwhCUsJhC8lkz24Za1+jQzzHWzH/XDpxrNiCTtGCfWTibLSsWI8qo
CkK0MW4lRBHtb1Md7/Sv7mTrOAihDQzv8DUriJAZyQO9MJryObQhL83iVyeI0bw5T39KwmgcgBDWrJ5Z
COmywzNb9Gf/RUpUbmAA+BgxDzSz2/pvB41caErJUlCmlmj+F3lpawAFCQhYIqyDCLd80iXmLa5rRJkC
sSQxNOH8quuALlHAuEIBFyjIiHybU2DqUgkgBcKhGbMItGNh14lmIw8dVNdBrL6hRh7xW/szQgIdtvHi
j+/qOkIZkASERFo18QfzS2giKVtbXWPWWDMFhOANKVyx5RJMHIQlY/42AQ6DC58xDl4pFGeE9bxGUzEO
AniaJ5R1T30ef1gyIma2DCcZ7dY3KtWWwyAQIx9dvbK7PECxRg3DLaGH2uU+noYVdhM9gDvl8BAl2x1X
2NlrStsmNM59tljWs5GMOMHpnv6BsCQHgaQSVay8w2wxdFxv5/6ta8mt1NgFbYVFaG35cQGy5EzC/wVV
Pk+uGRH35wKQ7RFws82VIEwuuShAtKZVmQpNMtQ2L/i6GfJZZZL3GdZTsQWyE4WRg8c5H5pqGAdui5gA
VQmG/voA0y3Ax0hGnWEf22PP1a6Vg+Z4BJshYS6AJKfVcgnikv4Ixwj9/W//+Gc0MDPR+nZDszrqkczP
QZAdThIsRJegXO4DOZ19vwAz3LLlBMmR86bzEqILXimQQYi+fJVKUJYOU+VmNn0JnVQp95kv3GPAa+Ow
axCcL9rIfZ9H/UM8fZzCc1d8hj8frq7Og1HSRkjALTpsZm4rkMqHVq6pijN9Y97i64v/4XOiMn/+d958
xPqV9oQkHveK+IYn95ETf/tgw7ra3uS5gfaUJ/ddlU5gCQK5Ofw25xKCx+l/hn2R0PIVoXU4MLMk710R
fRI+PY/2CEMse+/a913gdxRtu+5fk8Nb0qLob/ZCgFB2x1dgkmluhnikh4Pbxv9PICVJYaoX7RnZ/nPE
+UaLMmwZiz26BbE5DweHfVDbQ8Od4Sx9CZ2Ez0V94lkHe3uWyQda5Dbc6Ztu1beVYxeYkmmQDsIuIywS
ruscbKWFL+p2YZ8adunpF+LTunE4CPGbJAnmJjamjnQ1zSM0J2WZ05goytmCxwrUkTSCNe8e9kPUFM4a
G3EPXMTh7LvFNfw28gJ5LlnVj4G2Kxtolpb/EdXvaRM1qlaQFQTDp86EMDGaD27pkUfWb1v1LTA/S6op
C/d0M981Im57iWGu28drr1DWsk1fhl0v7RroYG20JtJev1CKUm7vn7H74sEFoD+FTGUXmi9fb+5VXzj2
nWlaUi4mZhqJKryrTp+5lqYomhswCCfX7oPV/8SmcnRY4U67GwFkNWlRT84YwF4fuWBmD1tdj0TXe34P
1WAtrZo7gLYvOSMPbedgM7apR3I2nRMJOXTe80+mUKvr9m8ZfKVz+/pIA3W8K6Wv+GpPLkfqYEdNdLZv
/guzV8T2C5qNSm/4nATtC+S4Fa9U+EshxDSLdVhTDBgR7E66C5lus80r9WLJXvuK8hm+KfvA9erllOqu
tSmY50njo0vkcfnfjf93Y8mezrmQ6Z9oPoKog+58D3aa0q4X/NcfFM0HXrCz/nU76zZeI4/n3mOSMvPf
je47cvxV041+x5/qvDv+pwAAAP//UoD+gUggAAA=
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
