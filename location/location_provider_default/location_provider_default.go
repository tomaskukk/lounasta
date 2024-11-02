package location_provider_default

import (
	"encoding/json"
	"errors"
	"fmt"
	"location-v2/location/location_provider"
	"net/http"
	"strconv"
	"strings"
)

type ipLocationResponse struct {
	Loc string `json:"loc"` // This will contain coordinates as "latitude,longitude"
}

// fetchLocationByIP fetches the location coordinates based on IP
func fetchLocationByIP() (float64, float64, error) {
	url := "https://ipinfo.io/json"

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get location by IP: %w", err)
	}
	defer resp.Body.Close()

	var location ipLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return 0, 0, fmt.Errorf("failed to decode IP location response: %w", err)
	}

	// Split "latitude,longitude" into separate float64 values
	coords := strings.Split(location.Loc, ",")
	if len(coords) != 2 {
		return 0, 0, errors.New("invalid location data from IP lookup")
	}

	lat, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse latitude: %w", err)
	}

	lon, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse longitude: %w", err)
	}

	return lat, lon, nil
}

type DefaultLocationManager struct{}

func (m DefaultLocationManager) GetLocation() (location_provider.Coordinate2D, error) {
	lat, lon, err := fetchLocationByIP()

	if err != nil {
		return location_provider.Coordinate2D{}, fmt.Errorf("failed to get location: %w", err)
	}

	return location_provider.Coordinate2D{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}
