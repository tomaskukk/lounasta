package main

import (
	"fmt"
	"location-v2/api"
	"location-v2/cli"

	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.15
#cgo LDFLAGS: -framework CoreLocation -framework Foundation -L. -llocation

#import "location_manager_darwin.h"
*/
import "C"

type Location struct {
	Coordinate         Coordinate2D
	Altitude           float64
	HorizontalAccuracy float64
	VerticalAccuracy   float64
	Timestamp          time.Time
}

type Coordinate2D struct {
	Latitude  float64
	Longitude float64
}

func CurrentLocation() (Location, error) {
	var cloc C.Location
	if ret := C.get_current_location(&cloc); int(ret) != 0 {
		return Location{}, fmt.Errorf("failed to get location, code %d", ret)
	}

	loc := Location{
		Coordinate: Coordinate2D{
			Latitude:  float64(C.float(cloc.coordinate.latitude)),
			Longitude: float64(C.float(cloc.coordinate.longitude)),
		},
		Altitude:           float64(C.float(cloc.altitude)),
		HorizontalAccuracy: float64(C.float(cloc.horizontalAccuracy)),
		VerticalAccuracy:   float64(C.float(cloc.verticalAccuracy)),
	}
	return loc, nil
}

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
	loc, err := CurrentLocation()
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
