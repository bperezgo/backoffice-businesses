package handler

import "github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"

type BasisHandler struct {
	HandlerMethod HandlerMethod
	Path          string
	BasisBody     interface{}
}

func (h *BasisHandler) GetMethod() HandlerMethod {
	return h.HandlerMethod
}

func (h *BasisHandler) GetPath() string {
	return h.Path
}

func (h *BasisHandler) GetEmptyRequest() handlertypes.Request {
	return handlertypes.Request{
		Body: h.BasisBody,
	}
}
