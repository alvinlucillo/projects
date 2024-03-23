package handlers

import (
	"fmt"
	"healthstats/pkg/services"
	"net/http"
)

const (
	packageName = "handlers"
)

type handler struct {
	service *services.Service
}

func NewFileHandler(service *services.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UploadFile(w http.ResponseWriter, req *http.Request) {
	l := h.service.Logger.With().Str("package", packageName).Str("function", "UploadFile").Logger()
	err := req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		l.Err(err).Msg("Error parsing multipart form")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the file from form data.
	file, handler, err := req.FormFile("file") // "file" is the key of the form data
	if err != nil {
		// fmt.Printf("Error retrieving the file: %s", err.Error())
		l.Err(err).Msg("Error retrieving the file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	result, err := h.service.S3Service.UploadFile(handler.Filename, file)
	if err != nil {
		l.Err(err).Msg("Error uploading file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%#v\n", result)

	// You can now use the file, for example, save it to disk.
	// For now, let's just respond with the name of the file.
	w.Write([]byte(fmt.Sprintf("Successfully uploaded file: %s", handler.Filename)))
}

func (h *handler) InitRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /upload", h.UploadFile)
}
