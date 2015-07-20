package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
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

type _escDir struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	local      string
	isDir      bool

	data []byte
	once sync.Once
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

func (dir _escDir) Open(name string) (http.File, error) {
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
	return nil, nil
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
	return time.Time{}
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
		return _escDir{fs: _escLocal, name: name}
	}
	return _escDir{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(f)
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

	"/template/html/HandshakeProtocolDetails.html": {
		local: "template/html/HandshakeProtocolDetails.html",
		size:  3188,
		compressed: `
H4sIAAAJbogA/7RXzW7bRhC+6ykmmyBpg5CsGwRtXUqAYTtIgQAJLOXQXooVuSZXXi6Z3aUbQdCxtz5C
+3J9ku4PKf6KVg3kYu/P/Hwz881yFD65+nC5+vXjNaQqY4vZbhe8hFUOZRFjRUClVIIiWcH07hWIksML
IiPwclBMxiKo7/wkP8i9gJfBfh8ag8AwT+aIcLSYhSnB8WIGEGZEYYhSLCRRc/Rp9db7EdkLRRUji9X7
JbzDPJYpviNwpYUpk2HgLo3YE8+D99qRVBDlWUEZiUHLQ0Y5vaV6c7lcgudZWUb5HQjC5kiqLSMyJUQh
SAW5naNUqUKeB0GGv0Qx99d5rqQSuDAbbTg4HASv/df+myCSsjnztTdfn6ADpA+FojnHTGeNZOSrA/Cs
mxaMMHAZDtd5vNX/YnoPEcNSzlGUc51FToRFq1GtRVCtaJaAFJEDo7FQX5+Uwvr/7dPFTx9vrv1NkSDA
TBcLBQMD1SpUeM1I7dBt7F8vze8bx6pmgduJemk28eKpLnPcPVqKCN7lUg1vLiKVi+HxleHFUQ1doREf
CqtSDs/fYsutG4JlV02vK+R65SKyu90OhGY8gWeUx+TLK3iWiLws4HwOPuz3lXYjJcv1L5WgLhEnFp+R
rvQOKlaJ3gL53CjBd819qEzRQTct9lSeJIyYojOGC0lQdYxFYvrtqZH8fbdzEPd7dCAJo9GdqRhqstD2
2yD0l2UUEamfBlGSNshDWWub0gn2LRImp/Vikx8xUNNNPtBypdrtGOFV2vTj0yllV64Th4gMU6YUvFNt
Gd79b2OPyK6z9O/ff47YGs9rpfHPX2MaUyltg3K94FphPMROU7g3qDZmyUnjOTrKvZqrPXh1N+kGIfeE
K9sbLVTX5lBeCIG3wzCa5Frd01lL+W2O+uE9krYP5/k4XcZv+mFdfj9NlecCC/HzUVtTnHnOplUfCmp2
hFg1dFu+1bYwACqN6Vgf2RonhDrSHt0oT66Q64PRNht0fqLgG/twtTh9Q/S3N9Ma2Gwdub9tP/YtE1V/
CBqbb4joaPZ7ZczueDYOfP4DC055MiD0VPhNcKZsBpnFbuV7EM6tfqsyjdEJelnXevJjssB8jt4g81J1
Qz/6BgftiaNfne7ZCFfGMtPB8gNy6Pt+RhjQd9vdd3et91QvzSda8ysM9ICnxzwZCVqoZoYzA2WUx8Tf
fC6J2NpBzi29M//sTA+SZmTcSAPV6R4xcupUuulPxT3bQTWPBu4nxuy/AAAA///LAF1jdAwAAA==
`,
	},

	"/template/txt/HandshakeProtocolDetails.txt": {
		local: "template/txt/HandshakeProtocolDetails.txt",
		size:  1046,
		compressed: `
H4sIAAAJbogA/6xTTW/iMBS88ytmJSR2EQmrPSIWCfVD7ZVyq3pIndc0UrBT26mKLP/3Oo6Lk8ARn5LJ
vHlv5jnGLOfYCzR1nmmCfi8VNB3qyr0tIBuOGSmGREBXKpfLn29pIU68GeZLa42BzHhBmJY8p68FpoUU
TY3Vf6SwNrnemcRWqnl9DN2Y4JyYLgVvW4bmfqzyDfQRufjr4GdjKuKBZu0LoCRbwZieTvok2YNQ2loA
OSk9/n7rsI4Qu/TLG8ZIuTxlQ/AiAXFsqlTA7rOyaiS1GM87aD1s0zIo31GmBLd20wva2aZP4to77lXc
taDaSpkdneIEwxOH9cXpzb/TjEmyicOtkyRO5Z4C3YvvjzUFeKA18nxmeOz29DDplAqN334xPTc7YuJw
cLSsfe1s/fFbjFHIMm8vgRxQx7FcEhpckVbFC/uckhHfr3/YoZs7LvPSuVCEc/8jmW1VQYUsfw1CGtVd
+b/qRL8DAAD//2SFvx0WBAAA
`,
	},

	"/": {
		isDir: true,
		local: "/",
	},

	"/template": {
		isDir: true,
		local: "/template",
	},

	"/template/html": {
		isDir: true,
		local: "/template/html",
	},

	"/template/txt": {
		isDir: true,
		local: "/template/txt",
	},
}
