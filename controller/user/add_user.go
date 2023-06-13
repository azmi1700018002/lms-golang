package user

import (
	"lms/model"
	ruser "lms/repository/r_user"
	"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepo ruser.UserRepository
}

func NewUserController(userRepo ruser.UserRepository) *userController {
	return &userController{
		userRepo: userRepo,
	}
}

func (c *userController) RegisterUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah file foto profil di-upload
	file, err := ctx.FormFile("profile_picture")
	if err == nil {
		// upload file ke Cloudinary
		cloudinaryConfig, err := cloudinary.NewFromParams("ddee7paye", "898949133356251", "Jn3rtgch_6Api6XU5BWmvBUMsuA")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// convert file ke format yang bisa diupload ke cloudinary
		fileReader, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer fileReader.Close()
		uploadParams := uploader.UploadParams{Format: "jpg"}
		uploadResult, err := cloudinaryConfig.Upload.Upload(ctx, fileReader, uploadParams)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.ProfilePicture = uploadResult.URL
	} else {
		// jika tidak, set avatar CDN sebagai profil gambar
		user.ProfilePicture = "https://www.pngall.com/wp-content/uploads/12/Avatar-No-Background.png"
	}

	if err := c.userRepo.RegisterUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}
