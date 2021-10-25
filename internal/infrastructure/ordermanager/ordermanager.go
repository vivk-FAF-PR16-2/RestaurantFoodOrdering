package ordermanager

import "github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"

type IManager interface {
	Add(data dto.RestaurantData)
	Get() menuData
}

type orderManager struct {
	restaurants []dto.RestaurantData
}

// region Singleton

var instance IManager

func Get() IManager {
	if instance == nil {
		instance = newManager()
	}
	return instance
}

// endregion

func newManager() IManager {
	return &orderManager{
		restaurants: make([]dto.RestaurantData, 0),
	}
}

func (manager orderManager) Add(data dto.RestaurantData) {
	manager.restaurants = append(manager.restaurants, data)
}

func (manager orderManager) Get() menuData {
	count := len(manager.restaurants)
	return menuData{
		count,
		manager.restaurants,
	}
}
