package errcodes

import (
	"net/http"

	"github.com/Brainsoft-Raxat/tech-task/pkg/apperror"
)

var (
	NotFoundError       = apperror.NewErrorCode(1, http.StatusNotFound, "Resource not found")
	InternalServerError = apperror.NewErrorCode(2, http.StatusInternalServerError, "Internal server error")
	InvalidRequest      = apperror.NewErrorCode(3, http.StatusBadRequest, "Invalid request")
	Unauthorized        = apperror.NewErrorCode(4, http.StatusUnauthorized, "Unauthorized")
	Forbidden           = apperror.NewErrorCode(5, http.StatusForbidden, "Forbidden")
)
