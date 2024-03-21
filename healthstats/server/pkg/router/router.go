package router

import (
	"fmt"
	"net/http"
)

type router struct {
	router *http.ServeMux
}

func NewRouter() *http.ServeMux {
	r := &router{router: http.NewServeMux()}

	r.setupRoutes()

	return r.router
}

func (r *router) setupRoutes() {
	r.router.HandleFunc("POST /upload", func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve the file from form data.
		file, handler, err := req.FormFile("file") // "file" is the key of the form data
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// You can now use the file, for example, save it to disk.
		// For now, let's just respond with the name of the file.
		w.Write([]byte(fmt.Sprintf("Successfully uploaded file: %s", handler.Filename)))
	})

}
