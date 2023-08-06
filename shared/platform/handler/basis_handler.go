package handler

import (
	"encoding/json"
	"errors"

	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
)

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

func (h *BasisHandler[T]) ValidateAndDecode(body interface{}, req []byte) error {
	err := json.Unmarshal(req, body)

	if err != nil {
		return errors.New("Invalid request body")
	}

	err = Validator(body)
	if err != nil {
		return err
	}

	return nil
}
