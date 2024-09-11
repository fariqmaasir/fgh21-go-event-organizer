package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProfileUser(ctx *gin.Context) {
	account := models.Regist{}

	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.CreateProfile(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := models.FindOneProfile(profile.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Users Failed To Created",
				Results: result,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Create User success",
			Results: result,
		})
}

func DetailProfile(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	data, err := models.FindOneProfile(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Acc Not Found",
				// Results: data,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Create Profile Success",
			Results: data,
		})
}

func ChangeUserPassword(ctx *gin.Context) {
	account := models.ChangePass{}
	id := ctx.GetInt("userId")

	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := models.ChangePassword(account, id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Incorrect Password",
				// Results: data,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Change Password Success",
			Results: data,
		})
}

func EditProfileUser(ctx *gin.Context) {
	account := models.Regist{}
	id := ctx.GetInt("userId")

	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.EditProfile(account, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := models.FindOneProfile(profile.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Users Failed To Find",
				// Results: result,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Edit User success",
			Results: result,
		})
}

func UploadProfileImage(c *gin.Context) {
	id := c.GetInt("userId")
	fmt.Println(id)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest,
			lib.Message{
				Success: true,
				Message: "no files uploaded",
				// Results: result,
			})
		return
	}
	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		c.JSON(http.StatusBadRequest,
			lib.Message{
				Success: true,
				Message: "no files uploaded",
				// Results: result,
			})
		return
	}

	image := uuid.New().String() + fileExt

	root := "./images/"
	if err := c.SaveUploadedFile(file, root+image); err != nil {
		c.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Upload image failed",
			})
		return
	}
	fmt.Println(err)
	img := "http://localhost:8888/images/" + image

	result, err := models.UpdateProfileImage(models.Profile{Picture: &img}, id)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			lib.Message{
				Success: true,
				Message: "Update image failed",
			})
		return
	}
	c.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Upload image success",
			Results: result,
		})
}
