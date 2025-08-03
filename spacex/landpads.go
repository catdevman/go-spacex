package spacex

import (
	"context"
	"fmt"
)

// LandpadsService handles communication with the landpad related
// methods of the SpaceX API.
type LandpadsService struct {
	client *Client
}

// Landpad represents a SpaceX landpad.
type Landpad struct {
	Name             *string  `json:"name"`
	FullName         *string  `json:"full_name"`
	Status           string   `json:"status"`
	Type             *string  `json:"type"`
	Locality         *string  `json:"locality"`
	Region           *string  `json:"region"`
	Latitude         *float64 `json:"latitude"`
	Longitude        *float64 `json:"longitude"`
	LandingAttempts  int      `json:"landing_attempts"`
	LandingSuccesses int      `json:"landing_successes"`
	Wikipedia        *string  `json:"wikipedia"`
	Details          *string  `json:"details"`
	Launches         []string `json:"launches"`
}

// ListAllLandpads lists all landpads.
func (s *LandpadsService) ListAllLandpads(ctx context.Context) ([]*Landpad, error) {
	u := "landpads"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var landpads []*Landpad
	_, err = s.client.do(ctx, req, &landpads)
	if err != nil {
		return nil, err
	}

	return landpads, nil
}

// GetLandpad retrieves a specific landpad.
func (s *LandpadsService) GetLandpad(ctx context.Context, id string) (*Landpad, error) {
	u := fmt.Sprintf("landpads/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	landpad := new(Landpad)
	_, err = s.client.do(ctx, req, landpad)
	if err != nil {
		return nil, err
	}

	return landpad, nil
}
