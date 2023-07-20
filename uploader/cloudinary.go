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

// UploadOnCloudinary uploads image into Cloudinary provider
// and then, returns a hash of the image
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

// DeleteOnCloudinary allows to remove a picture on
// Cloudinary provider
func DeleteOnCloudinary(hash string) (string, error) {
	cld, _ := cloudinary.New()

	uploadResult, err := cld.Upload.Destroy(
		ctx,
		uploader.DestroyParams{PublicID: hash},
	)

	if err != nil {
		return "", err
	}

	return uploadResult.Error.Message, nil
}
