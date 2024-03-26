package handler

import (
	"context"
	"fmt"
	"healthstats/pkg/middleware"
	"healthstats/pkg/model"
	"healthstats/pkg/repo"
	"healthstats/pkg/service"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	packageName = "handler"
)

type handler struct {
	service *service.Service
}

func NewFileHandler(service *service.Service) *handler {
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

	tx := req.Context().Value("tx").(*sqlx.Tx)
	rqRepo := repo.NewRequestRepo(tx, h.service.Logger)

	newRequest := model.Request{
		FileName: handler.Filename,
		Status:   model.RequestStatusPending,
	}

	id, err := rqRepo.CreateRequest(newRequest)
	if err != nil {
		l.Err(err).Msg("Error creating request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("ID: %s\n", id)

	go func() {
		newRequest.ID = id
		newRequest.Status = model.RequestStatusSuccess

		// time.Sleep(5 * time.Second)

		_, err := h.service.S3Service.UploadFile(handler.Filename, file)
		if err != nil {
			l.Err(err).Msg("Error uploading file")

			newRequest.Status = model.RequestStatusFailed

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("File uploaded")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		tx, err := h.service.DB.BeginTxx(ctx, nil)
		if err != nil {
			l.Err(err).Msg("Error starting transaction")
			return
		}

		rqRepo := repo.NewRequestRepo(tx, h.service.Logger)

		err = rqRepo.UpdateRequest(newRequest)
		if err != nil {
			l.Err(err).Msg("Error updating request")
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tx.Commit()
		if err != nil {
			l.Err(err).Msg("Error committing transaction")
			return
		}

		// fmt.Printf("%#v\n", result)
	}()

	// You can now use the file, for example, save it to disk.
	// For now, let's just respond with the name of the file.
	fmt.Println("Waiting for file to be uploaded. For now, sending response")
	w.Write([]byte(id))
}

func (h *handler) InitRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /upload", middleware.TransactionMiddleware(h.service, h.UploadFile))
}
