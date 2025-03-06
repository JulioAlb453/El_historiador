package domain

import (
	"errors"
	"fmt"
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("campo: '%s' error:%s", e.Field, e.Message)
}

var (
	ErrMissingField           = errors.New("Faltan campos obligatorios")
	ErrInvalidData            = errors.New("Datos invalidos")
	ErrInternalError          = errors.New("Ocurrio un error interno al procesar la solicitud")
	ErrMissingTitle           = errors.New("El titulo es requerido")
	ErrMissingAuthor          = errors.New("El nombre del autor es requerido")
	ErrMissingContent         = errors.New("El contenido es requerido")
	ErrMissingDescription     = errors.New("La descripcion es requerida")
	ErrMissingPublicationDate = errors.New("La fecha de publicacion es requerida")
	ErrMissingTopic           = errors.New("El tema es requerido")
	ErrInvalidPublicationDate = errors.New("La fecha de publicacion es invalida ")
)
