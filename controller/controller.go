package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastructure/orderconvert"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastructure/ordermanager"
	"io/ioutil"
	"log"
)

type IController interface {
	RegisterRoutes(r *gin.Engine)
}

type controller struct {
}

func New() IController {
	return &controller{}
}

func (c *controller) RegisterRoutes(r *gin.Engine) {
	r.GET("/menu", c.getMenu)
	r.POST("/register", c.register)
	r.POST("/order", c.order)
	r.POST("/rating", c.rating)
}

func (c *controller) getMenu(ctx *gin.Context) {
	manager := ordermanager.Get()
	menu := manager.Get()

	ctx.JSON(200, &menu)
}

func (c *controller) register(ctx *gin.Context) {
	var data dto.RestaurantData

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		message := fmt.Sprintf("Error from `/register` route: %v\n", err)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": message,
		})
		log.Panic(message)
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		message := fmt.Sprintf("Error from `/register` route: %v\n", err)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": message,
		})
		log.Panic(message)
	}

	// TODO: Append new `RestaurantData` to menu list

	ctx.JSON(200, gin.H{
		"msg": "ok",
	})
}

func (c *controller) order(ctx *gin.Context) {
	var data dto.ClientInData

	err := json.NewDecoder(ctx.Request.Body).Decode(&data)
	if err != nil {
		message := fmt.Sprintf("Error from `/order` route: %v\n", err)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": message,
		})
		log.Panic(message)
	}

	responseData := orderconvert.OrderConvert(data)
	response, err := json.Marshal(responseData)
	if err != nil {
		message := fmt.Sprintf("Error from `/order` route: %v\n", err)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": message,
		})
		log.Panic(message)
	}

	ctx.JSON(200, &response)
}

func (c *controller) rating(ctx *gin.Context) {
	// TODO: `rating` method implementation
	// res, err := http.Post("http://dinner_hall/rating", "application/json", data from Request)

	ctx.JSON(200, gin.H{
		"msg": "ok",
	})
}
