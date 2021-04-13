// Code generated by "esc -o data_gen.go -pkg main ../templates"; DO NOT EDIT.

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
		local:   "../templates/hrpc-client-go.htmpl",
		size:    5769,
		modtime: 1614909041,
		compressed: `
H4sIAAAAAAAC/+RYy27bOBfe5ylOhaSQCoXeG+midf6/LaZpgjpFF0VRMNKRTJgiFZJKJjD87gOSuka+
xG0zM8BkEQs8F57vXD5SWq3g+IZqhOlrCI1ixbzKMvYnhBnjWFKzeGuF5BMtMIKAlEoaGUSwXh+tVqDR
/J9xFLRACHWpmDAZBCd6cqLJQpUJSThDYUguAwhTVJgBuSwNk0KTd/KKJkuaY+QDcD6t02NmsLDhELdS
ei1YrWAY0siRU2dFKZWBQKCZLIwpg3YllzLnSHLJqciJVPnEgbmpMv/QaTI5YbIyjHdLWWF6nphZVDck
kcUkl4pxTif3eKNlssSelo2gUj0fNw8GdVBjdGvawhR4r9E06BUVOQKZo7pjiYcEANBJLtAsZNoKauGx
Qv1BuCIq1JLfYfpOXj+UWKeTAPkgysq4JZfqR8aXldlpfVmZgXnfnpYlinSOpsPl4yG9yuzXv6xM36DB
jSJtktN79Nk4XsZwfGcDbz01lk3WLb4lrNfBYw8sG6TZt/PlH2BU1SyM9ttUl2ON6q7t13bRDYXNaG8u
Zm4cullwU9Wl09jsdsbrNWijqsTAqs2eHyh4ZVubeHetzMaB6svnj9aMibyr0XukKSpwRv65E11fX13Z
9m+MGsHX+Wi5DjOrRAKf8L4faVgpXqtG8GqAoQteoamUgJd9cSft4E3t48sexNU6Hqi1SKdQKT6UeXzT
PtjH5i3mKQRWTQdDeQ19CsG97svW/Sz8VyeytmAZhEIaCKWCcEHr1p4bhbQAErm1uStTuxYNYbatFDZN
3e+LyG7SnxNnHCqv5bF8OHeK4asuY34tBlRKqgiG3eVowebcyW1yHfGTC6r0gvJQRWP1zKm+eA2C8Q3u
em0tGI8hKwz5n907C4NEVjwFm6SCKsMoB4W3FWozhZP7wMUw3nB9NHZ/28brmvoT3n/2jsLg6nJ+Hfht
555nwuBETyeTEz1pztJeCUnDVz6l67VVap6DuJ4/0g5Iu9JOXBSDO8d8FHa6QpfS6PfmztUPEoXUMJH/
XOIyqWAZg+NmP6oNPE+H2yK6rRW+Lb/Da7h7WolqG/ImTW3lhUFhTi2hBzEEtCw5S6i9r0zsvSiINrjQ
ZVvmOs7651yGCm+fI7+l1D+f3hQzVC5u8lamD2TGpcZwbJlSQ1tk/l5FbOe84TxsrZ8DnUKaenS6lELj
YfCko1cb88vH9LIaa9tQGzr5IoqaUDx07+k5EFbNTvzncdY7+CBju9PRZgN7H+IaHfNTkW5i+PpIGB8H
0TMxf7KgAg6h/5IKloRBJVhRcixQGEwfTeOBiDeA/WWsEYRMgAV3djqC7RoKzk53YN+Ov7INXSlOvnz+
uJonCyxw2tBNfe2J4b20XPCY+WO4omYxheDwg6V3dWhvejH8aFmhfXsi55jRiptzRjkqYn/CiszdvTKM
4iGB//pEuX+o1IYx2XBxmNlQC7rEYdu17biJQrbYtOWKxhvl0jVMGG0B4Hk3afl2o1Khcz3c+tt3e25v
Ud+3Z3Ocbpfavx8xFKg1zbE7yRzVX/jVDYfDodUbvizYBFik0V7dG4V0uVNrvVPq8nl22uA7OszLeluZ
dudUI8fB298oA1SjDS0GubTZPju1cU73ZfmFXD45vUzMoieq2nbfr+vnb08pdorNwp52/stJOB6oXaab
j2mXQef19zbor0N1DHJ26oPb3QYaRdr1AROz39oGf1dtB5fF4duZBfhvq4/vp4R8Vcxgw3LdafaWCaoe
LhpOtOD+AQQHctXRk9irPkaZmLn7yGzfzVGke77p/RUAAP//+4OCk4kWAAA=
`,
	},

	"/templates/hrpc-scanner.htmpl": {
		name:    "hrpc-scanner.htmpl",
		local:   "../templates/hrpc-scanner.htmpl",
		size:    591,
		modtime: 1618289015,
		compressed: `
H4sIAAAAAAAC/3SQQW/UMBCFz+tfMURCslHkvVPtBQk4oFIkChfEYZKMXauOHcYOSmXlvyMnu10E6s2e
9/zeNy4FEuUPzlPAkUDa+HzW0KQeQyBuFKyrEBP2j2gJSgHjPE2YH95hItB3U3YxJP0xfjlbqt2NU+QM
zYAZO0x0TL/8cWD3m7h5Fm2M1pO20WOwOrI9Thxz7GazH65OM+ZGiFLAGdC3lBJaun+atq59i7tPkHm+
DCgMG0YpwBgs/ffIzKEHucCbUkB/rhuvq4Lv6GeSCuROqrd7C8QcWUERB6Y8c4CNTt8ipwf0clFiFS8k
fu0xyMQ9uJCJDfZUVrUH1jxnoIqnEwTn6+DSEJwXh3UzdC3ER3h7qk4tf/zsnjKpmzqrfmdqWpV3qG9h
PGN1LSzqZlNfXfMvBcQsDrXh38bz1YxZv6+YRjZzoGWiPtMAuf7f6/umrTR1778++08AAAD//6ZjmSRP
AgAA
`,
	},

	"/templates/hrpc-server-echo-go.htmpl": {
		name:    "hrpc-server-echo-go.htmpl",
		local:   "../templates/hrpc-server-echo-go.htmpl",
		size:    9796,
		modtime: 1618288944,
		compressed: `
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
uHRsbr/yVWNeAAf+sOZfAQAA///04BkQRCYAAA==
`,
	},

	"/templates/hrpc-server-go.htmpl": {
		name:    "hrpc-server-go.htmpl",
		local:   "../templates/hrpc-server-go.htmpl",
		size:    8256,
		modtime: 1613010077,
		compressed: `
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
r5pu9Dt+ovPu+J8CAAD//6+DjS9AIAAA
`,
	},

	"/templates": {
		name:  "templates",
		local: `../templates`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"../templates": {
		_escData["/templates/hrpc-client-go.htmpl"],
		_escData["/templates/hrpc-scanner.htmpl"],
		_escData["/templates/hrpc-server-echo-go.htmpl"],
		_escData["/templates/hrpc-server-go.htmpl"],
	},
}
