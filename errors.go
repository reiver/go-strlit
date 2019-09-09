package strlit

import (
	"errors"
)

var (
	errEmptySource    = errors.New("strlit: Empty Source")
	errNilDestination = errors.New("strlit: Nil Destination")
	errNilSource      = errors.New("strlit: Nil Source")
	errUTF8RuneError  = errors.New("strlit: UTF-8 Rune Error")
)
