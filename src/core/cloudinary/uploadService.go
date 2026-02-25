package cloudinary

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(fileBytes []byte, folder, fileName string) (string, error) {
	if CloudinaryInstance == nil {
		return "", fmt.Errorf("Cloudinary no está inicializado")
	}

	ctx := context.Background()

	fileReader := bytes.NewReader(fileBytes)

	uploadResult, err := CloudinaryInstance.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		Folder:       folder,
		ResourceType: "image",
	})

	if err != nil {
		return "", fmt.Errorf("error al subir imagen a Cloudinary: %v", err)
	}

	return uploadResult.SecureURL, nil
}

func UploadAvatar(fileBytes []byte, fileName string) (string, error) {
	return UploadImage(fileBytes, "avatars", fileName)
}

func UploadCourseImage(fileBytes []byte, fileName string) (string, error) {
	return UploadImage(fileBytes, "courses", fileName)
}

func DeleteImage(imageURL string) error {
	if CloudinaryInstance == nil {
		return fmt.Errorf("Cloudinary no está inicializado")
	}

	ctx := context.Background()

	publicID := ExtractPublicID(imageURL)

	_, err := CloudinaryInstance.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
	})

	if err != nil {
		return fmt.Errorf("error al eliminar imagen de Cloudinary: %v", err)
	}

	fmt.Println("Imagen eliminada de Cloudinary:", publicID)
	return nil
}

func ExtractPublicID(imageURL string) string {
	parts := strings.Split(imageURL, "/upload/")
	if len(parts) < 2 {
		return ""
	}

	pathParts := strings.Split(parts[1], "/")
	if len(pathParts) < 2 {
		return ""
	}

	versionIndex := -1
	for i, part := range pathParts {
		if strings.HasPrefix(part, "v") && len(part) > 1 {
			versionIndex = i
			break
		}
	}

	if versionIndex == -1 || versionIndex+1 >= len(pathParts) {
		return ""
	}

	relevantParts := pathParts[versionIndex+1:]
	fileNameWithExtension := relevantParts[len(relevantParts)-1]
	fileName := strings.TrimSuffix(fileNameWithExtension, strings.TrimPrefix(fileNameWithExtension[strings.LastIndex(fileNameWithExtension, "."):], ""))
	folder := strings.Join(relevantParts[:len(relevantParts)-1], "/")

	if folder != "" {
		return folder + "/" + fileName
	}
	return fileName
}
