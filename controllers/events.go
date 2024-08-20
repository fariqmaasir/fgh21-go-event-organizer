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

type PageInfo struct {
	TotalData  int `json:"totalData"`
	TotalPages int `json:"totalPages"`
	PageLimit  int `json:"pageLimit"`
	Page       int `json:"page"`
	Next       int `json:"next"`
	Prev       int `json:"prev"`
}

func CreateEvents(ctx *gin.Context) {
	event := models.Events{}

	if err := ctx.ShouldBind(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := ctx.Get("userId")
	fmt.Print("----")
	fmt.Print()
	fmt.Println(event)
	result, err := models.CreateOneEvent(event, id.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	event.CreatedBy = id.(int)
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Event created successfully",
		Results: result,
	})
}

func ListAllEvents(ctx *gin.Context) {
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
			Message: "Get All Events",
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

func DetailEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	results := models.FindOneEvent(id)
	fmt.Println(id)
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List Event By User Id",
		Results: results,
	})
}

func UserCreatedEvent(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	results := models.FindOneEventByUserId(id)
	fmt.Println(id)
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List Event",
		Results: results,
	})
}

func UpdateEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	newEvent := models.Events{}

	if err := ctx.ShouldBind(&newEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}
	createId, _ := ctx.Get("userId")
	err := models.EditOneEvent(newEvent, id, createId.(int))
	fmt.Println(err)
	fmt.Println("--")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: "Failed to create Event",
		})
		return
	}
	newEvent.CreatedBy = createId.(int)
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Event updated successfully",
		Results: newEvent,
	})
}

func DeleteEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result, err := models.DeleteOneEvent(id)
	fmt.Println(result)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Event Doesn't Find",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Delete Success",
		Results: result,
	})
}

func SectionEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result, err := models.FindSectionsByEventId(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Section Doesn't Find",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Section Found",
		Results: result,
	})
}

func FindAllPaymentMethod(ctx *gin.Context) {
	result, err := models.ListAllPaymentMethod()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Payment Not Found",
			Results: result,
		})
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Payment",
		Results: result,
	})
}

func CreateWishlist(ctx *gin.Context) {
	eventId, _ := strconv.Atoi(ctx.Param("id"))
	userId := ctx.GetInt("userId")

	result, err := models.CreateOneWishlist(eventId, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Message{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	fmt.Println(err)
	// wish.UserId = id
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Add Event To Wishlist",
		Results: result,
	})
}

func DeleteWishlist(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	results, err := models.DeleteWishlistById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Wishlist Doesn't Find",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Delete Success",
		Results: results,
	})
}

func FindWishlist(ctx *gin.Context) {

	id := ctx.GetInt("userId")
	wish, err := models.GetWishlistByUserId(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Wishlist Doesn't Find",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Wishlist Finded",
		Results: wish,
	})
}
