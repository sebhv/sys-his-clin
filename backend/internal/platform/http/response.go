package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/sebhv/sys-his-clin/backend/internal/shared/apperrors"
)

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	status, msg := statusFromErr(err)

	if status == http.StatusInternalServerError {
		slog.Error("error no manejado", "err", err, "path", r.URL.Path, "method", r.Method)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func statusFromErr(err error) (int, string) {
	switch {
	case errors.Is(err, apperrors.ErrNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, apperrors.ErrAlreadyExists):
		return http.StatusConflict, err.Error()
	case errors.Is(err, apperrors.ErrInvalidInput):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, apperrors.ErrUnauthorized):
		return http.StatusUnauthorized, err.Error()
	case errors.Is(err, apperrors.ErrForbidden):
		return http.StatusForbidden, err.Error()
	default:
		return http.StatusInternalServerError, "error interno"
	}
}
