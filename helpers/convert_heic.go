package helpers

import (
	"fmt"
	"os"

	vips "github.com/davidbyttow/govips/v2"
)

func ConvertHEIC(image []byte) ([]byte, error) {
	vips.Startup(nil)
	defer vips.Shutdown()

	// Convertir l'image en PNG
	new_image, err := vips.NewImageFromBuffer(image)
	if err != nil {
		fmt.Println("Erreur lors de la création de l'image à partir des bytes :", err)
		os.Exit(1)
	}

	pngBytes, err := new_image.ToBytes()
	if err != nil {
		fmt.Println("Erreur lors de la conversion de l'image en PNG :", err)
		os.Exit(1)
	}

	return pngBytes, nil
}
