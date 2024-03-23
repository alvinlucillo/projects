package router

import (
	"healthstats/pkg/handlers"
	"healthstats/pkg/services"
	"net/http"
)

type router struct {
	router  *http.ServeMux
	service *services.Service
}

func NewRouter(service *services.Service) *http.ServeMux {
	r := &router{router: http.NewServeMux(), service: service}

	r.setupRoutes()

	return r.router
}

func (r *router) setupRoutes() {

	fileHandler := handlers.NewFileHandler(r.service)
	fileHandler.InitRoutes(r.router)

	// r.router.HandleFunc("POST /upload", func(w http.ResponseWriter, req *http.Request) {
	// 	err := req.ParseMultipartForm(10 << 20) // 10 MB
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	// Retrieve the file from form data.
	// 	file, handler, err := req.FormFile("file") // "file" is the key of the form data
	// 	if err != nil {
	// 		fmt.Printf("Error retrieving the file: %s", err.Error())
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	defer file.Close()

	// 	// // Create a new session in the us-west-2 region.
	// 	// sess, err := session.NewSession(&aws.Config{
	// 	// 	Region: aws.String("us-east-1")},
	// 	// )
	// 	// if err != nil {
	// 	// 	fmt.Printf("Error creating session: %s", err.Error())
	// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// 	return
	// 	// }

	// 	// Assume the role
	// 	// stsSvc := sts.New(sess)
	// 	// creds := stscreds.NewCredentials(sess, "arn:aws:iam::488035000657:role/healthstats-terraform")

	// 	// Create an uploader with the session and default options
	// 	// uploader := s3manager.NewUploader(sess)

	// 	// Create an uploader with the session and assumed role
	// 	// uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
	// 	// 	u.PartSize = 5 * 1024 * 1024 // 5MB part size
	// 	// 	u.LeavePartsOnError = true   // Don't delete the parts if the upload fails.
	// 	// 	u.Concurrency = 3            // Download parts concurrently.
	// 	// 	u.S3 = s3.New(sess, &aws.Config{Credentials: creds})
	// 	// })

	// 	// Upload the file to S3.
	// 	// result, err := uploader.Upload(&s3manager.UploadInput{
	// 	// 	Bucket: aws.String("healthstats-files"),
	// 	// 	Key:    aws.String(handler.Filename),
	// 	// 	Body:   file,
	// 	// })
	// 	// if err != nil {
	// 	// 	fmt.Printf("Error uploading file: %s", err.Error())
	// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// 	return
	// 	// }

	// 	result, err := r.service.S3Service.UploadFile(handler.Filename, file)

	// 	fmt.Printf("%#v\n", result)

	// 	// You can now use the file, for example, save it to disk.
	// 	// For now, let's just respond with the name of the file.
	// 	w.Write([]byte(fmt.Sprintf("Successfully uploaded file: %s", handler.Filename)))
	// })

}
