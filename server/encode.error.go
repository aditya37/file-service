package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aditya37/file-service/service"
	"github.com/aditya37/file-service/utils"
)

type Errors interface {
	Errs() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		log.Panic("Error not nil")
	}
	switch e := err.(type) {
	case *utils.CustomError:
		switch e.Code {
		case service.ErrCodeWrongFileFormat:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeFailedWriteTemp:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeReadFile:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeFailedCreateTemp:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeFileLarge:
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeFailedParseObject:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeFailedUploadToServer:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(e)
		case service.ErrCodeWrongRequest:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.CustomError{
				InternalError: e.InternalError,
				Description:   e.Description,
				Code:          e.Code,
			})
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CustomError{
			InternalError: e.Error(),
			Description:   e.Error(),
			Code:          0,
		})
	}
}
