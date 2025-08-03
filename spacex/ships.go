package spacex

import (
	"context"
	"fmt"
)

// ShipsService handles communication with the ship related
// methods of the SpaceX API.
type ShipsService struct {
	client *Client
}

// Ship represents a SpaceX ship.
type Ship struct {
	Name          string   `json:"name"`
	LegacyID      *string  `json:"legacy_id"`
	Model         *string  `json:"model"`
	Type          *string  `json:"type"`
	Roles         []string `json:"roles"`
	Active        bool     `json:"active"`
	Imo           *int     `json:"imo"`
	Mmsi          *int     `json:"mmsi"`
	Abs           *int     `json:"abs"`
	Class         *int     `json:"class"`
	MassKg        *int     `json:"mass_kg"`
	MassLbs       *int     `json:"mass_lbs"`
	YearBuilt     *int     `json:"year_built"`
	HomePort      *string  `json:"home_port"`
	Status        *string  `json:"status"`
	SpeedKn       *float64 `json:"speed_kn"`
	CourseDeg     *float64 `json:"course_deg"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	LastAisUpdate *string  `json:"last_ais_update"`
	Link          *string  `json:"link"`
	Image         *string  `json:"image"`
	Launches      []string `json:"launches"`
}

// ListAllShips lists all ships.
func (s *ShipsService) ListAllShips(ctx context.Context) ([]*Ship, error) {
	u := "ships"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var ships []*Ship
	_, err = s.client.do(ctx, req, &ships)
	if err != nil {
		return nil, err
	}

	return ships, nil
}

// GetShip retrieves a specific ship.
func (s *ShipsService) GetShip(ctx context.Context, id string) (*Ship, error) {
	u := fmt.Sprintf("ships/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	ship := new(Ship)
	_, err = s.client.do(ctx, req, ship)
	if err != nil {
		return nil, err
	}

	return ship, nil
}
