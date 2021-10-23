package http

import (
	"net/http"
	"path"
)

type fileServer struct {
	prefix string
	fs     http.FileSystem
}

func PrefixFileSystem(prefix string, fs http.FileSystem) http.FileSystem {
	return &fileServer{
		prefix: prefix,
		fs:     fs,
	}
}

// PrefixFileServer return a Handler
func PrefixFileServer(prefix string, fs http.FileSystem) http.Handler {
	return http.FileServer(PrefixFileSystem(prefix, fs))
}

func (f *fileServer) Open(name string) (http.File, error) {
	name = path.Join(f.prefix, name)
	return f.fs.Open(name)
}
