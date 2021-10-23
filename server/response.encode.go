package server

import (
	"context"
	"encoding/json"
	"net/http"
)

// encodeFileUploadResponse
func encodeFileUploadResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if e, ok := response.(Errors); ok && e.Errs() != nil {
		encodeError(ctx, e.Errs(), w)
	}
	return json.NewEncoder(w).Encode(response)
}
