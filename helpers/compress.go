package helpers

import (
	"os"
	"strconv"

	"github.com/discord/lilliput"
)

func Compress(image []byte, width int32, height int32) ([]byte, error) {
	var outputWidth, _ = strconv.Atoi(os.Getenv("DEFAULT_OUTPUT_WIDTH"))

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
	} else if outputWidth > header.Width() {
		outputWidth = header.Width()
	}

	resizeMethod := lilliput.ImageOpsResize
	if outputWidth == header.Width() {
		resizeMethod = lilliput.ImageOpsNoResize
	}

	// Calcul what height it needs to be
	outputHeight := float64(outputWidth) / float64(header.Width()) * float64(header.Height())
	if outputHeight == 0 {
		defaultOutputHeight, _ := strconv.Atoi(os.Getenv("DEFAULT_OUTPUT_HEIGHT"))
		outputHeight = float64(defaultOutputHeight)
	}

	if height != 0 {
		outputHeight = float64(height)
	}

	opts := &lilliput.ImageOptions{
		FileType:             ".webp",
		Width:                outputWidth,
		Height:               int(outputHeight),
		ResizeMethod:         resizeMethod,
		NormalizeOrientation: true,
		EncodeOptions:        map[int]int{lilliput.WebpQuality: 100},
	}

	// resize and transcode image
	outputImg, err = ops.Transform(decoder, opts, outputImg)
	if err != nil {
		return []byte{}, err
	}

	return outputImg, nil
}
