//go:build darwin
// +build darwin

package location_provider_darwin

import (
	"fmt"

	"location-v2/location/location_provider"
	"time"
)

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.15
#cgo LDFLAGS: -framework CoreLocation -framework Foundation -L. -llocation

#import "location_provider_darwin.h"
*/
import "C"

type Location struct {
	Coordinate         location_provider.Coordinate2D
	Altitude           float64
	HorizontalAccuracy float64
	VerticalAccuracy   float64
	Timestamp          time.Time
}

type MacLocationProvider struct{}

func (m MacLocationProvider) GetLocation() (location_provider.Coordinate2D, error) {
	var cloc C.Location
	if ret := C.get_current_location(&cloc); int(ret) != 0 {
		return location_provider.Coordinate2D{}, fmt.Errorf("failed to get location, code %d", ret)
	}

	loc := location_provider.Coordinate2D{
		Latitude:  float64(C.float(cloc.coordinate.latitude)),
		Longitude: float64(C.float(cloc.coordinate.longitude)),
	}

	return loc, nil
}
