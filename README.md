# Geodistance Package

The `geodistance` package provides functionality to calculate the distance between two geographical points on Earth using the Haversine formula. It supports distance calculations in both kilometers and miles.

## Installation

To use the `geodistance` package in your Go project, you can import it directly:

```go
import "github.com/boydmeyer/geodistance"
```

Ensure you have Go installed and run:

```bash
go get github.com/boydmeyer/geodistance
```

## Usage

The package provides a `GeoPoint` struct to represent geographical coordinates and methods to compute distances.

### Creating a GeoPoint

Create a new `GeoPoint` with latitude and longitude:

```go
point1 := geodistance.NewGeoPoint(40.7128, -74.0060) // New York
point2 := geodistance.NewGeoPoint(51.5074, -0.1278)  // London
```

### Calculating Distance

Calculate the distance between two points, specifying the unit (`Kilometers` or `Miles`):

```go
distanceKm := point1.GetDistance(*point2, geodistance.Kilometers)
distanceMi := point1.GetDistance(*point2, geodistance.Miles)
```

## Example

Below is a complete example demonstrating how to use the package:

```go
package main

import (
	"fmt"
	"github.com/boydmeyer/geodistance"
)

func main() {
	// Create two geographical points
	newYork := geodistance.NewGeoPoint(40.7128, -74.0060) // New York
	london := geodistance.NewGeoPoint(51.5074, -0.1278)   // London

	// Calculate distance in kilometers
	distanceKm := newYork.GetDistance(*london, geodistance.Kilometers)
	fmt.Printf("Distance in kilometers: %.2f km\n", distanceKm)

	// Calculate distance in miles
	distanceMi := newYork.GetDistance(*london, geodistance.Miles)
	fmt.Printf("Distance in miles: %.2f mi\n", distanceMi)
}
```

### Output

```
Distance in kilometers: 5570.04 km
Distance in miles: 3461.04 mi
```

## Features

- **Haversine Formula**: Accurately calculates distances between two points on a sphere.
- **Unit Support**: Supports both kilometers and miles.
- **Simple API**: Easy-to-use struct and methods for distance calculations.

## License

This package is licensed under the MIT License. See the LICENSE file for details.
