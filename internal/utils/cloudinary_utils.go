package utils

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/crisnet-dev/fastfood/internal/config"
)

func UploadFileToCloudinary(file multipart.File) (*uploader.UploadResult, error) {
	env := config.GetEnv()

	cld, err := cloudinary.NewFromParams(env.CLOUD_NAME, env.API_KEY, env.API_SECRET)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	response, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "products",
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
