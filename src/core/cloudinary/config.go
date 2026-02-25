package cloudinary

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

var CloudinaryInstance *cloudinary.Cloudinary

func InitCloudinary() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Advertencia: No se pudo cargar .env: %v", err)
	}

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	if cloudName == "" {
		log.Fatal("ERROR: CLOUDINARY_CLOUD_NAME vacío")
	}
	if apiKey == "" {
		log.Fatal("ERROR: CLOUDINARY_API_KEY vacío")
	}
	if apiSecret == "" {
		log.Fatal("ERROR: CLOUDINARY_API_SECRET vacío")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("ERROR al crear instancia de Cloudinary: %v", err)
	}

	CloudinaryInstance = cld
	fmt.Println("✅ Cloudinary configurado correctamente")
}
