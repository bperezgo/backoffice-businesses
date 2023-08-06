package handler

import "github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"

type BasisHandler[T any] struct {
	HandlerMethod HandlerMethod
	Path          string
	BasisBody     T
}

func (h *BasisHandler[T]) GetMethod() HandlerMethod {
	return h.HandlerMethod
}

func (h *BasisHandler[T]) GetPath() string {
	return h.Path
}

func (h *BasisHandler[T]) GetEmptyRequest() handlertypes.Request {
	return handlertypes.Request{
		Body: []byte{},
	}
}
