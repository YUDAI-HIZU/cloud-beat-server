package usecase

import (
	"app/domain/models"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

func resizeImage(i *models.Image) error {
	img, t, err := image.Decode(i.Buf)

	if err != nil {
		return err
	}

	switch t {
	case "jpeg":
		if err := jpeg.Encode(i.Buf, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(i.Buf, img); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(i.Buf, img, nil); err != nil {
			return err
		}
	}
	return nil
}
