package location_provider

type Coordinate2D struct {
	Latitude  float64
	Longitude float64
}

type LocationProvider interface {
	GetLocation() (Coordinate2D, error)
}
