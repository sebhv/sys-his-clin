package apperrors

import "errors"

var (
	ErrNotFound      = errors.New("recurso no encontrado")
	ErrAlreadyExists = errors.New("el recurso ya existe")
	ErrInvalidInput  = errors.New("datos de entrada inválidos")
	ErrUnauthorized  = errors.New("no autorizado")
	ErrForbidden     = errors.New("acceso denegado")
)
