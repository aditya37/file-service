package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aditya37/file-service/endpoint"
	firebase_repo "github.com/aditya37/file-service/repository/firebase"
	mysql_repo "github.com/aditya37/file-service/repository/mysql"
	"github.com/aditya37/file-service/service"
	env "github.com/aditya37/get-env"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type Http interface {
	Start()
}

type httpServer struct {
	fileUploadTransport *kithttp.Server
	getUploadedFiles    *kithttp.Server
	detailFile          *kithttp.Server
}

func NewHttpServer() (Http, error) {

	// register firebase app
	firebaseApp, err := firebase_repo.NewFirebaseClient(firebase_repo.FirebaseClientConfig{
		BucketName:     env.GetString("FIREBASE_BUCKET_NAME", "upload-service-329303.appspot.com"),
		CredentialFile: env.GetString("FIREBASE_CRED_FILE_PATH", "firebase-admin-key.json"),
		ProjectID:      env.GetString("FIREBASE_PROJECT_ID", "upload-service-329303"),
	})
	if err != nil {
		return nil, err
	}

	// import repository for handle file storage
	firebaseRepo, err := firebase_repo.NewFirebaseStorage(*firebaseApp)
	if err != nil {
		return nil, err
	}

	// mysql repo
	db, err := mysql_repo.NewMysqlDataSource(mysql_repo.MysqlConfig{
		Host:              env.GetString("MYSQL_DBHOST", "localhost"),
		Port:              env.GetInt("MYSQL_DBPORT", 3306),
		Name:              env.GetString("MYSQL_DBNAME", "db_file_service"),
		User:              env.GetString("MYSQL_DBUSER", "root"),
		Password:          env.GetString("MYSQL_DBPASSWORD", "lym0us1n"),
		MaxConnection:     env.GetInt("MYSQL_MAX_CONNECTION", 10),
		MaxIdleConnection: env.GetInt("MYSQL_MAX_IDLE_CONNECTION", 10),
	})
	if err != nil {
		return nil, err
	}

	// service layer
	srv, err := service.NewFileService(firebaseRepo, db)
	if err != nil {
		return nil, err
	}

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	e := endpoint.NewFileServiceEndpoint(srv)
	return &httpServer{
		fileUploadTransport: kithttp.NewServer(
			e.FileUploadEndpoint,
			decodeRequestFileUpload,
			encodeFileUploadResponse,
			opts...),
		getUploadedFiles: kithttp.NewServer(
			e.UploadedFilesEndpoint,
			decodeRequestUploadedFile,
			encodeUploadedFileResponse,
			opts...,
		),
		detailFile: kithttp.NewServer(
			e.DetailFile,
			decodeDetailFileRequest,
			encodeDetailFile,
			opts...,
		),
	}, nil
}

func (g *httpServer) muxHandler() http.Handler {
	m := mux.NewRouter()
	m.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			next.ServeHTTP(w, r)
		})
	})
	m.Methods(http.MethodPost).Path("/file/upload").Handler(g.fileUploadTransport)
	m.Methods(http.MethodGet).Path("/files").Handler(g.getUploadedFiles)
	m.Methods(http.MethodGet).Path("/file/{object}").Handler(g.detailFile)
	return m
}

func (h *httpServer) Start() {
	errChan := make(chan error)
	// os signal
	go func() {
		chanSignal := make(chan os.Signal)
		signal.Notify(
			chanSignal,
			syscall.SIGINT,
			syscall.SIGALRM,
			syscall.SIGTERM,
		)
		errChan <- fmt.Errorf("%s", <-chanSignal)
	}()
	port := fmt.Sprintf(":%s", env.GetString("SERVICE_PORT", "4444"))
	go func() {
		serve := &http.Server{
			Addr:    port,
			Handler: h.muxHandler(),
		}
		log.Printf("Server run in port %s", port)
		if err := serve.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	log.Fatalf("Stop server with error detail: %v", <-errChan)
}
