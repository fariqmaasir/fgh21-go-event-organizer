package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func CreateEventCategory(ctx *gin.Context) {
	EventCategory := models.EventCategory{}

	if err := ctx.ShouldBind(&EventCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := models.CreateOneEventCategory(EventCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Event Category created successfully",
		Results: result,
	})
}

func ListAllEventCategory(ctx *gin.Context) {
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
			Message: "Get All EventCategory",
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

func DetailEventCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	results := models.FindOneEventCategory(id)
	fmt.Println(id)
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List EventCategory",
		Results: results,
	})
}

func UpdateEventCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	newEventCategory := models.EventCategory{}

	if err := ctx.ShouldBind(&newEventCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}
	createId, _ := ctx.Get("userId")
	err := models.EditOneEventCategory(newEventCategory, id, createId.(int))
	fmt.Println(err)
	fmt.Println("--")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to create EventCategory",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "EventCategory updated successfully",
		Results: newEventCategory,
	})
}

func DeleteEventCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result, err := models.DeleteOneEventCategory(id)
	fmt.Println(result)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Account Doesn't Find",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Delete Success",
		Results: result,
	})
}
