package main

import (
	"fmt"
	"location-v2/api"
	"location-v2/cli"
	"location-v2/location"

	"location-v2/utils"

	"os"

	"github.com/spf13/cobra"
)

func run(name string, food string) {
	loc, err := location.GetLocation()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var restaurants []api.Restaurant

	lat := loc.Latitude
	lon := loc.Longitude

	restaurants, err = api.FetchLounaat(lat, lon)

	if err != nil {
		return
	}

	if name != "" {
		restaurants = utils.FilterRestaurantsByName(restaurants, name)
	}

	if food != "" {
		restaurants = utils.FilterRestaurantByFood(restaurants, food)
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
