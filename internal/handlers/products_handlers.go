package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/models"
	"github.com/crisnet-dev/fastfood/internal/services"
	"github.com/crisnet-dev/fastfood/internal/utils"
)

func UploadProduc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.HttpError(w, "Error to process file", 400)
		return
	}

	jsonData := r.FormValue("data")

	var product models.Product

	if err := json.Unmarshal([]byte(jsonData), &product); err != nil {
		utils.HttpError(w, "Invalid JSON", 400)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		utils.HttpError(w, "Error to read file", 400)
		return
	}
	defer file.Close()

	if err := services.UploadProductService(&product, file); err != nil {
		utils.HttpError(w, "Error to upload file", 500) //Validate type error
		return
	}

	utils.HttpResponse(w, map[string]string{
		"message": "Product added",
	}, 200)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := services.GetProductService()
	if err != nil {
		utils.HttpError(w, err.Error(), 500)
		return
	}

	utils.HttpResponse(w, map[string]any{
		"products": products,
	}, 200)
}
