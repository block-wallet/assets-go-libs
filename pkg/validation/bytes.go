package validation

import (
	"encoding/json"
	"fmt"
	"image/png"

	"github.com/trustwallet/assets-go-libs/pkg/file"
)

const kbInByte = 1000

func ValidateImageDimension(file *file.AssetFile, maxW int, maxH int, minW int, minH int) error {
	img, err := png.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	if img.Width > maxW || img.Height > maxH || img.Height < minH || img.Width < minW {
		return fmt.Errorf("%w: %s, max - %dx%d, min - %dx%d; given %dx%d",
			ErrInvalidImgDimension, file.Name(), maxW, maxH, minW, minH, img.Width, img.Height)
	}

	return nil
}

func ValidateSize(file *file.AssetFile, limit int) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	logoSize := fileInfo.Size() / kbInByte

	if logoSize > int64(limit) {
		return fmt.Errorf("%w: logo should be less than %dKB instead of %dKB", ErrInvalidFileSize, limit, logoSize)
	}

	return nil
}

func ValidateJson(b []byte) error {
	if !json.Valid(b) {
		return ErrInvalidJson
	}

	return nil
}
