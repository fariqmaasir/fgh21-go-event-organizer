package controllers

import (
	"fmt"
	"net/http"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
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
				Success: true,
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
			Message: "OK",
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
				Message: "Acc Not Found",
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
				Success: true,
				Message: "Users Failed To Find",
				Results: result,
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
