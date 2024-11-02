//go:build darwin
// +build darwin

package location

import (
	"location-v2/location/location_provider"
	"location-v2/location/location_provider_darwin"
)

func getProvider() location_provider.LocationProvider {
	return location_provider_darwin.MacLocationProvider{}
}
