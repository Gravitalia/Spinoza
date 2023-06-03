package uploader

import (
	"bytes"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gravitalia/spinoza/helpers"
)

var (
	ctx = context.Background()
)

func UploadOnCloudinary(image []byte) (string, error) {
	cld, _ := cloudinary.New()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		bytes.NewReader(image),
		uploader.UploadParams{PublicID: helpers.GetHash(image)})

	if err != nil {
		return "", err
	}

	return uploadResult.PublicID, nil
}
