package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/crisnet-dev/fastfood/internal/config"
	"github.com/crisnet-dev/fastfood/internal/utils"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	env := config.GetEnv()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.HttpError(w, err.Error(), 400)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		utils.HttpError(w, err.Error(), 400)
		return
	}
	defer file.Close()

	cld, err := cloudinary.NewFromParams(env.CLOUD_NAME, env.API_KEY, env.API_SECRET)
	if err != nil {
		log.Println(err)
		utils.HttpError(w, "Error to upload file", 500)
		return
	}

	ctx := context.Background()

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "products",
	})
	if err != nil {
		log.Println(err)
		utils.HttpError(w, "Error to upload file", 500)
		return
	}

	utils.HttpResponse(w, map[string]string{
		"url": resp.SecureURL,
	}, 200)
}
