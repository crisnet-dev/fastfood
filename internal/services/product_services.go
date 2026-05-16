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

	if err := repository.AddNewProductRepository(product); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetProductService() ([]models.Product, error) {
	products, err := repository.GetProduct()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}
