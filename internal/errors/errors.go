package errors

import "errors"

var (
	ErrConfigNotFound = errors.New("config not found")
	ErrNoKeyFound     = errors.New("no ssh key found")
	ErrWrongTaskBody  = errors.New("wrong task body")
	ErrDataBaseBak    = errors.New("database backup error")
	ErrMediaBak       = errors.New("media backup error")
	ErrFileNotFound   = errors.New("source file not found")
)
