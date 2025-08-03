package spacex

import "context"

// RoadsterService handles communication with the roadster related
// methods of the SpaceX API.
type RoadsterService struct {
	client *Client
}

// Roadster represents the Tesla Roadster launched into space.
type Roadster struct {
	Name            string   `json:"name"`
	LaunchDateUTC   string   `json:"launch_date_utc"`
	LaunchDateUnix  int64    `json:"launch_date_unix"`
	LaunchMassKg    int      `json:"launch_mass_kg"`
	LaunchMassLbs   int      `json:"launch_mass_lbs"`
	NoradID         int      `json:"norad_id"`
	EpochJD         float64  `json:"epoch_jd"`
	OrbitType       string   `json:"orbit_type"`
	ApoapsisAU      float64  `json:"apoapsis_au"`
	PeriapsisAU     float64  `json:"periapsis_au"`
	SemiMajorAxisAU float64  `json:"semi_major_axis_au"`
	Eccentricity    float64  `json:"eccentricity"`
	Inclination     float64  `json:"inclination"`
	Longitude       float64  `json:"longitude"`
	PeriapsisArg    float64  `json:"periapsis_arg"`
	PeriodDays      float64  `json:"period_days"`
	SpeedKph        float64  `json:"speed_kph"`
	SpeedMph        float64  `json:"speed_mph"`
	EarthDistanceKm float64  `json:"earth_distance_km"`
	EarthDistanceMi float64  `json:"earth_distance_mi"`
	MarsDistanceKm  float64  `json:"mars_distance_km"`
	MarsDistanceMi  float64  `json:"mars_distance_mi"`
	FlickrImages    []string `json:"flickr_images"`
	Wikipedia       string   `json:"wikipedia"`
	Video           string   `json:"video"`
	Details         string   `json:"details"`
}

// GetRoadsterInfo retrieves roadster information.
func (s *RoadsterService) GetRoadsterInfo(ctx context.Context) (*Roadster, error) {
	u := "roadster"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	roadster := new(Roadster)
	_, err = s.client.do(ctx, req, roadster)
	if err != nil {
		return nil, err
	}

	return roadster, nil
}
