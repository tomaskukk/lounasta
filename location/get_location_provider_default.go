//go:build !darwin
// +build !darwin

package location

import (
	"github.com/tomaskukk/lounasta/location/location_provider"
	"github.com/tomaskukk/lounasta/location/location_provider_default"
)

func getProvider() location_provider.LocationProvider {
	return location_provider_default.DefaultLocationProvider{}
}
