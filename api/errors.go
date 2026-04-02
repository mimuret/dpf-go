package api

import "errors"

var (
	// ErrRecordNotFound はレコードが見つからない場合に返されます。
	ErrRecordNotFound = errors.New("record not found")
	// ErrZoneNotFound はゾーンが見つからない場合に返されます。
	ErrZoneNotFound = errors.New("zone not found")
)
