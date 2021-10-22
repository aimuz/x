package http

import (
	"net/http"
	"path"
)

type FileServer struct {
	prefix string
	fs     http.FileSystem
}

// PrefixFileServer return a Handler
func PrefixFileServer(prefix string, fs http.FileSystem) http.Handler {
	return http.FileServer(&FileServer{
		prefix: prefix,
		fs:     fs,
	})
}

func (f *FileServer) Open(name string) (http.File, error) {
	name = path.Join(f.prefix, name)
	return f.fs.Open(name)
}

type SPAServer struct {
}
