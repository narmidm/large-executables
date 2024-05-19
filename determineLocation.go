package main

import (
	"math"
)

// Region Define a struct to hold region information
type Region struct {
	Name string
	Lat  float64
	Long float64
}

// Define a slice to hold all regions
var regions = []Region{
	{"asia-east1", 25.0330, 121.5654},                  // Taiwan
	{"asia-east2", 22.3964, 114.1095},                  // Hong Kong
	{"asia-northeast1", 35.6895, 139.6917},             // Tokyo, Japan
	{"asia-northeast2", 34.6937, 135.5023},             // Osaka, Japan
	{"asia-northeast3", 37.5665, 126.9780},             // Seoul, South Korea
	{"asia-south1", 19.0760, 72.8777},                  // Mumbai, India
	{"asia-south2", 23.0225, 72.5714},                  // Delhi, India
	{"asia-southeast1", 1.3521, 103.8198},              // Singapore
	{"asia-southeast2", -6.2088, 106.8456},             // Jakarta, Indonesia
	{"australia-southeast1", -33.8688, 151.2093},       // Sydney, Australia
	{"australia-southeast2", -37.8136, 144.9631},       // Melbourne, Australia
	{"europe-north1", 60.1699, 24.9384},                // Finland
	{"europe-west1", 53.3478, -6.2597},                 // Dublin, Ireland
	{"europe-west2", 51.5074, -0.1278},                 // London, UK
	{"europe-west3", 50.1109, 8.6821},                  // Frankfurt, Germany
	{"europe-west4", 48.8566, 2.3522},                  // Paris, France
	{"europe-west6", 47.3769, 8.5417},                  // Zurich, Switzerland
	{"northamerica-northeast1", 45.5017, -73.5673},     // Montreal, Canada
	{"northamerica-northeast2", 43.651070, -79.347015}, // Toronto, Canada
	{"southamerica-east1", -23.5505, -46.6333},         // São Paulo, Brazil
	{"southamerica-west1", -12.0464, -77.0428},         // Santiago, Chile
	{"us-central1", 41.8781, -87.6298},                 // Chicago, USA
	{"us-east1", 33.8361, -84.3879},                    // Atlanta, USA
	{"us-east4", 35.2271, -80.8431},                    // North Carolina, USA
	{"us-west1", 37.7749, -122.4194},                   // San Francisco, USA
	{"us-west2", 34.0522, -118.2437},                   // Los Angeles, USA
	{"us-west3", 45.5051, -122.6750},                   // Oregon, USA
	{"us-west4", 40.7128, -74.0060},                    // New York, USA
}

func determineLocation(latitude, longitude float64) string {
	nearestRegion := ""
	minDistance := math.MaxFloat64

	for _, region := range regions {
		distance := haversine(latitude, longitude, region.Lat, region.Long)
		if distance < minDistance {
			minDistance = distance
			nearestRegion = region.Name
		}
	}

	return nearestRegion
}

// Haversine function to calculate the distance between two lat-long points
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in kilometers
	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180.0))*math.Cos(lat2*(math.Pi/180.0))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
