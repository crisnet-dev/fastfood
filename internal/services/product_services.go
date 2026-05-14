package services

import (
	"log"
	"mime/multipart"

	"github.com/crisnet-dev/fastfood/internal/models"
	"github.com/crisnet-dev/fastfood/internal/repository"
	"github.com/crisnet-dev/fastfood/internal/utils"
)

func UploadProductService(product *models.Product, file multipart.File) error {
	response, err := utils.UploadFileToCloudinary(file)
	if err != nil {
		log.Println(err)
		return err
	}

	product.ImageURL = response.SecureURL

	log.Print(response.SecureURL)

	repository.AddNewProductRepository(product)

	return nil
}
