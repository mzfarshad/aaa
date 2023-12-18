package errs

import (
	"fmt"
)

func NewNotFound(entity string) error {
	return fmt.Errorf("%s not found", entity)
}
