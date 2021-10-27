package orderconvert

import (
	"encoding/json"
	"fmt"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastructure/orderid"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastructure/ordermanager"
	"github.com/vivk-FAF-PR16-2/RestaurantDinnerHall/internal/infrastructure/sendrequest"
	"log"
)

func OrderConvert(source dto.ClientInData) dto.ClientOutData {
	idProvider := orderid.Get()
	id := idProvider.Get()
	count := len(source.Orders)
	orders := make([]dto.OrderOutData, count)

	for i := 0; i < count; i++ {
		orders[i] = dto.OrderOutData{
			RestaurantId: source.Orders[i].RestaurantId,
			CreatedTime:  source.Orders[i].CreatedTime,
			OrderId:      id,
		}
	}

	restaurants := ordermanager.Get().Get().RestaurantsData
	for _, order := range orders {
		var restaurantIndex int = -1
		for i := range restaurants {
			if restaurants[i].RestaurantId == order.RestaurantId {
				restaurantIndex = i
				break
			}
		}

		if restaurantIndex == -1 {
			log.Fatalln("error: wrong restaurant id in order!")
		}

		restaurant := restaurants[restaurantIndex]
		url := fmt.Sprintf("http://%s/v2/order", restaurant.Address)
		payload, err := json.Marshal(source.Orders[restaurantIndex])
		if err != nil {
			message := fmt.Sprintf("error: %v\n", err)
			log.Fatal(message)
		}

		result := sendrequest.Post(url, payload)
		err = json.NewDecoder(result.Body).Decode(&order)
		if err != nil {
			message := fmt.Sprintf("error: %v\n", err)
			log.Fatal(message)
		}
	}

	return dto.ClientOutData{
		OrderId: id,
		Orders:  orders,
	}
}
