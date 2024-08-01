package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type results struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
}
type account struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
}
type message struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results,omitempty"`
}

func main() {
	data :=
		[]results{
			{
				Id:       1,
				Name:     "fazz",
				Email:    "fazz@mail.com",
				Password: "12345678",
			},
		}
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			message{
				Success: true,
				Message: "OK",
				Results: data,
			})
	})
	r.POST("/users", func(c *gin.Context) {
		user := results{}

		c.Bind(&user)

		user.Id = len(data) + 1

		data = append(data, user)

		c.JSON(http.StatusOK,
			message{
				Success: true,
				Message: "Create User success",
				Results: user,
			})
	})
	r.POST("/auth/login", func(c *gin.Context) {
		user := account{}
		cont := false

		c.Bind(&user)
		for _, x := range data {
			if user.Email == x.Email && user.Password == x.Password {
				cont = true
			}
		}

		if cont {
			c.JSON(http.StatusOK,
				message{
					Success: true,
					Message: "Login Success",
				},
			)
		} else {
			c.JSON(http.StatusUnauthorized,
				message{
					Success: false,
					Message: "Wrong Email Or Password",
				},
			)
		}
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for _, account := range data {
			if account.Id == id {
				c.JSON(http.StatusOK, message{
					Success: true,
					Message: "Users Found",
					Results: []results{account},
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, message{
			Success: false,
			Message: "Users Not Found",
		})
	})
	r.PATCH("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		user := results{}

		c.Bind(&user)

		for i, x := range data {
			if x.Id == id {
				data[i].Name = user.Name
				data[i].Email = user.Email
				data[i].Password = user.Password
				c.JSON(http.StatusOK, message{
					Success: true,
					Message: "Data Users Success",
					Results: data,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, message{
			Success: false,
			Message: "User Not Found",
		})
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for i, x := range data {
			if x.Id == id {
				data = append(data[:i], data[i+1:]...)
				c.JSON(http.StatusOK, message{
					Success: true,
					Message: "Delete Success",
					Results: data,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, message{
			Success: false,
			Message: "Data Not Found",
		})
	})
	r.Run("localhost:8888")
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}