package controllers

import (
	"net/http"
	"strconv"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator"
)

// var validate *validator.Validate

func ListAllUsers(ctx *gin.Context) {
	data := models.FindAllUsers()
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "OK",
			Results: data,
		})
}

func DetailUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindOneUser(id)
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "OK",
			Results: data,
		})
}
func CreateUser(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate = validator.New()
	results := models.CreateOneUser(user)
	// if err := validate.Struct(user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// } else {

	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Create User success",
			Results: results,
		})
	// }
}
func UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user := models.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	results := models.UpdateOneUser(id, user)
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Create User success",
			Results: results,
		})
}
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	models.DeleteOneUser(id)
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Create User success",
		})
}
