package utils

import "errors"

var (
	ErrConfigNotFound         = errors.New("config file not found")
	ErrInvalidCopyDestination = errors.New("copy destination is invalid")
	ErrInvalidCopyFrom        = errors.New("copy from is invalid")
	ErrMapKeyNotMatch         = errors.New("map's key type doesn't match")
	ErrNotSupported           = errors.New("not supported")
)
