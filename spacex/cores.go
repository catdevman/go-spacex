package spacex

import (
	"context"
	"fmt"
)

// CoresService handles communication with the core related
// methods of the SpaceX API.
type CoresService struct {
	client *Client
}

// Core represents a SpaceX core.
type Core struct {
	Serial       string   `json:"serial"`
	Block        *int     `json:"block"`
	Status       string   `json:"status"`
	ReuseCount   int      `json:"reuse_count"`
	RTLSAttempts int      `json:"rtls_attempts"`
	RTLSLandings int      `json:"rtls_landings"`
	ASDSAttempts int      `json:"asds_attempts"`
	ASDSLandings int      `json:"asds_landings"`
	LastUpdate   *string  `json:"last_update"`
	Launches     []string `json:"launches"`
}

// ListAllCores lists all cores.
func (s *CoresService) ListAllCores(ctx context.Context) ([]*Core, error) {
	u := "cores"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var cores []*Core
	_, err = s.client.do(ctx, req, &cores)
	if err != nil {
		return nil, err
	}

	return cores, nil
}

// GetCore retrieves a specific core.
func (s *CoresService) GetCore(ctx context.Context, id string) (*Core, error) {
	u := fmt.Sprintf("cores/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	core := new(Core)
	_, err = s.client.do(ctx, req, core)
	if err != nil {
		return nil, err
	}

	return core, nil
}
