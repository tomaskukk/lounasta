package main

import (
	"fmt"
	"location-v2/api"
	"location-v2/cli"
	"location-v2/location_manager"

	"os"
	"strings"

	"github.com/spf13/cobra"
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

func run(name string, food string) {
	loc, err := location_manager.CurrentLocation()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var restaurants []api.Restaurant

	lat := loc.Coordinate.Latitude
	lon := loc.Coordinate.Longitude

	restaurants, err = api.FetchLounaat(lat, lon)

	if err != nil {
		return
	}

	if name != "" {
		restaurants = FilterRestaurantsByName(restaurants, name)
	}

	if food != "" {
		restaurants = FilterRestaurantByFood(restaurants, food)
	}

	cli.PrintRestaurants(restaurants, food)

}

func main() {
	var name string
	var food string

	var rootCmd = &cobra.Command{
		Use:   "lounaatapp",
		Short: "CLI for finding lunch restaurants",
		Run: func(cmd *cobra.Command, args []string) {
			if name != "" {
				fmt.Printf("Searching for restaurants with name: %s\n", name)
			} else {
				fmt.Println("Showing all restaurants")
			}
			run(name, food)
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Filter by restaurant name")
	rootCmd.Flags().StringVarP(&food, "food", "f", "", "Filter by food name")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
