package location

import (
	"location-v2/location/location_provider"
)

func test() {
	// Initialize macOS provider on darwin, otherwise use the default provider

}

func GetLocation() (location_provider.Coordinate2D, error) {
	locationProvider := getProvider()

	coordinate, err := locationProvider.GetLocation()
	if err != nil {
		return location_provider.Coordinate2D{}, err
	}
	return coordinate, nil
}
