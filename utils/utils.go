package utils

import (
	"strings"

	"github.com/tomaskukk/lounasta/api"
)

func FilterRestaurantsByName(restaurants []api.Restaurant, name string) []api.Restaurant {
	var filteredRestaurants []api.Restaurant
	name = strings.ToLower(name)

	for _, restaurant := range restaurants {
		if strings.Contains(strings.ToLower(restaurant.Name), name) {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}

	return filteredRestaurants
}

func FilterRestaurantByFood(restaurants []api.Restaurant, food string) []api.Restaurant {
	var filtered []api.Restaurant
	food = strings.ToLower(food) // Normalize case for searching

	for _, restaurant := range restaurants {
		for _, dish := range restaurant.Dishes {
			if strings.Contains(strings.ToLower(dish), food) {
				filtered = append(filtered, restaurant)
				break // No need to check further dishes for this restaurant
			}
		}
	}

	return filtered
}
