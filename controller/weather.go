package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	db "example.com/openweather/database"
)

type CreateWeatherRequest struct {
	CityId int
	Time   int
	Data   string
}

type UpdateWeatherRequest struct {
	Id     int
	CityId int
	Time   int
	Data   string
}

type DeleteWeatherRequest struct {
	Id int
}

func CreateWeatherRecord(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.POST("/record", func(ctx *gin.Context) {

		body := CreateWeatherRequest{}
		ctx.BindJSON(&body)

		weatherId := db.CreateWeather(body.CityId, body.Time, body.Data)
		if weatherId == -1 {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "Create weather record failed",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
		})
	})

}

func UpdateWeatherRecord(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.PUT("/record/:id", func(ctx *gin.Context) {

		body := UpdateWeatherRequest{}
		ctx.BindJSON(&body)

		id := ctx.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code":   400,
				"status": "Id is invalid",
			})
			return
		}

		body.Id = parsedId

		queryResult := db.UpdateWeather(body.Id, body.CityId, body.Time, body.Data)
		if !queryResult {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "Update weather record failed",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
		})
	})

}

func DeleteWeatherRecord(route *gin.Engine) {
	auth := route.Group("/api/weather")
	auth.DELETE("/record/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code":   400,
				"status": "Id is invalid",
			})
			return
		}

		queryResult := db.DeleteWeather(parsedId)
		if !queryResult {
			ctx.JSON(400, gin.H{
				"code":    400,
				"message": "Delete weather record failed",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"code":   200,
			"status": "OK",
		})
	})

}