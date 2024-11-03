package location

import (
	"github.com/tomaskukk/lounasta/location/location_provider"
)

func GetLocation() (location_provider.Coordinate2D, error) {
	locationProvider := getProvider()

	coordinate, err := locationProvider.GetLocation()
	if err != nil {
		return location_provider.Coordinate2D{}, err
	}
	return coordinate, nil
}
