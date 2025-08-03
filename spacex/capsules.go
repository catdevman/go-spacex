package spacex

import (
	"context"
	"fmt"
)

// CapsulesService handles communication with the capsule related
// methods of the SpaceX API.
type CapsulesService struct {
	client *Client
}

// Capsule represents a SpaceX capsule.
type Capsule struct {
	Serial        string   `json:"serial"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Dragon        string   `json:"dragon"`
	ReuseCount    int      `json:"reuse_count"`
	WaterLandings int      `json:"water_landings"`
	LandLandings  int      `json:"land_landings"`
	LastUpdate    *string  `json:"last_update"`
	Launches      []string `json:"launches"`
}

// ListAllCapsules lists all capsules.
func (s *CapsulesService) ListAllCapsules(ctx context.Context) ([]*Capsule, error) {
	u := "capsules"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var capsules []*Capsule
	_, err = s.client.do(ctx, req, &capsules)
	if err != nil {
		return nil, err
	}

	return capsules, nil
}

// GetCapsule retrieves a specific capsule.
func (s *CapsulesService) GetCapsule(ctx context.Context, id string) (*Capsule, error) {
	u := fmt.Sprintf("capsules/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	capsule := new(Capsule)
	_, err = s.client.do(ctx, req, capsule)
	if err != nil {
		return nil, err
	}

	return capsule, nil
}
