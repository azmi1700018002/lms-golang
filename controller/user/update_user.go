package user

import (
	"lms/model"
	ruser "lms/repository/r_user"

	// vuser "lms/validation/v_user"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userUpdateController struct {
	userUpdateRepo ruser.UserUpdateRepository
}

func UpdateUserController(userUpdateRepo ruser.UserUpdateRepository) *userUpdateController {
	return &userUpdateController{
		userUpdateRepo: userUpdateRepo,
	}
}

func (c *userUpdateController) UpdateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan ID user dari parameter route
	idUser := ctx.Param("id_user")

	// Set ID user ke dalam struct user
	user.IDUser = uuid.MustParse(idUser)

	// if err := vuser.ValidateUser(&user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

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

		// hapus gambar lama dari Cloudinary
		if strings.HasPrefix(user.ProfilePicture, "http://res.cloudinary.com") {
			publicID := strings.TrimPrefix(user.ProfilePicture, "http://res.cloudinary.com/ddee7paye/image/upload/")
			publicID = strings.TrimSuffix(publicID, ".jpg")
			_, err = cloudinaryConfig.Admin.DeleteAssetsByPrefix(ctx, admin.DeleteAssetsByPrefixParams{
				Prefix: []string{publicID},
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		user.ProfilePicture = uploadResult.URL
	}

	if err := c.userUpdateRepo.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
