package hstorage_common

import (
	"errors"
	"fmt"
)

type MyError struct {
	Err string `json:"error"`
}

func (m *MyError) Error() string {
	return fmt.Sprintf(m.Err)
}

func (m *MyError) Is(tgt error) bool {
	var target *MyError
	ok := errors.As(tgt, &target)
	if !ok {
		return false
	}
	return m.Err == target.Err
}
