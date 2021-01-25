package form

import "errors"

var (
	ErrUserCancelRequest = errors.New("form: user has cancelled the request")
	ErrNoSelectableItem  = errors.New("form: form must contain at least 1 selectable item")
)
