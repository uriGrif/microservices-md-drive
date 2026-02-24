package objectstorage

import "io"

type ObjectStorage interface {
	CreateObject(key string, data []byte) error
	GetObject(key string) (io.Reader, error)
	GetObjectUrl(key string, expirationSeconds int64) (string, error)
	UpdateObject(key string, data []byte) error
	DeleteObject(key string) error
}
