package location

import (
	"location-v2/location/location_provider"
	"location-v2/location/location_provider_darwin"
	"location-v2/location/location_provider_default"
)

func GetLocation() (location_provider.Coordinate2D, error) {
	macProvider := location_provider_darwin.MacLocationProvider{}

	coordinateMac, err := macProvider.GetLocation()
	if err == nil {
		return coordinateMac, nil
	}

	// Fallback to IP-based location provider if macOS location fails

	defaultProvider := location_provider_default.DefaultLocationManager{}
	coordinateIp, err := defaultProvider.GetLocation()

	if err != nil {
		return location_provider.Coordinate2D{}, err
	}

	return coordinateIp, nil
}
