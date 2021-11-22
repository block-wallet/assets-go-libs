package validation

import (
	"fmt"
	"path/filepath"

	"github.com/trustwallet/assets-go-libs/pkg"
)

func ValidateLowercase(name string) error {
	if !pkg.IsLowerCase(name) {
		return fmt.Errorf("%w: file - %s should be in lowercase", ErrInvalidFileCase, name)
	}

	return nil
}

func ValidateExtension(name, ext string) error {
	fileExtension := filepath.Ext(name)
	if fileExtension != ext {
		return fmt.Errorf("%w: %s, only .png allowed in folder", ErrInvalidFileExt, name)
	}

	return nil
}
