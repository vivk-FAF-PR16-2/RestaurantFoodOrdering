package ordermanager

import "github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"

type menuData struct {
	Restaurants     int                  `json:"restaurants"`
	RestaurantsData []dto.RestaurantData `json:"restaurants_data"`
}
