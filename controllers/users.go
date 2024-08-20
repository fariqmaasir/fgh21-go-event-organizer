package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator"
)

// var validate *validator.Validate

func ListAllUsers(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	data, totalData := models.FindAllEvents(search, page, limit)
	totalPage := math.Ceil(float64(totalData) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "OK",
			Results: data,
			PageInfo: PageInfo{
				TotalData:  totalData,
				TotalPages: int(totalPage),
				PageLimit:  limit,
				Page:       page,
				Next:       next,
				Prev:       prev,
			},
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
