package helpers

import (
	"bytes"
	"image/jpeg"
	"os"
	"strconv"

	"github.com/chai2010/webp"
	"github.com/discord/lilliput"
)

func Compress(image []byte, width int32, height int32) ([]byte, error) {
	var outputWidth, _ = strconv.Atoi(os.Getenv("DEFAULT_OUTPUT_WIDTH"))
	var outputHeight, _ = strconv.Atoi(os.Getenv("DEFAULT_OUTPUT_HEIGHT"))

	decoder, err := lilliput.NewDecoder(image)
	// this error reflects very basic checks,
	// mostly just for the magic bytes of the file to match known image formats
	if err != nil {
		return []byte{}, err
	}
	defer decoder.Close()

	header, err := decoder.Header()
	// this error is much more comprehensive and reflects
	// format errors
	if err != nil {
		return []byte{}, err
	}

	// get ready to resize image,
	// using 8192x8192 maximum resize buffer size
	ops := lilliput.NewImageOps(8192)
	defer ops.Close()

	// create a buffer to store the output image, 10MB in this case
	outputImg := make([]byte, 10*1024*1024)

	if width != 0 {
		outputWidth = int(width)
	}

	if height != 0 {
		outputHeight = int(height)
	}

	var outputPixels = outputWidth * outputHeight

	// Calcul new width and height
	if header.Width()*header.Height() > outputPixels {
		outputWidth, _ = strconv.Atoi(os.Getenv("DEFAULT_OUTPUT_WIDTH"))
		outputHeight = int(float64(outputWidth) / float64(header.Width()) * float64(header.Height()))
	} else {
		outputWidth = header.Width()
		outputHeight = header.Height()
	}

	opts := &lilliput.ImageOptions{
		FileType:             ".jpeg",
		Width:                outputWidth,
		Height:               int(outputHeight),
		ResizeMethod:         lilliput.ImageOpsResize,
		NormalizeOrientation: true,
		EncodeOptions:        map[int]int{lilliput.JpegQuality: 90, lilliput.JpegProgressive: 100},
	}

	// resize and transcode image
	outputImg, err = ops.Transform(decoder, opts, outputImg)
	if err != nil {
		return []byte{}, err
	}

	resizedImage, err := jpeg.Decode(bytes.NewReader(outputImg))
	if err != nil {
		return nil, err
	}

	// Convert JPEG to WebP
	webpImg := new(bytes.Buffer)
	err = webp.Encode(webpImg, resizedImage, &webp.Options{Lossless: false, Quality: 90})
	if err != nil {
		return nil, err
	}

	// Read the encoded WebP image
	encodedImg := webpImg.Bytes()

	return encodedImg, nil
}
