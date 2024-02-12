package utils

import (
	geo "github.com/kellydunn/golang-geo"
)

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	point1 := geo.NewPoint(lat1, lon1)
	point2 := geo.NewPoint(lat2, lon2)

	distance := point1.GreatCircleDistance(point2)

	return distance
}
