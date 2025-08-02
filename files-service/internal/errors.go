package internal

import "errors"

var ErrNotAuthorized error = errors.New("user does not have permission to perform this action")

var ErrFileNotFound error = errors.New("file not found")
