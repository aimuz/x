package http

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type spaHandler struct {
	index string
	fs    http.FileSystem
}

func SPAServer(index string, fs http.FileSystem) http.Handler {
	return &spaHandler{
		index: index,
		fs:    fs,
	}
}

const indexPage = "/index.html"

func (s *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check whether a file exists at the given path
	_, err = s.fs.Open(path)
	if os.IsNotExist(err) {
		if !strings.HasPrefix(s.index, "/") {
			s.index = "/" + s.index
		}

		r.URL.Path = s.index
		if indexPage == s.index {
			r.URL.Path = "/"
		}

		http.FileServer(s.fs).ServeHTTP(w, r)
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(s.fs).ServeHTTP(w, r)
}
