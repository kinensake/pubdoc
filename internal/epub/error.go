package epub

import "errors"

var (
	ErrInvalid  = errors.New("epub: invalid file")
	ErrNotFound = errors.New("epub: file not found")
)
