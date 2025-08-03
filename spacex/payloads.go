package spacex

import (
	"context"
	"fmt"
)

// PayloadsService handles communication with the payload related
// methods of the SpaceX API.
type PayloadsService struct {
	client *Client
}

// Payload represents a SpaceX payload.
type Payload struct {
	Name            *string        `json:"name"`
	Type            *string        `json:"type"`
	Reused          bool           `json:"reused"`
	Launch          *string        `json:"launch"`
	Customers       []string       `json:"customers"`
	NoradIDs        []int          `json:"norad_ids"`
	Nationalities   []string       `json:"nationalities"`
	Manufacturers   []string       `json:"manufacturers"`
	MassKg          *int           `json:"mass_kg"`
	MassLbs         *float64       `json:"mass_lbs"`
	Orbit           *string        `json:"orbit"`
	ReferenceSystem *string        `json:"reference_system"`
	Regime          *string        `json:"regime"`
	Longitude       *float64       `json:"longitude"`
	SemiMajorAxisKm *float64       `json:"semi_major_axis_km"`
	Eccentricity    *float64       `json:"eccentricity"`
	PeriapsisKm     *float64       `json:"periapsis_km"`
	ApoapsisKm      *float64       `json:"apoapsis_km"`
	InclinationDeg  *float64       `json:"inclination_deg"`
	PeriodMin       *float64       `json:"period_min"`
	LifespanYears   *int           `json:"lifespan_years"`
	Epoch           *string        `json:"epoch"`
	MeanMotion      *float64       `json:"mean_motion"`
	Raan            *float64       `json:"raan"`
	ArgOfPericenter *float64       `json:"arg_of_pericenter"`
	MeanAnomaly     *float64       `json:"mean_anomaly"`
	Dragon          *DragonPayload `json:"dragon"`
}

// DragonPayload represents dragon specific payload info.
type DragonPayload struct {
	Capsule         *string  `json:"capsule"`
	MassReturnedKg  *float64 `json:"mass_returned_kg"`
	MassReturnedLbs *float64 `json:"mass_returned_lbs"`
	FlightTimeSec   *int     `json:"flight_time_sec"`
	Manifest        *string  `json:"manifest"`
	WaterLanding    *bool    `json:"water_landing"`
	LandLanding     *bool    `json:"land_landing"`
}

// ListAllPayloads lists all payloads.
func (s *PayloadsService) ListAllPayloads(ctx context.Context) ([]*Payload, error) {
	u := "payloads"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var payloads []*Payload
	_, err = s.client.do(ctx, req, &payloads)
	if err != nil {
		return nil, err
	}

	return payloads, nil
}

// GetPayload retrieves a specific payload.
func (s *PayloadsService) GetPayload(ctx context.Context, id string) (*Payload, error) {
	u := fmt.Sprintf("payloads/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	payload := new(Payload)
	_, err = s.client.do(ctx, req, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
