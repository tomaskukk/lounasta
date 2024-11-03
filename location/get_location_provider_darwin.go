//go:build darwin
// +build darwin

package location

import (
	"github.com/tomaskukk/lounasta/location/location_provider"
	"github.com/tomaskukk/lounasta/location/location_provider_darwin"
)

func getProvider() location_provider.LocationProvider {
	return location_provider_darwin.MacLocationProvider{}
}
