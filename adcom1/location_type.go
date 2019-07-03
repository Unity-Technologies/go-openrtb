package adcom1

// LocationType represents options to indicate how the geographic information was determined.
type LocationType int8

// Options to indicate how the geographic information was determined.
const (
	LocationGPS          LocationType = 1 // GPS/Location Services
	LocationIP           LocationType = 2 // IP Address
	LocationUserProvided LocationType = 3 // User Provided (e.g., registration data)
)
