package geodistance

import (
	"errors"
	"math"
)

const (
	earthRadiusKm = 6371.0 // Earth's radius in kilometers
)

// DistanceUnit represents the unit of measurement for distance
type DistanceUnit string

const (
	Kilometers    DistanceUnit = "km"
	Meters        DistanceUnit = "m"
	Miles         DistanceUnit = "mi"
	NauticalMiles DistanceUnit = "nm"
	Yards         DistanceUnit = "yd"
	Feet          DistanceUnit = "ft"
)

// conversionFactors defines conversion from kilometers to other units
var conversionFactors = map[DistanceUnit]float64{
	Kilometers:    1.0,
	Meters:        1000.0,
	Miles:         0.621371,
	NauticalMiles: 0.539957,
	Yards:         1093.6133,
	Feet:          3280.8399,
}

// GeoPoint represents a geographical point with latitude and longitude
type GeoPoint struct {
	Latitude  float64
	Longitude float64
}

// NewGeoPoint creates a new GeoPoint
func NewGeoPoint(lat, lon float64) (*GeoPoint, error) {
	if lat < -90 || lat > 90 {
		return nil, errors.New("latitude must be in range [-90, 90]")
	}
	if lon < -180 || lon > 180 {
		return nil, errors.New("longitude must be in range [-180, 180]")
	}
	return &GeoPoint{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}

// toRad converts degrees to radians
func (gp *GeoPoint) toRad(n float64) float64 {
	return n * math.Pi / 180
}

// GetDistance calculates the distance to another GeoPoint using the Haversine formula
func (gp *GeoPoint) GetDistance(to *GeoPoint, unit DistanceUnit) (float64, error) {
	dLat := gp.toRad(to.Latitude - gp.Latitude)
	dLon := gp.toRad(to.Longitude - gp.Longitude)
	fromLat := gp.toRad(gp.Latitude)
	toLat := gp.toRad(to.Latitude)

	a := math.Pow(math.Sin(dLat/2), 2) +
		(math.Pow(math.Sin(dLon/2), 2) * math.Cos(fromLat) * math.Cos(toLat))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate distance in kilometers
	distanceKm := earthRadiusKm * c

	// Convert to requested unit
	factor, ok := conversionFactors[unit]
	if !ok {
		return 0, errors.New("invalid distance unit")
	}

	return distanceKm * factor, nil
}
