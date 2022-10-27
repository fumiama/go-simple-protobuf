package spb

import "errors"

var (
	ErrInvalidStructLen = errors.New("invalid struct_len") // 1B<struct_len<1MB
	ErrInvalidDataLen   = errors.New("invalid data_len")
)
